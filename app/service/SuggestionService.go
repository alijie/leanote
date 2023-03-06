package service

import (
	"leanote/app/db"
	"leanote/app/info"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//	. "leanote/app/lea"
	//	"time"
	//	"sort"
)

type SuggestionService struct {
}

// 得到某博客具体信息
func (this *SuggestionService) AddSuggestion(suggestion info.Suggestion) bool {
	if suggestion.Id.IsZero() {
		suggestion.Id = primitive.NewObjectID()
	}
	return db.Insert(db.Suggestions, suggestion)
}
