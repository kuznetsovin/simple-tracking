package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/streadway/amqp"
	"net/http"
	"os"
	"simple-tracking/backend/handlers"
	"simple-tracking/backend/models"
)

var (
	config Config
)

func main() {
	var (
		err error
	)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	if len(os.Args) == 2 {
		if err := config.Load(os.Args[1]); err != nil {
			e.Logger.Fatal(err)
		}
	} else {
		e.Logger.Fatal("Не задан файл конфигурации")
	}

	e.Logger.SetLevel(config.GetLogLever())

	h := handlers.Handler{
		BrokerCfg: config.Broker,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	h.BrokerConn, err = amqp.Dial(config.Broker.Conn)
	if err != nil {
		e.Logger.Fatalf("Ошибка соединения с RabbitMQ: %v", err)
	}
	defer h.BrokerConn.Close()

	h.ErrChan = h.BrokerConn.NotifyClose(make(chan *amqp.Error))

	if h.DB, err = models.NewDataStore(config.DbConn); err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/ws", h.GeoWebsocket)

	api := e.Group("/api")
	api.GET("/last-positions", h.LastPosition)
	api.GET("/tracks/:client", h.Track)

	api.GET("/geo-objects", h.GeoObjects)
	api.POST("/geo-objects", h.AddObject)

	api.GET("/vehicle-dict", h.GetVehicles)
	api.POST("/vehicle-dict", h.AddVehicle)

	report := api.Group("/report")
	report.GET("/object-dist/:client", h.ReportObjectDist)

	e.Logger.Fatal(e.Start(config.Addr))
}
