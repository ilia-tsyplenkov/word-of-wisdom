package utils_test

import (
	"net"
	"testing"

	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestSendReceive(t *testing.T) {
	client, server := net.Pipe()
	defer func() {
		client.Close()
		server.Close()
	}()
	msg := []byte("Hi there!")
	go func() {
		err := utils.SendMessage(client, msg)
		require.NoError(t, err)
	}()

	received, err := utils.ReceiveMessage(server)
	require.NoError(t, err)
	require.Equal(t, msg, received)
}

func TestSendReceiveMoreThanOnce(t *testing.T) {
	client, server := net.Pipe()
	defer func() {
		client.Close()
		server.Close()
	}()
	messages := [][]byte{[]byte("First one"), []byte("Second one")}

	go func() {
		for _, msg := range messages {
			err := utils.SendMessage(client, msg)
			require.NoError(t, err)
		}
	}()

	for _, expected := range messages {
		received, err := utils.ReceiveMessage(server)
		require.NoError(t, err)
		require.Equal(t, expected, received)
	}
}
