package handler

import (
	"net/http"
	"strconv"

	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/service"
	"github.com/KaBel11/RandomFact/shared/utils"
)

type FactHandler struct {
	svc *service.FactService
}

func NewFactHandler(svc *service.FactService) *FactHandler {
	return &FactHandler{svc: svc}
}

func (h *FactHandler) List(w http.ResponseWriter, r *http.Request) {
	facts, err := h.svc.GetAllFacts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, facts)
}

func (h *FactHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fact, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, fact)
}

func (h *FactHandler) GetRandomFact(w http.ResponseWriter, r *http.Request) {
	fact, err := h.svc.GetRandomFact(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, fact)
}

func (h *FactHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateFactRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fact, err := h.svc.CreateFact(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, fact)
}

func (h *FactHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	fact, err := h.svc.UpdateFact(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, fact)
}

func (h *FactHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseUint(idParam, 10, 64) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.svc.DeleteFact(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusNoContent, nil)
}