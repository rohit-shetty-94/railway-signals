package track

import "github.com/pushpanthraj/crosstech/railway-signals/entity/signal"

type Track struct {
	ID      int             `json:"track_id" pg:",pk"`
	Source  string          `json:"source"`
	Target  string          `json:"target"`
	Signals []signal.Signal `json:"signal_ids" pg:"rel:has-many"`
}
