package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/fahmiaz411/devcode/helper/constant"
	"github.com/fahmiaz411/devcode/modules/todo/domain"
	"github.com/fahmiaz411/devcode/modules/todo/interfaces"
)

type MysqlRepository struct {
	Conn *sql.DB
}

func NewMysqlRepository(Conn *sql.DB) interfaces.TodoRepoMysql {
	return &MysqlRepository{
		Conn: Conn,
	}
}

func (m *MysqlRepository) Create(ctx context.Context, req domain.TodoCreateRequest) (res domain.TodoCreateResponse, err error) {
	now := time.Now().UTC()

	var isActive sql.NullBool
	if req.IsActive != nil {
		isActive.Valid = true
		isActive.Bool = *req.IsActive
	}

	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		INSERT INTO todos (
			title,
			activity_group_id,
			is_active,
			created_at,
			updated_at
		) VALUES (
			?,
			?,
			?,
			?,
			?
		)
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	values := []any{
		req.Title,
		req.ActivityGroupID,
		isActive,
		now,
		now,
	}

	var result sql.Result
	result, err = stmt.ExecContext(ctx, values...)
	if err != nil {
		return
	}

	res.ID, _ = result.LastInsertId()
	res.Title = req.Title
	res.ActivityGroupID = req.ActivityGroupID
	res.Priority = domain.PriorityDefault
	res.CreatedAt = now
	res.UpdatedAt = now

	if isActive.Valid {
		res.IsActive = *req.IsActive
	} else {
		res.IsActive = domain.IsActiveDefault
	}
	
	return
}

func (m *MysqlRepository) Update(ctx context.Context, req domain.TodoUpdateRequest) (res domain.TodoUpdateResponse, err error) {
	fields := []string{}
	values := []any{}

	if req.Title != constant.EmptyString {
		fields = append(fields, "title")
		values = append(values, req.Title)
	}

	if req.IsActive != nil {
		fields = append(fields, "is_active")
		values = append(values, *req.IsActive)
	}

	if req.Priority != constant.EmptyString {
		fields = append(fields, "priority")
		values = append(values, req.Priority)
	}

	if len(fields) == constant.ZeroValue {
		return
	}

	// Updated At
	fields = append(fields, "updated_at")
	values = append(values, req.UpdatedAt)

	// Id
	values = append(values, req.ID)

	for key, field := range fields {
		fields[key] = field + " = ?"
	}

	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, fmt.Sprintf(`
		UPDATE todos 
		SET
			%s			
		WHERE todo_id = ?
	`, strings.Join(fields, ", ")))
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.ExecContext(ctx, values...)
	
	return
}

func (m *MysqlRepository) Delete(ctx context.Context, req domain.TodoDeleteRequest) (res domain.TodoDeleteResponse, err error) {
	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		DELETE FROM todos WHERE todo_id = ?
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.ExecContext(ctx, req.ID)
	
	return
}

func (m *MysqlRepository) GetAll(ctx context.Context, req domain.TodoGetAllRequest) (res domain.TodoGetAllResponse, err error) {
	res = []domain.Todo{}
	
	values := []any{}

	var queryWhere string
	if req.ActivityGroupID != int64(constant.ZeroValue) {
		queryWhere = fmt.Sprintf(`WHERE activity_group_id = ?`)
		values = append(values, req.ActivityGroupID)
	}

	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, fmt.Sprintf(`
		SELECT 
			todo_id,
			activity_group_id,
			title,
			is_active,
			priority,
			created_at,
			updated_at
		FROM todos
		%s
	`, queryWhere))
	if err != nil {
		return
	}
	defer stmt.Close()

	var rows *sql.Rows
	rows, err = stmt.QueryContext(ctx, values...)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo domain.Todo

		var isActive sql.NullBool

		if err = rows.Scan(
			&todo.ID,
			&todo.ActivityGroupID,
			&todo.Title,
			&isActive,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		); err != nil {
			return
		}

		if isActive.Valid {
			todo.IsActive = isActive.Bool
		}

		res = append(res, todo)
	}

	return
}

func (m *MysqlRepository) GetOne(ctx context.Context, req domain.TodoGetOneRequest) (res domain.TodoGetOneResponse, err error) {
	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		SELECT 
			todo_id,
			activity_group_id,
			title,
			is_active,
			priority,
			created_at,
			updated_at
		FROM todos
		WHERE todo_id = ?
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	var rows *sql.Rows
	rows, err = stmt.QueryContext(ctx, req.ID)
	if err != nil {
		return
	}

	if rows.Next() {
		var isActive sql.NullBool

		if err = rows.Scan(
			&res.ID,
			&res.ActivityGroupID,
			&res.Title,
			&isActive,
			&res.Priority,
			&res.CreatedAt,
			&res.UpdatedAt,
		); err != nil {
			return
		}

		if isActive.Valid {
			res.IsActive = isActive.Bool
		}
	}

	return
}