package repository

import (
	"database/sql"

	"github.com/fahmiaz411/devcode/modules/todo/interfaces"
	"github.com/fahmiaz411/devcode/modules/todo/repository/mysql"
)

type Repository struct {
	MySQL interfaces.TodoRepoMysql
}

// NewRepository constructor
func NewRepository(mysqlConn *sql.DB) *Repository {
	return &Repository{
		MySQL: mysql.NewMysqlRepository(mysqlConn),
	}
}