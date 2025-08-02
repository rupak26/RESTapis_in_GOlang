package students

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/rupak26/RESTapis_in_GOlang/internal/types"
	"github.com/rupak26/RESTapis_in_GOlang/internal/utils/responses"
)

func New() http.HandlerFunc {
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

		  responses.WriteJson(w , http.StatusCreated , map[string]string{"sucess" : "OK"})
	   }
    }
}