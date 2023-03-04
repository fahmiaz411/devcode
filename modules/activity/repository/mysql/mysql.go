package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/fahmiaz411/devcode/helper/constant"
	"github.com/fahmiaz411/devcode/modules/activity/domain"
	"github.com/fahmiaz411/devcode/modules/activity/interfaces"
)

type MysqlRepository struct {
	Conn *sql.DB
}

func NewMysqlRepository(Conn *sql.DB) interfaces.ActivityRepoMysql {
	return &MysqlRepository{
		Conn: Conn,
	}
}

func (m *MysqlRepository) Create(ctx context.Context, req domain.ActivityCreateRequest) (res domain.ActivityCreateResponse, err error) {
	now := time.Now().UTC()

	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		INSERT INTO activities (
			title,
			email,
			created_at,
			updated_at
		) VALUES (
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

	// Email nullable
	var email sql.NullString
	if req.Email != constant.EmptyString {
		email.Valid = true
		email.String = req.Email
	}

	var result sql.Result
	result, err = stmt.ExecContext(ctx, req.Title, email, now, now)
	if err != nil {
		return
	}

	res.ID, _ = result.LastInsertId()
	res.Title = req.Title
	res.Email = req.Email
	res.CreatedAt = now
	res.UpdatedAt = now
	
	return
}

func (m *MysqlRepository) Update(ctx context.Context, req domain.ActivityUpdateRequest) (res domain.ActivityUpdateResponse, err error) {
	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		UPDATE activities 
		SET
			title = ?
		WHERE activity_id = ?
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.ExecContext(ctx, req.Title, req.ID)
	
	return
}

func (m *MysqlRepository) Delete(ctx context.Context, req domain.ActivityDeleteRequest) (res domain.ActivityDeleteResponse, err error) {
	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		UPDATE activities SET deleted_at = NOW() WHERE activity_id = ?
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.ExecContext(ctx, req.ID)
	
	return
}

func (m *MysqlRepository) GetAll(ctx context.Context, req domain.ActivityGetAllRequest) (res domain.ActivityGetAllResponse, err error) {
	res = []domain.Activity{}
	
	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		SELECT 
			activity_id,
			title,
			email,
			created_at,
			updated_at,
			deleted_at
		FROM activities
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	var rows *sql.Rows
	rows, err = stmt.QueryContext(ctx)
	if err != nil {
		return
	}

	for rows.Next() {
		var act domain.Activity

		var (
			email sql.NullString
			deletedAt sql.NullTime
		)

		if err = rows.Scan(
			&act.ID,
			&act.Title,
			&email,
			&act.CreatedAt,
			&act.UpdatedAt,
			&deletedAt,
		); err != nil {
			return
		}

		if email.Valid {
			act.Email = email.String
		}

		if deletedAt.Valid {
			act.DeletedAt = &deletedAt.Time
		}

		res = append(res, act)
	}

	return
}

func (m *MysqlRepository) GetOne(ctx context.Context, req domain.ActivityGetOneRequest) (res domain.ActivityGetOneResponse, err error) {
	var stmt *sql.Stmt
	stmt, err = m.Conn.PrepareContext(ctx, `
		SELECT 
			activity_id,
			title,
			email,
			created_at,
			updated_at,
			deleted_at
		FROM activities
		WHERE activity_id = ?
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
		var (
			email sql.NullString
			deletedAt sql.NullTime
		)

		if err = rows.Scan(
			&res.ID,
			&res.Title,
			&email,
			&res.CreatedAt,
			&res.UpdatedAt,
			&deletedAt,
		); err != nil {
			return
		}

		if email.Valid {
			res.Email = email.String
		}

		if deletedAt.Valid {
			res.DeletedAt = &deletedAt.Time
		}
	}

	return
}