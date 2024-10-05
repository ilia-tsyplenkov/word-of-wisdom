package utils

import (
	"encoding/binary"
	"io"
	"net"
)

func SendMessage(conn net.Conn, msg []byte) error {
	// to use same connection more than once
	// need to let reader know size of a message
	if err := binary.Write(conn, binary.LittleEndian, uint32(len(msg))); err != nil {
		return err
	}
	// now we can send a message
	_, err := conn.Write(msg)
	return err
}

func ReceiveMessage(conn net.Conn) ([]byte, error) {
	var msgSize uint32
	// read size of incoming message first
	if err := binary.Read(conn, binary.LittleEndian, &msgSize); err != nil {
		return nil, err
	}
	msg := make([]byte, msgSize)
	_, err := io.ReadFull(conn, msg)
	return msg, err
}
