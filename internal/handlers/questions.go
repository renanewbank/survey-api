package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/renanewbank/survey-api/pkg/jsonschema"
)

func QuestionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		schemaPath, _ := filepath.Abs("api/question_schema.json")
		validationErrors, err := jsonschema.ValidateQuestionJSON(body, schemaPath)
		if err != nil {
			http.Error(w, "Error validating JSON", http.StatusInternalServerError)
			return
		}

		if validationErrors != nil {
			http.Error(w, "Invalid JSON:\n"+joinErrors(validationErrors), http.StatusBadRequest)
			return
		}

		// Aqui entraria a lógica de salvar no banco (futura)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Question validated and created (mock)"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func joinErrors(errors []string) string {
	result := ""
	for _, e := range errors {
		result += "- " + e + "\n"
	}
	return result
}
