package models

type Submission struct {
	Id         int64  `json:"id"`
	Title      string `json:"tittle"`
	AuthorName string `json:"author_name"`
	Email      string `json:"email"`
	Status     string `json:"status"`
}
