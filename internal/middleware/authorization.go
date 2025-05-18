package middleware

import (
  "erorrs"
  "net/http"

  "github.com/mihajlo-spasic/goapi/api"
  "github.com/mihajlo-spasic/goapi/internals/tools"
  log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username token")

func Authrorization(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    var username string = r.URL.query().get("username")  
    var token = r.Header.get("Authrorization")
    var err error

    if username == "" || token = "" {
      log.Error(UnAuthorizedError)
      api.RequestErrorHandler(w, UnAuthorizedError)
      return 
    }

    var database *tools.DatabaseInterface
    database, err = tools.NewDatabase()
    if err != nil { 
      api.InternalErrorHandler(w)
      return
    }

    var loginDetails *tools.loginDetails
    loginDetails = (*database).GetUserLoginDetails(username)

    if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
      log.Error(UnAuthorizedError)
      api.RequestErrorHandler(w, UnAuthorizedError)
      return 
    }

    next.ServeHTTP(w ,r)
  })
}
