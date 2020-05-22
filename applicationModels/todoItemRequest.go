package applicationModels

type TodoItemRequestModel struct {
	Content string
	Status bool
}

type HealthyCheckResponse struct {
	Message string
	Version string
}