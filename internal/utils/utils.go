package utils

import (
	"io"
	"net"
)

func SendMessage(conn net.Conn, msg []byte) error {
	_, err := conn.Write(msg)
	return err
}

func ReceiveMessage(conn net.Conn) ([]byte, error) {
	return io.ReadAll(conn)
}
