package models

type User struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name string `json:"username"` 
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Playlists []Playlist
}

type Playlist struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId uint
	PlayListName string `json:"plName"`
	NumberOfSongs int `json:"noSongs"`
	FavoriteSong string `json:"favSong"`
}

