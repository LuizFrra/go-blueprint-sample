package service

import (
	"context"
	"fmt"
	"test-blueprint/internal/repository/common"
	"test-blueprint/internal/repository/model"
)

type ListUserReposByPlatformService struct {
	platforms map[string]common.RepositoryService
}

func (l *ListUserReposByPlatformService) ListUserRepos(ctx context.Context, username string, platform string) ([]model.Repository, error) {
	platformService, ok := l.platforms[platform]

	if !ok {
		return nil, fmt.Errorf("platform not found")
	}

	return platformService.ListUserRepos(ctx, username)
}

type ListUserReposByPlatformServiceBuilder struct {
	platforms map[string]common.RepositoryService
}

func NewListUserReposByPlatformServiceBuilder() *ListUserReposByPlatformServiceBuilder {
	return &ListUserReposByPlatformServiceBuilder{
		platforms: make(map[string]common.RepositoryService),
	}
}

func (b *ListUserReposByPlatformServiceBuilder) AddPlatform(name string, service common.RepositoryService) *ListUserReposByPlatformServiceBuilder {
	b.platforms[name] = service
	return b
}

func (b *ListUserReposByPlatformServiceBuilder) Build() ListUserReposByPlatformService {
	return ListUserReposByPlatformService{
		platforms: b.platforms,
	}
}
