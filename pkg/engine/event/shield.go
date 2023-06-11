package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type ShieldAddedEventHandler = handler.EventHandler[ShieldAddedEvent]
type ShieldAddedEvent struct {
	ID           key.Shield
	Info         info.Shield
	ShieldHealth float64
}

type ShieldRemovedEventHandler = handler.EventHandler[ShieldRemovedEvent]
type ShieldRemovedEvent struct {
	ID     key.Shield
	Target key.TargetID
}

type ShieldChangeEventHandler = handler.EventHandler[ShieldChangeEvent]
type ShieldChangeEvent struct {
	ID     key.Shield
	Target key.TargetID
	OldHP  float64
	NewHP  float64
}