package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/renanewbank/survey-api/internal/models"
	"github.com/renanewbank/survey-api/internal/repository"
	"github.com/renanewbank/survey-api/pkg/jsonschema"
)

type QuestionHandler struct {
	repo       *repository.QuestionRepository
	schemaPath string
}

func NewQuestionHandler(repo *repository.QuestionRepository) *QuestionHandler {
	schemaPath, _ := filepath.Abs("api/question_schema.json")
	return &QuestionHandler{repo: repo, schemaPath: schemaPath}
}

func (h *QuestionHandler) HandleListOrCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		questions := h.repo.GetAll()
		json.NewEncoder(w).Encode(questions)

	case http.MethodPost:
		body, _ := io.ReadAll(r.Body)
		errors, err := jsonschema.ValidateQuestionJSON(body, h.schemaPath)
		if err != nil {
			http.Error(w, "Schema validation error", http.StatusInternalServerError)
			return
		}
		if errors != nil {
			http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
			return
		}

		var q models.Question
		if err := json.Unmarshal(body, &q); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		created := h.repo.Create(q)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *QuestionHandler) HandleByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/questions/")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		q, err := h.repo.GetByID(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(q)

	case http.MethodPut:
		body, _ := io.ReadAll(r.Body)
		errors, err := jsonschema.ValidateQuestionJSON(body, h.schemaPath)
		if err != nil {
			http.Error(w, "Schema validation error", http.StatusInternalServerError)
			return
		}
		if errors != nil {
			http.Error(w, strings.Join(errors, "\n"), http.StatusBadRequest)
			return
		}

		var updated models.Question
		if err := json.Unmarshal(body, &updated); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		q, err := h.repo.Update(id, updated)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(q)

	case http.MethodDelete:
		err := h.repo.Delete(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
