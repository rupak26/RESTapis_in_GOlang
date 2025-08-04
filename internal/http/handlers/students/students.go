package students

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/rupak26/RESTapis_in_GOlang/internal/storage"
	"github.com/rupak26/RESTapis_in_GOlang/internal/types"
	"github.com/rupak26/RESTapis_in_GOlang/internal/utils/responses"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func (w http.ResponseWriter , r *http.Request) {
	   if r.Method == "GET" {	
	      w.Write([]byte ("Welcome to student apis"))
	   }
	   if r.Method == "POST" {
		  slog.Info("Creating Student") 
		  var student types.Student 

          decoder := json.NewDecoder(r.Body) 
		  err := decoder.Decode(&student)
          
		  if err != nil {
			responses.WriteJson(w , http.StatusBadRequest , responses.GeneralError(err))
			return 
		  }
        
		  err2 := validator.New().Struct(student)

		  if err2 != nil {
			ValidateErrs := err2.(validator.ValidationErrors)
			responses.WriteJson(w , http.StatusBadRequest , responses.ValidationError(ValidateErrs))
			return 
		  }
          
          lastId , err := storage.CreateStudent(
				student.Name , 
				student.Email ,
				student.Age ,
		  )
          
		  slog.Info("user created successfully" , slog.String("userId" , fmt.Sprint(lastId)))

		  if err != nil {
			responses.WriteJson(w , http.StatusInternalServerError , err)
			return 
		  }

		  responses.WriteJson(w , http.StatusCreated , map[string]int64{"Id" : lastId})
	   }
    }
}


func GetByID(storage storage.Storage) http.HandlerFunc {
	return func (w http.ResponseWriter , r *http.Request) {
		if r.Method != "GET" {	
	      w.Write([]byte ("Method did not Match"))
		  return 
	    }

		id := r.PathValue("id")
		
		slog.Info("Getting a Student" , slog.String("id" , id))

		intId , err := strconv.ParseInt(id , 10 , 64)

		if err != nil {
			responses.WriteJson(w , http.StatusBadRequest , responses.GeneralError(err))
			return 
		}
		student , e := storage.GetStudentById(intId)
		
		if e != nil {
			slog.Error("error getting user" , slog.String("id" , id))
            responses.WriteJson(w , http.StatusInternalServerError , responses.GeneralError(err))
			return 
		}
		responses.WriteJson(w , http.StatusOK , student)
	}
}