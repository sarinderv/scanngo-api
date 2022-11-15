package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientCreate(t *testing.T) {

	store := &ClientService{
		db: db,
	}
	tx = db.Begin()
	b, err := store.Create(context.TODO(), &Client{
		Name:  "some Client",
		Adderss: "some address",
	})

	require.NoError(t, err)
	assert.NotZero(t, b.ID)
}

func TestClientBulkCreate(t *testing.T) {

	store := &ClientService{
		db: db,
	}
	tx = db.Begin()
	err := store.BulkCreate(context.TODO(),[]{ &Client{
		Name:  "some Client",
		Adderss: "some otheraddress",,
	},
	&Client{
		Name:  "some other Client",
		Adderss: "some random address",
	}
})

	require.NoError(t, err)
}

func TestClientGet(t *testing.T) {

	store := &ClientService{
		db: db,
	}
	tx = db.Begin()

	id := 2
	b, err := store.Find(context.TODO(), 2)

	require.NoError(t, err)
	
	assert.equals(b.Name, "that Client")
	assert.equals(b.Address, "this address")
}
