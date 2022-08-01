package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Invj(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["ip"] = r.RemoteAddr
	resp["user-agent"] = r.UserAgent()
	resp["accept-language"] = r.Header.Get("Accept-Language")
	resp["message"] = "Inverse JC bot...starting"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error in JSON marshall.\n%s", err)
	} else {
		w.Write(jsonResp)
	}
	return
}
