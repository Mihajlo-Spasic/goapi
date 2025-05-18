package main 

import (
  "fmt"
  "net/http"

  "github.com/go-chi/chi"
  "github.com/mihajlo-spasic/goapi/internal/handlers"
  log "github.com/sirupsen/logrus"
)

func main() {
  log.SetReportCaller(true)
  var r *chi.Mux = chi.NewRouter()
  handlers.Handler(r)

  fmt.Println("Starting GO API service")
  
  if err := http.ListenAndServe("localhost:8080", r); err != nil {
  log.Error(err)

  }
}

