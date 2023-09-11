package application

import v1 "agahi/internal/delivery/v1"

func (a *App) RegisterRoutes() {
	a.Router.HandleFunc("/register", v1.RegisterUserAction(a.Repository.UserRepo)).Methods("POST")
}
