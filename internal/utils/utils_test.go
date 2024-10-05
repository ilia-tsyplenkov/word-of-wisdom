package utils_test

import (
	"io"
	"net"
	"testing"

	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	client, server := net.Pipe()
	defer func() {
		server.Close()
	}()
	msg := []byte("Hi there!")
	go func() {
		err := utils.SendMessage(client, msg)
		require.NoError(t, err)
		client.Close()
	}()

	received, err := io.ReadAll(server)
	require.NoError(t, err)
	require.Equal(t, msg, received)
}

func TestReceiveMessage(t *testing.T) {
	client, server := net.Pipe()
	defer func() {
		server.Close()
	}()
	msg := []byte("Hi there!")
	go func() {
		_, err := client.Write(msg)
		require.NoError(t, err)
		client.Close()
	}()

	received, err := utils.ReceiveMessage(server)
	require.NoError(t, err)
	require.Equal(t, msg, received)

}
