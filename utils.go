//Misc helper methods that don't really belong anywhere else...
package main

import (
	"errors"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//VerifyBody is a helper function to ensure all http requests contain the requisite fields returns error if fields missing
func (h *Handler) VerifyBody(body map[string]interface{}, str ...string) error {
	for _, s := range str {
		fmt.Println(s)
		if _, ok := body[s]; !ok {
			return errors.New("error: missing field: " + s)
		}
	}

	return nil
}
