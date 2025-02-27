package service

import (
	"leanote/app/info"
	//	. "leanote/app/lea"
	"time"

	"leanote/app/db"

	"go.mongodb.org/mongo-driver/bson"
)

const IMAGE_TYPE = 0

type AlbumService struct {
}

// add album
func (this *AlbumService) AddAlbum(album info.Album) bool {
	album.CreatedTime = time.Now()
	album.Type = IMAGE_TYPE
	return db.Insert(db.Albums, album)
}

// get albums
func (this *AlbumService) GetAlbums(userId string) []info.Album {
	albums := []info.Album{}
	db.ListByQ(db.Albums, bson.M{"UserId": db.ObjectIDFromHex(userId)}, &albums)
	return albums
}

// delete album
// presupposition: has no images under this ablum
func (this *AlbumService) DeleteAlbum(userId, albumId string) (bool, string) {
	if db.Count(db.Files, bson.M{"AlbumId": db.ObjectIDFromHex(albumId),
		"UserId": db.ObjectIDFromHex(userId),
	}) == 0 {
		return db.DeleteByIdAndUserId(db.Albums, albumId, userId), ""
	}
	return false, "has images"
}

// update album name
func (this *AlbumService) UpdateAlbum(albumId, userId, name string) bool {
	return db.UpdateByIdAndUserIdField(db.Albums, albumId, userId, "Name", name)
}
