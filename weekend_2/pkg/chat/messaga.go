package chat

import (
	"encoding/json"
	"net"
)

const (
	MsgTypeUnknown MessageType = iota
	// Client Side
	MsgTypeRegisterReq
	MsgTypeUnregisterReq
	MsgTypeJoinReq
	MsgTypeLeaveReq
	MsgTypeSendReq
	// Server Side
	MsgTypeRes
	MsgTypeChat
)

type MessageType int

type Message struct {
	Type MessageType
	Data json.RawMessage
}

func SendMessage(to net.Conn, msg Message) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = to.Write(append(b, '\n'))
	return err
}

func sendResponse(to net.Conn, reqType MessageType, err error) error {
	var errMessage *string
	if err != nil {
		errString := err.Error()
		errMessage = &errString
	}

	b, err := json.Marshal(struct {
		ReqType MessageType
		Error   *string
	}{
		ReqType: reqType,
		Error:   errMessage,
	})
	if err != nil {
		return err
	}

	return SendMessage(to, Message{
		Type: MsgTypeRes,
		Data: b,
	})
}
