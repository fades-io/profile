package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"strings"
	"userGoServ/iternal/user"
	"userGoServ/pgk/client/postgresql"
	"userGoServ/pgk/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, user *user.User) error {

	q := `
		INSERT INTO public.user 
				(id, username, email, password, level, daysinrow,daysinweek, doessendpushups, theme, language, image)
		VALUES 
		       ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, user.Id, user.Username, user.Email, user.Password, user.Level, user.DaysInRow, user.DaysInWeek, user.DoesSendPushUps, user.Theme, user.Language, user.Image).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) FindOne(ctx context.Context, username string) (user.User, error) {
	q := `
		SELECT id FROM public.user WHERE username = $1 
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var usr user.User
	err := r.client.QueryRow(ctx, q, username).Scan(&usr.Id)
	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

func (r *repository) Update(ctx context.Context, user user.User) error {
	q := `
		UPDATE public.user SET email = $2, password = $3, level = $4, daysinrow = $5, daysinweek = $6,  doessendpushups = $7, theme = $8, language = $9, image = $10 WHERE username = $1 returning image
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	json, _ := json.Marshal(user.DaysInWeek)

	if err := r.client.QueryRow(ctx, q, user.Username, user.Email, user.Password, user.Level, user.DaysInRow, string(json), user.DoesSendPushUps, user.Theme, user.Language, user.Image).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil

}

func (r *repository) Delete(ctx context.Context, username string) error {
	q := `
		DELETE FROM public.user WHERE username = $1 returning id
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	user := user.User{}
	if err := r.client.QueryRow(ctx, q, username).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil

}

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

// Do we need it?
func (r *repository) FindAll(ctx context.Context) (u []user.User, err error) {
	q := `
		SELECT id, username FROM public.user;
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]user.User, 0)

	for rows.Next() {
		var usr user.User

		err = rows.Scan(&usr.Id, &usr.DaysInRow)
		if err != nil {
			return nil, err
		}

		users = append(users, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
