package service

import (
"context"
"database/sql"
"errors"
"fmt"

"github.com/jmoiron/sqlx"
"github.com/rs/zerolog/log"
)

type Client struct {
ID               int           `json:"id" db:"id"`
Name             string        `json:"name" db:"name"`
Address          string        `json:"address" db:"address"`
Active           bool          `json:"active" db:"active"`
DeactivationDate sql.NullTime  `json:"deactivationDate" db:"deactivationDate"`
Insert_Timestamp sql.NullTime  `json:"ins_ts" db:"ins_ts"`
Update_Timestamp sql.NullTime  `json:"upd_ts" db:"upd_ts"`
Insert_User      int           `db:"ins_user"`
Update_User      sql.NullInt32 `db:"upd_user"`
}

type IClientService interface {
Create(ctx context.Context, tx *sqlx.Tx, r *Client) (*Client, error)
Update(ctx context.Context, tx *sqlx.Tx, r *Client) error
Delete(ctx context.Context, tx *sqlx.Tx, id int) error
FindAll(ctx context.Context) ([]*Client, error)
FindByID(ctx context.Context, id int) (*Client, error)
}

type clientService struct {
db *sqlx.DB
}

func NewClientService(db *sqlx.DB) IClientService {
return &clientService{
	db,
}
}

func (s *clientService) Create(ctx context.Context, tx *sqlx.Tx, r *Client) (*Client, error) {
log.Info().Msg("Inserting record into client")

rs, err := sqlx.NamedExec(tx, "INSERT INTO client (name, address, active , ins_user) VALUES (:name, :address, :active, :ins_user )", r)

if err != nil {
	log.Error().Msg(err.Error())
	return nil, err
}

id, _ := rs.LastInsertId()

r.ID = int(id)

fmt.Println(rs)

return r, nil
}

func (s *clientService) Update(ctx context.Context, tx *sqlx.Tx, r *Client) error {
log.Info().Msg("Inserting record into client")

rs, err := sqlx.NamedExec(tx, "UPDATE client SET name=:name , address=:address, active=:active, deactivationDate=:deactivationDate, upd_user=:upd_user WHERE id = :id", r)

if err != nil {
	log.Error().Msg(err.Error())
	return err
}

fmt.Println(rs)
return nil
}

func (s *clientService) Delete(ctx context.Context, tx *sqlx.Tx, id int) (err error) {
log.Info().Msg("Deleting record from client")

rs := sqlx.MustExec(tx, "DELETE FROM client WHERE id = ?", id)

rows, err := rs.RowsAffected()
if rows == 0 {

	return errors.New("asdasd")
}

if err != nil {

	return err
}

fmt.Println(rs)
return nil
}

func (s *clientService) FindAll(ctx context.Context) (c []*Client, err error) {
log.Info().Msg("Inserting record into client")

people := []*Client{}
err = s.db.Select(&people, "SELECT * FROM client where id > 1")
if err != nil {
	log.Error().Msg(err.Error())
	return nil, err
}
return people, nil
}

func (s *clientService) FindByID(ctx context.Context, id int) (c *Client, err error) {
log.Info().Msg("Getting client by id")

people := Client{}
err = s.db.Get(&people, "SELECT * FROM client where id=? LIMIT 1", id)
if err != nil {
	return nil, err
}
return &people, nil
}
