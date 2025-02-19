package signal

// type Signal struct {
// 	ID       int      `json:"id" pg:",pk"`
// 	SignalID int      `json:"signal_id"`
// 	Name     *string  `json:"name,omitempty"`
// 	ELR      *string  `json:"elr,omitempty"`
// 	Mileage  *float64 `json:"mileage,omitempty"`
// 	TrackID  int      `json:"track_id"`
// }

type Signal struct {
	ID       int     `json:"id" pg:",pk"`
	SignalID int     `json:"signal_id"`
	Name     *string `json:"signal_name"`
	ELR      string  `json:"elr"`
	Mileage  float64 `json:"mileage"`
	TrackID  int     `json:"track_id"`
}
