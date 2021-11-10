package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/tpranoto/bundelt-server/common"
)

const (
	driverPostgres = "postgres"

	DefaultMaxConn = 30
	DefaultMaxIdle = 10
)

type (
	PostgreSQLStorage interface {
		UsersStorage
		UserGroupsStorage
		GroupStorage
		GroupMessageStorage
		EventStorage
		EventGroupStorage
		UserEventStorage
	}

	postgreSQLStorage struct {
		ctx    context.Context
		db     *sqlx.DB
		logger *log.Logger
	}
)

func NewPostgreSQLStorage(log *log.Logger) PostgreSQLStorage {
	host := common.GetEnv("DBHOST", "localhost")
	port := common.GetEnv("DBPORT", "5432")
	user := common.GetEnv("DBUSER", "postgres")
	password := common.GetEnv("DBPASSWORD", "postgres")
	dbName := common.GetEnv("DBNAME", "bundelt")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	dbConn, err := sqlx.Connect(driverPostgres, dbInfo)
	if err != nil {
		log.Fatalf("failed to connect to %s db on %s:%s, %s", dbName, host, port, err.Error())
	}

	return &postgreSQLStorage{
		ctx:    context.Background(),
		db:     dbConn,
		logger: log,
	}
}
