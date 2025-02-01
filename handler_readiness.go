package main

import (
	"net/http"
)

// Specific function signature this is the func signature that go expects to be written if you want to define an HTTP Handler in the way that GO standard library expects

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, 200, struct{}{})
}