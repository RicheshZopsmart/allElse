package Net_rest

import (
	"fmt"
	"net/http"
)

func GetAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{
		{"Request-Type": "%v"},
		{"Request-User Agent": "%v"},
	}`, r.Method, r.UserAgent())))

}
