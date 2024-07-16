package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
// 	w.WriteHeader(statusCode)
// 	err := json.NewEncoder(w).Encode(data)
// 	if err != nil {
// 		fmt.Fprintf(w, "%s", err.Error())
// 	}
// }

func JSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, "An error occurred", struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, "Bad Request", nil)
}