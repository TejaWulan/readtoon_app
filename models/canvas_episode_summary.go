package models

type CanvasEpisodeSummary struct {
	EpisodeUID string   `json:"episode_uid"`
	Photos     []string `json:"photos"`
}
