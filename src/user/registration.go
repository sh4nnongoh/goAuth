package user

import (
	"errors"
	"time"
)

// ErrUnknown is used when a user could not be found.
var ErrUnknown = errors.New("unknown user")

// RegistrationActivity represents how and where a cargo can be handled, and can
// be used to express predictions about what is expected to happen to a cargo
// in the future.
type RegistrationActivity struct {
	Type RegistrationEventType
}

// RegistrationEvent is used to register the event when, for instance, a cargo is
// unloaded from a carrier at a some location at a given time.
type RegistrationEvent struct {
	UserID   UserID
	Activity RegistrationActivity
}

// RegistrationEventType describes type of a handling event.
type RegistrationEventType int

// Valid handling event types.
const (
	NotHandled RegistrationEventType = iota
	Registered
)

func (t RegistrationEventType) String() string {
	switch t {
	case NotHandled:
		return "Not Handled"
	case Registered:
		return "Registered"
	}
	return ""
}

// RegistrationHistory is the Registration history of a user.
type RegistrationHistory struct {
	RegistrationEvents []RegistrationEvent
}

// MostRecentlyCompletedEvent returns most recently completed Registration event.
func (h RegistrationHistory) MostRecentlyCompletedEvent() (RegistrationEvent, error) {
	if len(h.RegistrationEvents) == 0 {
		return RegistrationEvent{}, errors.New("delivery history is empty")
	}

	return h.RegistrationEvents[len(h.RegistrationEvents)-1], nil
}

// RegistrationEventRepository provides access a registration event store.
type RegistrationEventRepository interface {
	Store(e RegistrationEvent)
	QueryRegistrationHistory(UserID) RegistrationHistory
}

// RegistrationEventFactory creates registation events
type RegistrationEventFactory struct {
	UserRepository Repository
}

// CreateRegistrationEvent creates a validated registration event.
func (f *RegistrationEventFactory) CreateRegistrationEvent(registered time.Time, completed time.Time, id UserID,
	eventType RegistrationEventType) (RegistrationEvent, error) {

	if _, err := f.UserRepository.Find(id); err != nil {
		return RegistrationEvent{}, err
	}

	return RegistrationEvent{
		UserID: id,
		Activity: RegistrationActivity{
			Type: eventType,
		},
	}, nil
}
