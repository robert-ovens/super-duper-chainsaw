package activities

import (
	"context"
	"errors"
	"os"

	"github.com/google/uuid"
)

type CreateVirtualMachineParams struct {
	Name  string
	Cores int
}

// CreateVirtualMachine implements Activities.
func (*ActivitiesImpl) CreateVirtualMachine(ctx context.Context, params CreateVirtualMachineParams) (*string, error) {
	if os.Getenv("FAIL") == "true" {
		return nil, errors.New("some provisioning error")
	}
	id := uuid.NewString()
	return &id, nil
}
