package audio

import "time"

type Song struct {
	Name     string
	Artist   string
	Album    string
	Duration time.Time
}
