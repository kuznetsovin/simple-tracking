package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	"simple-tracking/backend/models"
	"simple-tracking/backend/utils"
)

type Handler struct {
	BrokerConn *amqp.Connection
	BrokerCfg  utils.BrokerCfg
	Upgrader   websocket.Upgrader
	ErrChan    chan *amqp.Error
	DB         models.DbStore
}
