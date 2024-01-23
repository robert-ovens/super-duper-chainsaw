package workflows

import (
	"github.com/robert-ovens/super-duper-chainsaw/activities"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

type CloudOperationsWorkflows interface {
	ProvisionInstanceWorkflow(ctx workflow.Context, params ProvisionInstanceWorkflowParams) error
}
type CloudOperationsWorkflowsImpl struct {
	activities     activities.Activities
	temporalClient client.Client
}

func NewCloudOperationsWorkflowsIpml(
	temporalClient client.Client,
	activities activities.Activities,
) CloudOperationsWorkflows {
	return &CloudOperationsWorkflowsImpl{
		temporalClient: temporalClient,
		activities:     activities,
	}
}
