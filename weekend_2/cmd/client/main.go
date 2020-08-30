package main

import (
	"bufio"
	"codelab/pkg/chat"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	chat.SendMessage(conn, chat.Message{
		Type: chat.MsgTypeRegisterReq,
		Data: MustJSONMarshal("larry"),
	})
	chat.SendMessage(conn, chat.Message{
		Type: chat.MsgTypeUnregisterReq,
	})

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}

func MustJSONMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}
