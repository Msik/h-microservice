package repository

import (
	"context"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CategoryRepositorySuite struct {
	suite.Suite
	categoryRepository *CategoryRepository
}

func TestCategoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(CategoryRepositorySuite))
}

func (s *CategoryRepositorySuite) SetupSuite() {
	sqlxDB, err := sqlx.Connect("postgres", os.Getenv("DB_SQLX_CONNECT"))
	require.NoError(s.T(), err)

	s.categoryRepository = NewCategoryRepository(sqlxDB)
}

func (s *CategoryRepositorySuite) TestStoreAndDeleteCategory() {
	title := "test-title"

	ctx := context.Background()
	id, err := s.categoryRepository.Store(ctx, title)
	require.NoError(s.T(), err)

	category, err := s.categoryRepository.GetById(ctx, id)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), category.Title, title)

	err = s.categoryRepository.Delete(ctx, id)
	assert.NoError(s.T(), err)
}
