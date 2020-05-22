package applicationModels

type TodoItemRequestModel struct {
	Content string
}

type HealthyCheckResponse struct {
	Message string
	Version string
}