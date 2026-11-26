package store

import "github.com/jackc/pgx/v5/pgxpool"

// EventStore is the only code that talks to the events table.
type EventStore struct {
	pool *pgxpool.Pool
}

func NewEventStore(pool *pgxpool.Pool) *EventStore {
	return &EventStore{pool: pool}
}

// We will add Insert and other method as required
