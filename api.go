package nutshttp

import (
	"encoding/json"
	"net/http"
)

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorReponse struct {
	Error ErrorInfo `json:"error"`
}

func WriteError(w http.ResponseWriter, e APIError) {
	w.WriteHeader(e.Code)
	w.Header().Set("Content-Type", "application/json")

	resp := ErrorReponse{
		Error: ErrorInfo{
			Code:    e.Code,
			Message: e.Message,
		},
	}

	data, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func WriteSucc(w http.ResponseWriter, v interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := map[string]interface{}{
		"data": v,
	}

	data, _ := json.Marshal(resp)

	_, _ = w.Write(data)
}
