package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type Listener struct {
	conn *pgx.Conn
}

func NewListener(databaseURL string) (*Listener, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &Listener{conn: conn}, nil
}

func (l *Listener) Listen(ctx context.Context, channel string) (<-chan string, error) {
	_, err := l.conn.Exec(ctx, "LISTEN "+channel)
	if err != nil {
		return nil, fmt.Errorf("unable to start listening: %v", err)
	}

	notifications := make(chan string)
	go func() {
		for {
			notification, err := l.conn.WaitForNotification(ctx)
			if err != nil {
				log.Printf("Error waiting for notification: %v", err)
				continue
			}
			notifications <- notification.Payload
		}
	}()

	return notifications, nil
}

func (l *Listener) Close() error {
	return l.conn.Close(context.Background())
}
