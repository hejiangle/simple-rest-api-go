package application_models

type TodoItemRequestModel struct {
	Content string
	Status  bool
}

type HealthyCheckResponse struct {
	Message string
	Version string
}
