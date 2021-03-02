package service



type Country struct {
	Code  string
	Regex string
}

type CategoryListResponse struct {
	Count int      `json:"count"`
}
