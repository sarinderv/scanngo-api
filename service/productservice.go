package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Product struct {
	ID               int          `json:"id" db:"id"`
	Name             string       `json:"name" db:"name"`
	Price            string       `json:"price" db:"price"`
	Active           bool         `json:"active" db:"active"`
	ClientID         int          `json:"clientId" db:"client_id"`
	ImageUrl         string       `json:"imageUrl" db:"image_url"`
	Insert_Timestamp sql.NullTime `json:"ins_ts" db:"ins_ts"`
	Update_Timestamp sql.NullTime `json:"upd_ts" db:"upd_ts"`
}

type IProductService interface {
	Create(ctx context.Context, tx *sqlx.Tx, r []*Product) error
	Update(ctx context.Context, tx *sqlx.Tx, r *Product, clientID int) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
	FindAll(ctx context.Context, clientID int) ([]*Product, error)
	FindByID(ctx context.Context, id int, clientID int) (*Product, error)
	FindByName(ctx context.Context, name string, clientID int) ([]Product, error)
}

type ProductService struct {
	db  *sqlx.DB
	log zerolog.Logger
}

func NewProductService(db *sqlx.DB, log zerolog.Logger) IProductService {
	return &ProductService{
		db,
		log,
	}
}

func (s *ProductService) Create(ctx context.Context, tx *sqlx.Tx, r []*Product) error {
	log.Info().Msg("Inserting record into inventory")

	rs, err := sqlx.NamedExec(tx, "INSERT INTO inventory (name, price, active, client_id, image_url) VALUES (:name, :price, :active, :client_id, :image_url )", r)

	if err != nil {
		return err
	}

	rows, err := rs.RowsAffected()

	if rows < 1 {
		if err != nil {
			err := errors.New("no rows were inserted")
			log.Error().Msg(err.Error())
			return err
		}
	}

	return nil
}

func (s *ProductService) Update(ctx context.Context, tx *sqlx.Tx, r *Product, clientID int) error {
	s.log.Debug().Msg("updating record into inventory")

	product, err := getProduct(s.log, tx, r.ID, clientID)
	if err != nil {
		return err
	}

	s.log.Debug().Msg("updating product from base")
	updateProduct(product, r)

	rs, err := sqlx.NamedExec(tx, "UPDATE inventory SET name=:name , price=:price, active=:active, image_url=:image_url WHERE id = :id", r)

	if err != nil {
		return err
	}

	fmt.Println(rs)
	return nil
}

func (s *ProductService) Delete(ctx context.Context, tx *sqlx.Tx, id int) (err error) {
	log.Info().Msg("Deleting record from inventory")

	rs := sqlx.MustExec(tx, "DELETE FROM inventory WHERE id = ?", id)

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

func (s *ProductService) FindAll(ctx context.Context, clientId int) (c []*Product, err error) {
	log.Info().Msg("Inserting record into Product")

	people := []*Product{}
	err = s.db.Select(&people, "SELECT * FROM inventory where id > 1 and client_id=?", clientId)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return people, nil
}

func (s *ProductService) FindByName(ctx context.Context, name string, clientId int) (c []Product, err error) {
	log.Info().Msg("Getting Product by name")

	products := []Product{}
	err = s.db.Select(&products, "SELECT * FROM inventory where name like ? and client_id=?", "%"+name+"%", clientId)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) FindByID(ctx context.Context, id int, clientID int) (c *Product, err error) {
	log.Info().Msg("Getting Product by id")
	return getProduct(s.log, s.db, id, clientID)
}

func getProduct(log zerolog.Logger, db sqlx.Ext, id int, clientID int) (c *Product, err error) {
	log.Info().Msgf("Getting product id %d", clientID)

	var product Product
	err = sqlx.Get(db, &product, "select * from inventory where id = ? ", id)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get product")
	}

	return &product, nil
}

func updateProduct(baseProduct *Product, updatedProduct *Product) {

	if updatedProduct.ImageUrl == "" {
		updatedProduct.ImageUrl = baseProduct.ImageUrl
	}

	if updatedProduct.Price == "" {
		updatedProduct.Price = baseProduct.Price
	}
}
