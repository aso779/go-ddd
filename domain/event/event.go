package event

import "time"

type DomainEvent interface {
	OccurredOn() time.Time
}
