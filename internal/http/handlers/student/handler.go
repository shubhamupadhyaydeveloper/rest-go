package student

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shubhamupadhyaydeveloper/students-api/utils/response"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post method is valid", http.StatusMethodNotAllowed)
		return
	}

	var data Student
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	if err := validator.New().Struct(data); err != nil {
		validateErrors := err.(validator.ValidationErrors)
		http.Error(w, validateErrors.Error(), http.StatusBadRequest)
		return
	}

	slog.Info("student data received", slog.Any("data", data))

	res := map[string]string{
		"status": "success",
		"name":   data.Name,
		"email":  data.Email,
	}

	response.WriteJSON(w,http.StatusCreated,res)
}
