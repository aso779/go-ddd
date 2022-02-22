package events

import "time"

type DomainEvent interface {
	OccurredOn() time.Time
}
