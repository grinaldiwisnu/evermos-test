package controllers

import "go-awesomeapi/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Items Routes
	s.Router.HandleFunc("/item", middlewares.SetMiddlewareJSON(s.CreateItem)).Methods("POST")
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.GetItems)).Methods("GET")
	s.Router.HandleFunc("/item/{id}", middlewares.SetMiddlewareJSON(s.GetItem)).Methods("GET")
	s.Router.HandleFunc("/item", middlewares.SetMiddlewareJSON(s.GetItem)).Methods("PUT")

	// Users Routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")

	// Cart Routes
	s.Router.HandleFunc("/carts/{id}", middlewares.SetMiddlewareJSON(s.GetCartByUID)).Methods("GET")
	s.Router.HandleFunc("/carts", middlewares.SetMiddlewareJSON(s.AddToCart)).Methods("POST")
}
