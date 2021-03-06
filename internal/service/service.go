package service

import (
	"WB-test-L0/internal/domain/model"
	"WB-test-L0/internal/repository/cache"
	"WB-test-L0/internal/repository/database"
	"encoding/json"
	"log"
)

type userService struct {
	cache cache.Cache
	repo  database.Repository
}

//NewUserService - create new Service interface reference
func NewUserService(cache cache.Cache, repo database.Repository) Service {
	return &userService{cache: cache, repo: repo}
}

//FindByUUID - take entity from cache
func (s *userService) FindByUUID(uuid string) (model.Entity, error) {
	//return from cache
	return s.cache.Get(uuid), nil
}

//SetEntity - set entity to cache + repository
func (s *userService) SetEntity(message []byte) error {
	entity := model.Entity{}
	//parse message
	err := json.Unmarshal(message, &entity)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	//set to cache & database message
	s.cache.Set(entity.OrderUID, entity)
	err = s.repo.CreateEntity(entity.OrderUID, message)
	if err != nil {
		return err
	}

	return nil
}

//GetAllEntity - take all entities from repo to cache
func (s *userService) GetAllEntity() error {
	rows, err := s.repo.FindAllEntities()
	if err != nil {
		return err
	}

	var uuid string
	var data []byte
	var entity model.Entity

	for rows.Next() {
		err = rows.Scan(&uuid, &data)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &entity)
		if err != nil {
			return err
		}

		s.cache.Set(uuid, entity)
	}

	return nil
}
