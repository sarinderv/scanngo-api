package service

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/rs/zerolog/log"
// )

// type Product struct {
// 	ID               int           `json:"id" db:"id"`
// 	Name             string        `json:"name" db:"name"`
// 	Address          string        `json:"address" db:"address"`
// 	Active           bool          `json:"active" db:"active"`
// 	DeactivationDate sql.NullTime  `json:"deactivationDate" db:"deactivationDate"`
// 	Insert_Timestamp sql.NullTime  `json:"ins_ts" db:"ins_ts"`
// 	Update_Timestamp sql.NullTime  `json:"upd_ts" db:"upd_ts"`
// 	Insert_User      int           `db:"ins_user"`
// 	Update_User      sql.NullInt32 `db:"upd_user"`
// }

// type IProductService interface {
// 	BulkCreate(ctx context.Context, tx *sqlx.Tx, r []*Product) (error)
// 	Create(ctx context.Context, tx *sqlx.Tx, r *Product) (*Product, error)
// 	Update(ctx context.Context, tx *sqlx.Tx, r *Product) error
// 	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
// 	FindAll(ctx context.Context) ([]*Product, error)
// 	FindByID(ctx context.Context, id int) (*Product, error)
// }

// type ProductService struct {
// 	db *sqlx.DB
// }

// func NewProductService(db *sqlx.DB) IProductService {
// 	return &ProductService{
// 		db,
// 	}
// }

// func (s *ProductService) Create(ctx context.Context, tx *sqlx.Tx, r *Product) (*Product, error) {
// 	log.Info().Msg("Inserting record into Product")

// 	rs, err := sqlx.NamedExec(tx, "INSERT INTO Product (name, price, active , ins_user) VALUES (:name, :price, :active, :ins_user )", r)

// 	if err != nil {
// 		log.Error().Msg(err.Error())
// 		return nil, err
// 	}

// 	id, _ := rs.LastInsertId()

// 	r.ID = int(id)

// 	fmt.Println(rs)

// 	return r, nil
// }

// func (s *ProductService) Update(ctx context.Context, tx *sqlx.Tx, r *Product) error {
// 	log.Info().Msg("Inserting record into Product")

// 	rs, err := sqlx.NamedExec(tx, "UPDATE Product SET name=:name , address=:address, active=:active, deactivationDate=:deactivationDate, upd_user=:upd_user WHERE id = :id", r)

// 	if err != nil {
// 		log.Error().Msg(err.Error())
// 		return err
// 	}

// 	fmt.Println(rs)
// 	return nil
// }

// func (s *ProductService) Delete(ctx context.Context, tx *sqlx.Tx, id int) (err error) {
// 	log.Info().Msg("Deleting record from Product")

// 	rs := sqlx.MustExec(tx, "DELETE FROM Product WHERE id = ?", id)

// 	rows, err := rs.RowsAffected()
// 	if rows == 0 {

// 		return errors.New("asdasd")
// 	}

// 	if err != nil {

// 		return err
// 	}

// 	fmt.Println(rs)
// 	return nil
// }

// func (s *ProductService) FindAll(ctx context.Context) (c []*Product, err error) {
// 	log.Info().Msg("Inserting record into Product")

// 	people := []*Product{}
// 	err = s.db.Select(&people, "SELECT * FROM Product where id > 1")
// 	if err != nil {
// 		log.Error().Msg(err.Error())
// 		return nil, err
// 	}
// 	return people, nil
// }

// func (s *ProductService) FindByID(ctx context.Context, id int) (c *Product, err error) {
// 	log.Info().Msg("Getting Product by id")

// 	people := Product{}
// 	err = s.db.Get(&people, "SELECT * FROM Product where id=? LIMIT 1", id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &people, nil
// }
