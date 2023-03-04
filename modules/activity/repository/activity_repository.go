package repository

import (
	"database/sql"

	"github.com/fahmiaz411/devcode/modules/activity/interfaces"
	"github.com/fahmiaz411/devcode/modules/activity/repository/mysql"
)

type Repository struct {
	MySQL interfaces.ActivityRepoMysql
}

// NewRepository constructor
func NewRepository(mysqlConn *sql.DB) *Repository {
	return &Repository{
		MySQL: mysql.NewMysqlRepository(mysqlConn),
	}
}