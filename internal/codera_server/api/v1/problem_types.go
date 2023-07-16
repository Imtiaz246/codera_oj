package v1

type CreateProblemOption struct {
	Title string `json:"title" validate:"required"`
}
