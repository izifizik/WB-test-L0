package service

import (
	"WB-test-L0/internal/domain/model"
	"WB-test-L0/internal/errors"
	"WB-test-L0/internal/repository/cache"
	"WB-test-L0/internal/repository/database"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

type userService struct {
	cache cache.Cache
	repo  database.Repository
	conn  stan.Conn
}

const sub = "WB"

func NewUserService(cache cache.Cache, repo database.Repository, conn stan.Conn) Service {
	return &userService{cache: cache, repo: repo, conn: conn}
}

func (s *userService) FindByUUID(uuid string) (model.Entity, error) {
	//return from cache
	return s.cache.Get(uuid), nil
}

func (s *userService) SetEntity(message []byte) error {
	entity := model.Entity{}
	//parse message
	err := json.Unmarshal(message, &entity)
	if err != nil {
		log.Println(err.Error())
		return errors.ServiceError
	}

	//set to cache & database message
	s.cache.Set(entity.OrderUID, entity)
	err = s.repo.CreateEntity(entity.OrderUID, message)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *userService) GetAllEntity() error {
	//here we get all entity from db to cache
	rows, err := s.repo.FindAllEntity()
	if err != nil {
		return err
	}

	var uuid string
	var data []byte
	var entity model.Entity
	//set all data to cache
	for rows.Next() {
		err = rows.Scan(&uuid, &data)
		if err != nil {
			return errors.DbError
		}

		err = json.Unmarshal(data, &entity)
		if err != nil {
			return errors.ServiceError
		}

		s.cache.Set(uuid, entity)
	}

	return nil
}

func (s *userService) SetToNats(message []byte) error {
	err := s.conn.Publish(sub, message)
	if err != nil {
		return err
	}

	return nil
}
