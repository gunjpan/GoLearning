package models

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler
    handler = route.HandlerFunc

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }

  return router
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },

  Route{
    "CreateLot",
    "POST",
    "/lots",
    CreateLot,
  },

  Route{
    "DeleteLot",
    "DELETE",
    "/lots/{name}",
    DeleteLot,
  },

  Route{
    "FindLot",
    "GET",
    "/lots/{name}",
    FindLot,
  },

  Route{
    "ParkCar",
    "POST",
    "/lots/{name}/enter",
    ParkCar,
  },

  Route{
    "UnparkCar",
    "POST",
    "/lots/{name}/exit",
    UnparkCar,
  },

  Route{
    "UpdateLot",
    "PUT",
    "/lots/{name}",
    UpdateLot,
  },
}
