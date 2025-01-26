package common

import (
	"context"
	"test-blueprint/internal/repository/model"
)

type RepositoryService interface {
	ListUserRepos(ctx context.Context, username string) ([]model.Repository, error)
}
