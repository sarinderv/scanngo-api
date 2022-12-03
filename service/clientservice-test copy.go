package service

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// var db *sqlx.DB

// func TestClientCreate(t *testing.T) {
// 	db , er : = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "a", "b", "c", "d"))
// 	if err != nil {
// 		log.Error().Msg(err.Error())
// 		return nil, err
// 	}

// 	db.Close()
// 	store := &clientService{
// 		db: db,
// 	}
// 	tx := db.MustBegin()

// 	b, err := store.Create(context.TODO(), tx, &Client{
// 		Name:    "some Client",
// 		Address: "some address",
// 	})

// 	require.NoError(t, err)
// 	assert.NotZero(t, b.ID)
// }

// func TestClientBulkCreate(t *testing.T) {

// 	store := &clientService{
// 		db: db,
// 	}
// 	tx = db.Begin()
// 	err := store.BulkCreate(context.TODO(), []*Client{
// 		&Client{
// 			Name:    "some Client",
// 			Adderss: "some otheraddress",
// 		},
// 		&Client{
// 			Name:    "some other Client",
// 			Adderss: "some random address",
// 		},
// 	})

// 	require.NoError(t, err)
// }

// func TestClientGet(t *testing.T) {
// 	db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "a", "b", "c", "d"))

// 	store := &ClientService{
// 		db: db,
// 	}
// 	tx = db.Begin()

// 	id := 2
// 	b, err := store.Find(context.TODO(), 2)

// 	require.NoError(t, err)

// 	assert.equals(b.Name, "that Client")
// 	assert.equals(b.Address, "this address")
// }
