package db

import (
	"errors"
	"strings"

	"github.com/jak103/usu-gdsf/config"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

const (
	MOCK      = "mock"
	FIRESTORE = "firestore"
	MONGO     = "mongo"
)

var connection Database

type Database interface {
	GetAllGames() ([]models.Game, error)
	AddGame(models.Game) (string, error)
	RemoveGame(models.Game) error
	GetGameByID(string) (models.Game, error)
	GetGamesByTag(string) ([]models.Game, error)
	GetGameID(models.Game) (string, error)
	Disconnect() error
	Connect() error
}

func NewDatabaseFromEnv() (Database, error) {
	if connection == nil {
		dbType := strings.ToLower(config.DbType)

		switch dbType {
		case "mock":
			connection = &Mock{}
		case "firestore":
			connection = &Firestore{}
		case "mongo":
			connection = &Mongo{}

		default:
			log.Error("Unknown DB_TYPE %v", dbType)
			return nil, errors.New("unknown DB_TYPE")
		}

		err := connection.Connect()
		if err != nil {
			log.WithError(err).Error("Could not connect to database")
			return nil, err
		}
	}
	return connection, nil
}
