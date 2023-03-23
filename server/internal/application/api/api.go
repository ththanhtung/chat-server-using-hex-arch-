package api

type Application struct {
	hub   Hub
}

func NewApplicaion(hub Hub) *Application {
	return &Application{
		hub: hub,
	}
}