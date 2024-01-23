package activities

import (
	"context"
)

type Activities interface {
	CreateVirtualMachine(ctx context.Context, params CreateVirtualMachineParams) (*string, error)
	RegisterWithCmdb(ctx context.Context, params RegisterWithCmdbParams) error
}
type ActivitiesImpl struct {
}

func NewActivitiesImpl() Activities {
	return &ActivitiesImpl{}
}
