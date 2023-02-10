package entity

type Reaction struct {
	Like    int64 `json:"like" bson:"like"`
	DisLike int64 `json:"dislike" bson:"dislike"`
	Love    int64 `json:"love" bson:"love"`
	Angry   int64 `json:"angry" bson:"angry"`
}

func NewReaction() *Reaction {
	return &Reaction{}
}
