package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/sahil3982/students-api/internal/storage"
	"github.com/sahil3982/students-api/internal/types"
	"github.com/sahil3982/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//vailidation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.VailidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		slog.Info("Student created", slog.Int64("lastId", lastId), slog.String("name", student.Name), slog.String("email", student.Email), slog.Int("age", student.Age))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{
			"id": lastId,
		})

	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
 fmt.Println("Get by id called")
}
