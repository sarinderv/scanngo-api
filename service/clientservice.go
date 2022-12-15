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
	ClientID         int          `json:"clientId" db:"id"`
	Name             string       `json:"name" db:"name"`
	Address          string       `json:"address" db:"address"`
	Active           bool         `json:"active" db:"active"`
	Primary_Contact  string       `json:"primaryContact" db:"primary_contact"`
	DeactivationDate sql.NullTime `json:"deactivationDate" db:"deactivationDate"`
	Client_image     string       `json:"clientImage" db:"client_image"`
	Logo_url         string       `json:"logo_url" db:"logo_url"`
	Insert_Timestamp sql.NullTime `json:"ins_ts" db:"ins_ts"`
	Update_Timestamp sql.NullTime `json:"upd_ts" db:"upd_ts"`
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

	rs, err := sqlx.NamedExec(tx, "INSERT INTO client (name, address, active,client_image, logo_url, primary_contact) VALUES (:name, :address, :active, :client_image, :logo_url , :primary_contact)", r)

	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	id, _ := rs.LastInsertId()

	r.ClientID = int(id)

	fmt.Println(rs)

	return r, nil
}

func (s *clientService) Update(ctx context.Context, tx *sqlx.Tx, r *Client) error {
	log.Info().Msg("Inserting record into client")

	rs, err := sqlx.NamedExec(tx, "UPDATE client SET name=:name , address=:address, active=:active, deactivationDate=:deactivationDate,client_image=:client_image, logo_url=:logo_url, primary_contact=:primary_contact WHERE id = :id", r)

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
