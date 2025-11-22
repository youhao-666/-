package v1

type Tag struct {
	ID      int    `json:"id"`
	TagName string `json:"tagName"`
}

type Tags struct {
	Tags []Tag `json:"tags"`
}

type TagCreate struct {
	TagName string `json:"tagName"`
}
