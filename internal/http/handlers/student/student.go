package student

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	types "github.com/Anurag340/student-api/internal/types"
	response "github.com/Anurag340/student-api/internal/utils"
	"github.com/Anurag340/student-api/storage"
	validator "github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if err!=nil {
			response.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "empty request body"})
			return
		}


		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSONResponse(w, http.StatusBadRequest, response.ValidationErrors(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJSONResponse(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJSONResponse(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
			return
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Error("error getting user", slog.String("id", id))
			response.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string {"error": "internal server error"})
			return
		}

		response.WriteJSONResponse(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all students")

		students, err := storage.GetStudents()
		if err != nil {
			response.WriteJSONResponse(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJSONResponse(w, http.StatusOK, students)
	}
}

func UpdateStudent(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id:=r.PathValue("id")
		slog.Info("updating a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
			return
		}

		var student types.Student
		err = json.NewDecoder(r.Body).Decode(&student)

		if err != nil {
			response.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "empty request body"})
			return
		}

		err = storage.UpdateStudent(intId , student.Name , student.Email , student.Age)
		if err != nil {
			response.WriteJSONResponse(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "student updated successfully"})
		slog.Info("user updated successfully", slog.String("userId", fmt.Sprint(student.Id)))
	}
}

func DeleteStudent(storage storage.Storage) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		
		id := r.PathValue("id")

		intId ,err := strconv.ParseInt(id, 10, 64)

		if err!=nil{
			response.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
			return
		}

		slog.Info("Deleting student",intId)

		err = storage.DeleteStudent(intId)
		if err!=nil{
			slog.Info("error deleting student",intId)
			response.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}
		response.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "student deleted successfully"})
		slog.Info("student deleted successfully", slog.String("id", id))

	}

}