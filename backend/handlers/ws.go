package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"simple-tracking/backend/utils"
)

func (h *Handler) GeoWebsocket(c echo.Context) error {
	ws, err := h.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	channel, err := h.BrokerConn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	if err = channel.Qos(h.BrokerCfg.PrefetchCount, h.BrokerCfg.PrefetchSize, h.BrokerCfg.QosGlobal); err != nil {
		return err
	}

	msgs, err := channel.Consume(
		h.BrokerCfg.Queue, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-msgs:
			c.Logger().Debugf("Получем сообщение из очереди")
			if err = msg.Ack(false); err != nil {
				return err
			}

			c.Logger().Debugf("Подтверждаем что считали сообщение")

			trackPoint := utils.BrokerMsg{}

			c.Logger().Debugf("Разбираем сообщение")
			if err := trackPoint.FromBytes(msg.Body); err != nil {
				return err
			}

			c.Logger().Debugf("Преобразовываем Feature")
			point, err := trackPoint.ToGeoFeature()
			if err != nil {
				return err
			}

			c.Logger().Debugf("Преобразуем сообщение в geojson")
			geoPkg, err := point.MarshalJSON()
			if err != nil {
				return err
			}

			if err := ws.WriteMessage(websocket.TextMessage, geoPkg); err != nil {
				return err
			}

		case notify := <-h.ErrChan:
			return fmt.Errorf("Ошибка соединения с RabbitMQ: %v", notify)
		}
	}
}
