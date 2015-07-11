package parsesearch

import (
	"encoding/json"
	"io"
	"log"
)

// Response is a Parse Cloud Code Webhook response
type Response struct {
	Error   interface{} `json:"error,omitempty"`
	Success interface{} `json:"success,omitempty"`
}

func writeErr(w io.Writer, msg error) {
	err := json.NewEncoder(w).Encode(Response{Error: msg.Error()})
	log.Println("wrote error:", err)
	if err != nil {
		log.Println("error encoding response:", err)
	}
}
