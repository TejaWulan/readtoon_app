package models

type MusicVideoFavouriteItem struct {
	MusicVideoUID string `json:"music_video_uid"`
	Photo         string `json:"photo"` // thumbnail
	Viewer        int    `json:"viewer"`
	Likes         int    `json:"likes"`
	AccountName   string `json:"account_name"` // nama_akun of video owner
	AccountUID    string `json:"account_uid"`  // nama_akun_uid of video owner
}
