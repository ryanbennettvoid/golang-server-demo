package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/cors"
)

func (server *Server) Listen() error {

  r := chi.NewRouter()

  cors := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  })

  r.Use(cors.Handler)
  r.Use(middleware.StripSlashes)
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/heartbeat", server.HandleGetHeartbeat)
  r.Get("/members", server.HandleGetMembers)
  r.Post("/members", server.HandleCreateMember)
  r.Get("/members/{id}", server.HandleGetMemberById)
  r.Put("/members/{id}", server.HandleUpdateMemberById)
  r.Delete("/members/{id}", server.HandleDeleteMemberById)

  fmt.Printf("listening on port %d...\n", server.Port)

  return http.ListenAndServe(fmt.Sprintf(":%d", server.Port), r)

}
