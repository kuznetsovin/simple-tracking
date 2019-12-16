/*
File describe main handle structure which includes broker and db connection.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/
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
