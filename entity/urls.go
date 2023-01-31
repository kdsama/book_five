package entity

type UrlObject struct {
	Name     string `bson:"name" json:"name"`
	Verified bool   `bson:"verified" json:"verified"`
}

func MakeUrlObject(url string) *UrlObject {
	return &UrlObject{url, false}
}
