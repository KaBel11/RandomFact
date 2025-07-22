package api

import (
	"net/http"
	"strconv"

	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/service"
	"github.com/KaBel11/RandomFact/shared/utils"
)

type FactsHandler struct {
	svc *service.FactsService
}

func NewFactsHandler(svc *service.FactsService) *FactsHandler {
	return &FactsHandler{svc: svc}
}

func (h *FactsHandler) List(w http.ResponseWriter, r *http.Request) {
	facts, err := h.svc.GetAllFacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, facts)
}

func (h *FactsHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fact, err := h.svc.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, fact)
}

func (h *FactsHandler) GetRandomFact(w http.ResponseWriter, r *http.Request) {
	fact, err := h.svc.GetRandomFact()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, fact)
}

func (h *FactsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateFactRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fact, err := h.svc.CreateFact(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, fact)
}

func (h *FactsHandler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseUint(idParam, 10, 64) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req dtos.UpdateFactRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id != req.ID {
		http.Error(w, "Ids doesn't match", http.StatusBadRequest)
		return
	}

	fact, err := h.svc.UpdateFact(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, fact)
}

func (h *FactsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseUint(idParam, 10, 64) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.svc.DeleteFact(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusNoContent, nil)
}
