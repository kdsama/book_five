package entity

type UrlObject struct {
	Name     string `bson:"name" json:"name"`
	Verified bool   `bson:"verified" json:"verified"`
}

func MakeUrlObject(url string) *UrlObject {
	return &UrlObject{url, false}
}

func MakeUrlObjects(urls []string) []UrlObject {
	to_return := []UrlObject{}
	for i := range urls {
		to_return = append(to_return, *MakeUrlObject(urls[i]))
	}
	return to_return
}
