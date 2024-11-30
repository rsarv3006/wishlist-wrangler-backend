package service

import (
	"fmt"
	"reflect"
	"wishlist-wrangler-api/exceptions"
	"wishlist-wrangler-api/logger"
	"wishlist-wrangler-api/repository"
)

type CrudService[T any, ID comparable, DTO any] struct {
	repository *repository.CrudRepository[T, ID]
	logger     logger.Logger
}

func NewCrudService[T any, ID comparable, DTO any](repository *repository.CrudRepository[T, ID], logger logger.Logger) *CrudService[T, ID, DTO] {
	return &CrudService[T, ID, DTO]{repository: repository, logger: logger}
}

func (s *CrudService[T, ID, DTO]) GetAll(opts *repository.FindOptions) ([]T, error) {
	s.logger.Info("Fetching all entities", logger.LogField{Key: "options", Value: opts})
	entities, err := s.repository.FindAll(opts)
	if err != nil {
		s.logger.Error("Failed to fetch entities", logger.LogField{Key: "error", Value: err})
		return nil, err
	}

	result := make([]T, len(entities))
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result, nil
}

func (s *CrudService[T, ID, DTO]) GetById(id ID) (*T, error) {
	entity, err := s.repository.FindById(id)
	if err != nil {
		s.logger.Error("Failed to fetch entity", logger.LogField{Key: "id", Value: id}, logger.LogField{Key: "error", Value: err})
		return nil, exceptions.NewEntityNotFoundException()
	}
	return entity, nil
}

func (s *CrudService[T, ID, DTO]) UpdateEntity(id ID, dto *DTO) (*T, error) {
	entity, err := s.convertDTOToEntity(dto)
	if err != nil {
		s.logger.Error("Failed to convert DTO to entity", logger.LogField{Key: "error", Value: err})
		return nil, err
	}

	existingEntity, err := s.GetById(id)
	if err != nil {
		s.logger.Error("Entity not found", logger.LogField{Key: "id", Value: id}, logger.LogField{Key: "error", Value: err})
		return nil, exceptions.NewEntityNotFoundException()
	}

	// Copy the ID from the existing entity to ensure it is not cleared
	s.CopyID(existingEntity, entity)

	updatedEntity, err := s.repository.Save(entity)
	if err != nil {
		s.logger.Error("Failed to update entity", logger.LogField{Key: "error", Value: err})
		return nil, err
	}

	return updatedEntity, nil
}

func (s *CrudService[T, ID, DTO]) CreateEntityFromDTO(dto *DTO) (*T, error) {
	entity, err := s.convertDTOToEntity(dto)
	if err != nil {
		s.logger.Error("Failed to convert DTO to entity", logger.LogField{Key: "error", Value: err})
		return nil, err
	}
	s.ClearID(entity)
	return s.repository.Save(entity)
}

func (s *CrudService[T, ID, DTO]) CopyID(entity *T, dto any) {
	if identifiableEntity, ok := any(entity).(interface{ GetID() ID }); ok {
		if identifiableDto, ok := any(dto).(interface{ SetID(ID) }); ok {
			identifiableDto.SetID(identifiableEntity.GetID())
		}
	}
}

func (s *CrudService[T, ID, DTO]) ClearID(entity *T) {
	if identifiable, ok := any(entity).(interface{ ClearID() }); ok {
		identifiable.ClearID()
	}
}

func (s *CrudService[T, ID, DTO]) convertDTOToEntity(dto *DTO) (*T, error) {
	if dto == nil {
		s.logger.Error("DTO cannot be nil")
		return nil, fmt.Errorf("dto cannot be nil")
	}

	if converter, ok := any(dto).(interface{ ToEntity() (*T, error) }); ok {
		return converter.ToEntity()
	}

	entity := new(T)
	dtoValue := reflect.ValueOf(dto).Elem()
	entityValue := reflect.ValueOf(entity).Elem()

	if dtoValue.Kind() != reflect.Struct || entityValue.Kind() != reflect.Struct {
		s.logger.Error("Both DTO and entity must be structs")
		return nil, fmt.Errorf("both DTO and entity must be structs")
	}

	for i := 0; i < dtoValue.NumField(); i++ {
		dtoField := dtoValue.Field(i)
		dtoFieldName := dtoValue.Type().Field(i).Name
		entityField := entityValue.FieldByName(dtoFieldName)

		if !entityField.IsValid() {
			continue // Skip fields that don't exist in entity
		}

		if !entityField.CanSet() {
			continue // Skip unexported fields
		}

		if dtoField.Type() != entityField.Type() {
			continue // Skip fields with different types
		}

		entityField.Set(dtoField)
	}

	return entity, nil
}
