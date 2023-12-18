package concourse

import (
	"argo-app-resource/pkg/argocd"
	"encoding/json"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func (task *Task) Check() error {
	setupLogging(task.stderr, true)

	req := DefaultCheckRequest()
	decoder := json.NewDecoder(task.stdin)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", err)
	}

	client, err := argocd.NewClient(&req.Source)
	if err != nil {
		return fmt.Errorf("can't create client: %s", err)
	}

	var application *v1alpha1.Application
	if req.Source.Project == "" {
		application, err = client.GetApplication(req.Source.App)
	} else {
		application, err = client.GetApplicationWithProject(req.Source.App, req.Source.Project)
	}

	if err != nil {
		return fmt.Errorf("can't load the application: %s", err)
	}

	version := Version{
		Revision:   application.Status.Sync.Revision,
		DeployedAt: application.Status.OperationState.StartedAt.UTC(),
		Health:     string(application.Status.Health.Status),
		SyncStatus: string(application.Status.Sync.Status),
	}
	var versions []Version
	versions = append(versions, version)

	if len(versions) == 0 {
		return fmt.Errorf("couldn't get application status")
	}

	err = json.NewEncoder(task.stdout).Encode(versions)
	if err != nil {
		return fmt.Errorf("could not serialize versions: %s", err)
	}

	return nil
}
