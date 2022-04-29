package event

import "time"

type DomainEvent interface {
	GetName() string
	OccurredOn() time.Time
}
