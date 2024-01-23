# super-duper-chainsaw

## Temporal local development setup

Follow instructions at [set up a local development environment for Temporal and Go](https://learn.temporal.io/getting_started/go/dev_environment).

## Running a local dev server

```shell
temporal server start-dev 
```

The admin ui will be available at http://localhost:8233


## Invoking a workflow

This is a fictional worklow that
- waits for an approval
- provisions an instance and gets the instance id
- registers the instance in the CMDB using the instance id from the provisioning activity

### Starting an instance
```shell
temporal workflow start \
--task-queue cloud-operations-queue \
--type ProvisionInstanceWorkflow \
--namespace default \
--workflow-id $(uuidgen)
```

### Approving the request

using the workflow-id from the previous step 
```shell
temporal workflow signal \
--workflow-id id-from-the-previous-step
--name "approve-request"
```

### Query the workflow

using the workflow-id from the previous step 
```shell
temporal workflow query \
--workflow-id id-from-the-previous-step
--type "current_state"
```
