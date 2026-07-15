package apps

import (
	"context"
	"fmt"

	"github.com/google/go-github/v89/github"
	"github.com/pkg/errors"
)

// Installation is a struct representing a GitHub installation.
type Installation struct {
	// ID is the installation ID.
	ID int64 `json:"id"`
	// Owner is the owner of the installation.
	Owner string `json:"owner"`
	// OwnerID is the owner's ID.
	OwnerID int64 `json:"repo"`
}

// Installations is an interface for managing GitHub installations.
type Installations interface {
	// ListAll returns all installations for this app.
	ListAll(ctx context.Context) ([]*Installation, error)
	// GetByOwner returns the installation for an owner (user or organization).
	// It returns an InstallationNotFound error if no installation exists for
	// the owner.
	GetByOwner(ctx context.Context, owner string) (*Installation, error)
	// GetByRepository returns the installation for a repository.
	// It returns an InstallationNotFound error if no installation exists for
	// the repository.
	GetByRepository(ctx context.Context, owner string, repo string) (*Installation, error)
}

type installationsImpl struct {
	client *github.Client
}

var _ Installations = (*installationsImpl)(nil)

// NewInstallations returns a new Installations instance.
func NewInstallations(client *github.Client) *installationsImpl {
	return &installationsImpl{client: client}
}

// ListAll returns all installations for this app.
func (i *installationsImpl) ListAll(ctx context.Context) ([]*Installation, error) {
	opt := &github.ListOptions{
		PerPage: 100,
	}

	var installations []*Installation
	for {
		list, resp, err := i.client.Apps.ListInstallations(ctx, opt)
		if err != nil {
			return nil, err
		}

		for _, item := range list {
			installations = append(installations, &Installation{
				ID:      item.GetAppID(),
				Owner:   item.GetAccount().GetLogin(),
				OwnerID: item.GetAccount().GetID(),
			})
		}

		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	return installations, nil
}

// GetByOwner returns the installation for an owner (user or organization).
// It returns an InstallationNotFound error if no installation exists for
// the owner.
func (i *installationsImpl) GetByOwner(ctx context.Context, owner string) (*Installation, error) {
	installation, _, err := i.client.Apps.GetOrganizationInstallation(ctx, owner)
	if err == nil {
		return &Installation{
			ID:      installation.GetAppID(),
			Owner:   installation.GetAccount().GetLogin(),
			OwnerID: installation.GetAccount().GetID(),
		}, nil
	}

	var rerr *github.ErrorResponse
	if ok := errors.As(err, &rerr); ok && rerr.Response.StatusCode != 404 {
		return nil, rerr
	}

	installation, _, err = i.client.Apps.GetUserInstallation(ctx, owner)
	if err != nil {
		return nil, err
	}

	return &Installation{
		ID:      installation.GetAppID(),
		Owner:   installation.GetAccount().GetLogin(),
		OwnerID: installation.GetAccount().GetID(),
	}, fmt.Errorf("failed to get installation for owner %q: %w", owner, err)
}

// GetByRepository returns the installation for a repository.
func (i *installationsImpl) GetByRepository(ctx context.Context, owner string, repo string) (*Installation, error) {
	installation, _, err := i.client.Apps.GetRepositoryInstallation(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	return &Installation{
		ID:      installation.GetAppID(),
		Owner:   installation.GetAccount().GetLogin(),
		OwnerID: installation.GetAccount().GetID(),
	}, nil
}
