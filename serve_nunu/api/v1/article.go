package v1

type ArticleWithTags struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  uint   `json:"userId"`
	Tags    []int  `json:"tags"`
}
