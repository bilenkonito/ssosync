package google

import (
	"context"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

// Instance

type Client interface {
	Users() UsersService
	Groups() GroupsService
}

type client struct {
	usersService *usersService
	groupsService *groupsService
}

func (r *client) Users() UsersService {
	return r.usersService
}

func (r *client) Groups() GroupsService {
	return r.groupsService
}

// Constructor

func NewClient(ctx context.Context, p ClientParams) (Client, error) {
	if ctx == nil {
		ctx = context.TODO()
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	config, err := parseConfig(p)
	if err != nil {
		return nil, err
	}

	options, err := parseOptions(ctx, config)
	if err != nil {
		return nil, err
	}

	service, err := createService(ctx, options)
	if err != nil {
		return nil, err
	}

	return &client{
		usersService: &usersService{
			ctx: ctx,
			usersService: service.Users,
		},
		groupsService: &groupsService{
			ctx: ctx,
			groupsService: service.Groups,
			membersService: service.Members,
		},
	}, nil
}

func parseConfig(p ClientParams) (*jwt.Config, error) {
	config, err := google.JWTConfigFromJSON(p.ServiceAccountKey,
		admin.AdminDirectoryGroupReadonlyScope,
		admin.AdminDirectoryGroupMemberReadonlyScope,
		admin.AdminDirectoryUserReadonlyScope)
	if err != nil {
		return nil, err
	}

	config.Subject = p.AdminEmail

	return config, nil
}

func parseOptions(ctx context.Context, config *jwt.Config) ([]option.ClientOption, error) {
	options := make([]option.ClientOption, 0)

	options = append(options, option.WithTokenSource(config.TokenSource(ctx)))

	return options, nil
}

func createService(ctx context.Context, options []option.ClientOption) (*admin.Service, error) {
	return admin.NewService(ctx, options...)
}
