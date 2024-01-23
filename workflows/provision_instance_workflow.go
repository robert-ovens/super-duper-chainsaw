package workflows

import (
	"time"

	"github.com/robert-ovens/super-duper-chainsaw/activities"
	"go.temporal.io/sdk/workflow"
)

type ProvisionInstanceWorkflowState struct {
	Id    *string
	State string
}
type ProvisionInstanceWorkflowParams struct {
	InstanceType string
	Cores        int
	Name         string
}

type ApproveRequestSignal struct {
}

// ProvisionInstanceWorkflow implements PetWorkflows.
func (p *CloudOperationsWorkflowsImpl) ProvisionInstanceWorkflow(ctx workflow.Context, params ProvisionInstanceWorkflowParams) error {
	ao := workflow.ActivityOptions{StartToCloseTimeout: 10000 * time.Second}

	ctx = workflow.WithActivityOptions(ctx, ao)

	state := ProvisionInstanceWorkflowState{State: "initial"}

	queryType := "current_state"
	err := workflow.SetQueryHandler(ctx, queryType, func() (ProvisionInstanceWorkflowState, error) {

		return state, nil
	})
	if err != nil {
		return err
	}

	err = workflow.SetQueryHandler(ctx, queryType, func() (ProvisionInstanceWorkflowState, error) {
		return state, nil
	})

	if err != nil {
		return err
	}

	var approveRequestSignal ApproveRequestSignal
	approveRequestSignalSignalChan := workflow.GetSignalChannel(ctx, "approve-request")
	approveRequestSignalSignalChan.Receive(ctx, &approveRequestSignal)

	var id string
	err = workflow.ExecuteActivity(ctx, p.activities.CreateVirtualMachine, activities.CreateVirtualMachineParams{
		Name:  params.Name,
		Cores: params.Cores,
	}).Get(ctx, &id)
	if err != nil {
		return err
	}

	state.State = "instance-provisioned"
	state.Id = &id

	err = workflow.ExecuteActivity(ctx, p.activities.RegisterWithCmdb, activities.RegisterWithCmdbParams{
		Id: *state.Id,
	}).Get(ctx, nil)
	if err != nil {
		return err
	}

	state.State = "instance-registered"

	return nil
}
