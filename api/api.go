package api

import (
  "encoding/json"
  "net/http"
)

type CoinBalanceParams struct{
  Username string 
}

type CoinBalanceResposne struct{
  // Code status
  Code int

  //Balance
  Balance int64 
}

type Error struct{
  // Code status
  Code int 
  
  // Message associated with the error 
  Message string
}



