package utils

import (
	"encoding/json"
	"net/http"
)

/*
*	Returns HTTP Response payload
*		@params status : bool -> HTTP Status
*		@params message : string -> HTTP Response's message
*		@return interface : map[string]interface{} -> Payload
*/
func Message(status bool, message string) map[string]interface{} {
	// Set response's payload
	return map[string]interface{} {"status" : status, "message": message}
}

/*
*	Receives a payload and serializes it into a JSON
*		@params w : http.ResponseWriter -> Util for HTTP responses
*		@params payload : map[string]interface{} -> response's payload
*/
func Respond(w http.ResponseWriter, payload map[string]interface{}) {
	// Set headers app content-type
	w.Header().Add("Content-Type", "application/json")
	// Encode payload to JSON
	json.NewEncoder(w).Encode(payload)
}