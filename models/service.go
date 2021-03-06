package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

const (
	SERVICE_OTHER    = "other"
	SERVICE_MIXPANEL = "mixpanel"
	SERVICE_GA       = "ga"
)

type Service struct {
	BaseModel     `sql:"-"`
	Id            int64
	Name          string
	Type          string
	Configuration string
	UserId        int64
}

func (s *Service) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name": s.Name,
		"type": s.Type,
	}
}

func (s *Service) GetConfiguration() map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(s.Configuration), &result)
	return result
}

func NewService(db *gorm.DB, name string, target string, configuration map[string]interface{}) *Service {
	marshalConfiguration, _ := json.Marshal(configuration)
	service := Service{
		Name:          name,
		Type:          target,
		Configuration: string(marshalConfiguration),
	}
	service.BaseModel = NewBaseModel(db, &service)
	return &service
}

func GetServicesForUser(db *gorm.DB, user *User) []Service {
	var services []Service
	db.Model(user).Related(&services)
	return services
}

func GetServiceFromId(db *gorm.DB, id string) *Service {
	var (
		service Service
		total   int
	)
	db.Where("id = ?", id).First(&service).Count(&total)

	if total == 0 {
		return nil
	}
	service.BaseModel = NewBaseModel(db, &service)
	return &service
}
