package nats

import (
	"github.com/nats-io/nats.go"
)

func getConnection(host, port, user, password string) (*nats.Conn, error) {
	//url := fmt.Sprintf("nats://%s:%s@%s:%s", user, password, host, port)
	connection, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		return nil, err
	}
	return connection, nil
}
