package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robert-ovens/super-duper-chainsaw/activities"
	"github.com/robert-ovens/super-duper-chainsaw/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	temporalClient, err := client.Dial(client.Options{
		HostPort:  fmt.Sprintf("%s:%s", os.Getenv("TEMPORAL_HOST"), os.Getenv("TEMPORAL_PORT")),
		Namespace: os.Getenv("TEMPORAL_NAMESPACE"),
	})
	if err != nil {
		log.Panic(err)
	}

	act := activities.NewActivitiesImpl()
	petWorkflows := workflows.NewCloudOperationsWorkflowsIpml(
		temporalClient,
		act,
	)

	defer temporalClient.Close()
	cloudOperationsWorker := worker.New(temporalClient, "cloud-operations-queue", worker.Options{})

	cloudOperationsWorker.RegisterWorkflow(petWorkflows.ProvisionInstanceWorkflow)

	cloudOperationsWorker.RegisterActivity(act)

	err = cloudOperationsWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start the Worker Process", err)
	}

}
