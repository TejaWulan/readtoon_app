package models

type MusicVideoSummary struct {
	MusicVideoUID string `json:"music_video_uid"`
	Photo         string `json:"photo"`
	Viewer        int    `json:"viewer"`
	Likes         int    `json:"likes"`
}
