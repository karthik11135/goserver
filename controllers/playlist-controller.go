package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/context"
	"github.com/karthik11135/golang-backend/db"
	"github.com/karthik11135/golang-backend/models"
)

func GetMyPlaylists(w http.ResponseWriter, r *http.Request) {
	email := context.Get(r, "email")
	fmt.Println("This is my email", email)

	var userId uint
	db.GetDb().Model(&models.User{}).Where("email = ?", email).Select("id").Scan(&userId)

	fmt.Println("My user Id", userId)

	var playlists []models.Playlist
	db.GetDb().Where("user_id = ?", userId).Find(&playlists)

	json.NewEncoder(w).Encode(playlists)
}

func PostMyPlaylist(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		return
	}

	email := context.Get(r, "email")

	fmt.Println("This is my email", email)

	defer r.Body.Close()

	var myPlaylist models.Playlist
	json.NewDecoder(r.Body).Decode(&myPlaylist)

	var UserId uint

	db.GetDb().Model(&models.User{}).Where("email = ?", email).Select("id").Scan(&UserId)

	var myPlayList = models.Playlist{UserId: UserId, PlayListName: myPlaylist.PlayListName, FavoriteSong: myPlaylist.FavoriteSong, NumberOfSongs: myPlaylist.NumberOfSongs}
	db.GetDb().Create(&myPlayList)

	json.NewEncoder(w).Encode("Successfully created")
}