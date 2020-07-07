package sql_store_test

import (
	"github.com/sgnatyuk/http-rest-api/internal/app/model"
	"github.com/sgnatyuk/http-rest-api/internal/app/store"
	"github.com/sgnatyuk/http-rest-api/internal/app/store/sql_store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	db, teardown := sql_store.TestDB(t, databaseURL)
	defer teardown("users")

	s := sql_store.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {

	db, teardown := sql_store.TestDB(t, databaseURL)
	defer teardown("users")

	s := sql_store.New(db)
	u1 := model.TestUser(t)

	_ = s.User().Create(u1)

	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	db, teardown := sql_store.TestDB(t, databaseURL)
	defer teardown("users")

	s := sql_store.New(db)

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	_ = s.User().Create(u)
	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}