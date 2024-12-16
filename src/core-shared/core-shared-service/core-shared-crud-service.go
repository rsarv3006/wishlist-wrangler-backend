package coresharedservice

import (
	"fmt"
	"reflect"
	coresharedexceptions "wishlist-wrangler-api/src/core-shared/core-shared-exceptions"
	coresharedlogger "wishlist-wrangler-api/src/core-shared/core-shared-logger"
	coresharedrepository "wishlist-wrangler-api/src/core-shared/core-shared-repository"
)

type CrudService[T any, ID comparable, DTO any] struct {
	repository *coresharedrepository.CrudRepository[T, ID]
	logger     coresharedlogger.Logger
}

func NewCrudService[T any, ID comparable, DTO any](repository *coresharedrepository.CrudRepository[T, ID], logger coresharedlogger.Logger) *CrudService[T, ID, DTO] {
	return &CrudService[T, ID, DTO]{repository: repository, logger: logger}
}

func (s *CrudService[T, ID, DTO]) GetAll(opts *coresharedrepository.FindOptions) ([]T, error) {
	s.logger.Info("Fetching all entities", coresharedlogger.LogField{Key: "options", Value: opts})
	entities, err := s.repository.FindAll(opts)
	if err != nil {
		s.logger.Error("Failed to fetch entities", coresharedlogger.LogField{Key: "error", Value: err})
		return nil, err
	}

	result := make([]T, 0, len(entities))
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result, nil
}

func (s *CrudService[T, ID, DTO]) GetById(id ID) (*T, error) {
	entity, err := s.repository.FindById(id)
	if err != nil {
		s.logger.Error("Failed to fetch entity", coresharedlogger.LogField{Key: "id", Value: id}, coresharedlogger.LogField{Key: "error", Value: err})
		return nil, coresharedexceptions.NewEntityNotFoundException()
	}
	return entity, nil
}

func (s *CrudService[T, ID, DTO]) UpdateEntity(id ID, dto *DTO) (*T, error) {
	entity, err := s.convertDTOToEntity(dto)
	if err != nil {
		s.logger.Error("Failed to convert DTO to entity", coresharedlogger.LogField{Key: "error", Value: err})
		return nil, err
	}

	existingEntity, err := s.GetById(id)
	if err != nil {
		s.logger.Error("Entity not found", coresharedlogger.LogField{Key: "id", Value: id}, coresharedlogger.LogField{Key: "error", Value: err})
		return nil, coresharedexceptions.NewEntityNotFoundException()
	}

	// Copy the ID from the existing entity to ensure it is not cleared
	s.CopyID(existingEntity, entity)

	updatedEntity, err := s.repository.Save(entity)
	if err != nil {
		s.logger.Error("Failed to update entity", coresharedlogger.LogField{Key: "error", Value: err})
		return nil, err
	}

	return updatedEntity, nil
}

func (s *CrudService[T, ID, DTO]) CreateEntityFromDTO(dto *DTO) (*T, error) {
	entity, err := s.convertDTOToEntity(dto)
	if err != nil {
		s.logger.Error("Failed to convert DTO to entity", coresharedlogger.LogField{Key: "error", Value: err})
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
