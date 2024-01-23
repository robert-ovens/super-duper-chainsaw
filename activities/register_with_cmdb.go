package activities

import (
	"context"
	"fmt"
	"log"
)

type RegisterWithCmdbParams struct {
	Id string
}

// CreateVirtualMachine implements Activities.
func (*ActivitiesImpl) RegisterWithCmdb(ctx context.Context, params RegisterWithCmdbParams) error {
	log.Println(fmt.Sprintf("Registering %s", params.Id))
	return nil
}
