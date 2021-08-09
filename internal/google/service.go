package google

import (
	"context"
	admin "google.golang.org/api/admin/directory/v1"
	"strconv"
)

// User

type UsersService interface {
	Get(GetParams) (*User, error)
	Query(QueryParams) ([]*User, error)
	Deleted(DeletedParams) ([]*User, error)
}

type usersService struct {
	ctx context.Context
	usersService *admin.UsersService
}

func (r *usersService) Get(p GetParams) (*User, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	ret, err := r.usersService.Get(p.Id).Do()
	return (*User)(ret), err
}

func (r *usersService) Query(p QueryParams) ([]*User, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	c := r.usersService.List().ShowDeleted(strconv.FormatBool(p.ShowDeleted))
	if p.Customer != "" {
		c = c.Customer(p.Customer)
	}
	if p.Domain != "" {
		c = c.Domain(p.Domain)
	}
	if p.Query != "" {
		c = c.Query(p.Query)
	}

	ret := make([]*User, 0)
	err = c.Query(p.Query).Pages(r.ctx, func(users *admin.Users) error {
		for _, user := range users.Users {
			ret = append(ret, (*User)(user))
		}
		return nil
	})

	return ret, err
}

func (r *usersService) Deleted(p DeletedParams) ([]*User, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return r.Query(QueryParams{
		AccountParams: p,
		ShowDeleted: true,
	})
}

// Group

type GroupsService interface {
	Get(GetParams) (*Group, error)
	Query(QueryParams) ([]*Group, error)
	Members(GroupMembersParams) ([]*Member, error)
}

type groupsService struct {
	ctx context.Context
	groupsService *admin.GroupsService
	membersService *admin.MembersService
}

func (r *groupsService) Get(p GetParams) (*Group, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	ret, err := r.groupsService.Get(p.Id).Do()
	return (*Group)(ret), err
}

func (r *groupsService) Query(p QueryParams) ([]*Group, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	c := r.groupsService.List()
	if p.Customer != "" {
		c = c.Customer(p.Customer)
	}
	if p.Domain != "" {
		c = c.Domain(p.Domain)
	}
	if p.Query != "" {
		c = c.Query(p.Query)
	}

	ret := make([]*Group, 0)
	err = c.Query(p.Query).Pages(r.ctx, func(groups *admin.Groups) error {
		for _, group := range groups.Groups {
			ret = append(ret, (*Group)(group))
		}
		return nil
	})

	return ret, err
}

func (r *groupsService) Members(p GroupMembersParams) ([]*Member, error) {
	err := p.Validate()
	if err != nil {
		return nil, err
	}

	id := p.Id
	if id == "" {
		id = p.Parent.Id
	}

	ret := make([]*Member, 0)
	err = r.membersService.List(id).Pages(r.ctx, func(members *admin.Members) error {
		for _, member := range members.Members {
			ret = append(ret, (*Member)(member))
		}
		return nil
	})

	return ret, err
}
