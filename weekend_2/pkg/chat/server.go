package chat

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	addr *net.TCPAddr

	mu       *sync.Mutex
	registry *UserRegistry
}

func NewServer(addr string) (*Server, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		addr:     tcpAddr,
		mu:       &sync.Mutex{},
		registry: NewUserRegistry(),
	}, nil
}

func (s *Server) Listen() error {
	l, err := net.ListenTCP("tcp", s.addr)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		log.Println("New Connection:", conn.RemoteAddr())

		go s.handleConnection(conn)
	}

	return nil
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		s.unregisterUser(conn)
		conn.Close()
	}()

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		b := sc.Bytes()

		var msg Message
		err := json.Unmarshal(b, &msg)
		if err != nil {
			log.Printf("Unmarshal Error [%s]: %v]\n", conn.RemoteAddr(), err)
		}

		err = s.handleMessage(conn, msg)
		if err != nil {
			log.Printf("Message Handling Error [%s]: %v\n", conn.RemoteAddr(), err)
		}
	}

	err := sc.Err()
	if err != nil {
		log.Printf("Disconnected By Error [%s]: %v\n", conn.RemoteAddr(), err)
		return
	}

	log.Printf("Disconnected: [%s]\n", conn.RemoteAddr())
}

func (s *Server) handleMessage(from net.Conn, msg Message) error {
	var err error
	switch msg.Type {
	case MsgTypeRegisterReq:
		var name string
		err = json.Unmarshal(msg.Data, &name)
		if err != nil {
			break
		}

		err = s.registerUser(from, name)
	case MsgTypeUnregisterReq:
		err = s.unregisterUser(from)
	default:
		err = fmt.Errorf("invalid message type: %v", msg.Type)
	}

	return sendResponse(from, msg.Type, err)
}

func (s *Server) registerUser(conn net.Conn, name string) error {
	if name == "" {
		return errors.New("invalid user name")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if user, ok := s.registry.Conns[conn]; ok {
		return fmt.Errorf("user already registered: %s", user.Name)
	}

	if _, ok := s.registry.Names[name]; ok {
		return fmt.Errorf("user already used: %s", name)
	}

	user := User{
		Conn: conn,
		Name: name,
	}
	s.registry.Conns[user.Conn] = user
	s.registry.Names[user.Name] = user

	return nil
}

func (s *Server) unregisterUser(conn net.Conn) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.registry.Conns[conn]
	if !ok {
		return errors.New("user is not registered")
	}

	delete(s.registry.Conns, user.Conn)
	delete(s.registry.Names, user.Name)

	return nil
}
