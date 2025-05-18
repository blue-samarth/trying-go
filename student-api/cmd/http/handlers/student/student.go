package student

import (
	"net/http"
	"errors"
	"encoding/json"
	"log/slog"
	"io"
	
	"github.com/blue-samarth/student-api/internal/models"
	"github.com/blue-samarth/student-api/internal/utils/responses"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var student models.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			responses.JsonResponse(w, "error", http.StatusBadRequest, err.Error())
			return 
		}

		slog.Info("Received request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Write([]byte("Welcome to the Student API!"))

		responses.JsonResponse(w, "success", http.StatusCreated, map[string] string {"success": "Student created successfully"})
	}
}