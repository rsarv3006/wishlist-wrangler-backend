package coresharedcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"
	coresharedrepository "wishlist-wrangler-api/src/core-shared/core-shared-repository"

	"github.com/go-chi/chi/v5"
)

type CrudController[K any, E any, D any, S interface {
	GetAll(opts *coresharedrepository.FindOptions) ([]E, error)
	GetById(id K) (*E, error)
	UpdateEntity(id K, dto *D) (*E, error)
	CreateEntityFromDTO(dto *D) (*E, error)
}] struct {
	Service S
}

func (c *CrudController[K, E, D, S]) GetAllEntities(w http.ResponseWriter, r *http.Request) {
	opts := &coresharedrepository.FindOptions{}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			handleError(w, "Invalid limit", http.StatusBadRequest)
			return
		}
		opts.Limit = limit
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			handleError(w, "Invalid offset", http.StatusBadRequest)
			return
		}
		opts.Offset = offset
	}

	if order := r.URL.Query().Get("order"); order != "" {
		opts.Order = order
	}

	entities, err := c.Service.GetAll(opts)
	if err != nil {
		handleError(w, "Failed to get entities", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(entities)
	if err != nil {
		handleError(w, "Failed to encode entities", http.StatusInternalServerError)
		return
	}
}

func (c *CrudController[K, E, D, S]) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	var id K
	switch any(id).(type) {
	case int:
		parsedId, err := strconv.Atoi(idStr)
		if err != nil {
			handleError(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		id = any(parsedId).(K)
	case uint:
		parsedId, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			handleError(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		id = any(uint(parsedId)).(K)
	default:
		handleError(w, "Unsupported ID type", http.StatusBadRequest)
		return
	}

	entity, err := c.Service.GetById(id)
	if err != nil {
		handleError(w, "Entity not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(entity)
	if err != nil {
		handleError(w, "Failed to encode entity", http.StatusInternalServerError)
		return
	}
}

func (c *CrudController[K, E, D, S]) UpdateData(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	var id K
	switch any(id).(type) {
	case int:
		parsedId, err := strconv.Atoi(idStr)
		if err != nil {
			handleError(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		id = any(parsedId).(K)
	case uint:
		parsedId, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			handleError(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		id = any(uint(parsedId)).(K)
	default:
		handleError(w, "Unsupported ID type", http.StatusBadRequest)
		return
	}

	var dto D
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		handleError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedEntity, err := c.Service.UpdateEntity(id, &dto)
	if err != nil {
		handleError(w, "Failed to update entity", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedEntity)
	if err != nil {
		handleError(w, "Failed to encode entity", http.StatusInternalServerError)
		return
	}
}

func (c *CrudController[K, E, D, S]) CreateData(w http.ResponseWriter, r *http.Request) {
	var dto D
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		handleError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdEntity, err := c.Service.CreateEntityFromDTO(&dto)
	if err != nil {
		if isForeignKeyViolation(err) {
			handleError(w, "Foreign key violation", http.StatusBadRequest)
			return
		}
		handleError(w, "Failed to create entity", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createdEntity)
	if err != nil {
		handleError(w, "Failed to encode entity", http.StatusInternalServerError)
		return
	}
}
