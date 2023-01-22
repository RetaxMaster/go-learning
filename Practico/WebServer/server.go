package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {

	// Esto va encadenando Middlewares, la función middleware retorna una HandlerFunc que tiene la lógica y encima se van aladiendo los middlewares
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}

func (s *Server) Listen() error {

	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
