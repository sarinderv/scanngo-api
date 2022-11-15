package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductBookCreate(t *testing.T) {

	store := &ProductService{
		db: db,
	}
	tx = db.Begin()
	b, err := store.Create(context.TODO(), &Product{
		Name:  "some product",
		Price: "100",
	})

	require.NoError(t, err)
	assert.NotZero(t, b.ID)
}

func TestProductBookBulkCreate(t *testing.T) {

	store := &ProductService{
		db: db,
	}
	tx = db.Begin()
	err := store.BulkCreate(context.TODO(),[]{ &Product{
		Name:  "some product",
		Price: "100",
	},
	&Product{
		Name:  "some other product",
		Price: "100",
	}
})

	require.NoError(t, err)
}

func TestProductGet(t *testing.T) {

	store := &ProductService{
		db: db,
	}
	tx = db.Begin()

	id := 2
	b, err := store.Find(context.TODO(), 2)

	require.NoError(t, err)
	
	assert.equals(b.Name, "that product")
	assert.equals(b.Price, "1001")
}
