package startup

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func DatabaseConnect(config DatabaseConfig) (*gorm.DB, error) {

	//dsn := fmt.Sprintf("%s:%s://%s:%s@postgres/%s?sslmode=disable", config.Host, config.Port,
	//	config.User, config.Password, config.DBName)
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	return nil, err
	//}
	//return db, nil
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.User, config.DBName, config.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
