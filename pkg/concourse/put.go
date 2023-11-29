package concourse

import (
	"argo-app-resource/pkg/argocd"
	"encoding/json"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func (task *Task) Put() error {
	setupLogging(task.stderr)

	var req GetRequest
	decoder := json.NewDecoder(task.stdin)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", err)
	}

	connection := argocd.Connection{
		Address: req.Source.Host,
		Token:   req.Source.Token,
	}

	client, err := argocd.NewClient(&connection)
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

	//application

	version := Version{
		Revision:   application.Status.Sync.Revision,
		DeployedAt: application.Status.OperationState.StartedAt.UTC(),
		Health:     string(application.Status.Health.Status),
		SyncStatus: string(application.Status.Sync.Status),
	}
	meta := make([]MetadataField, 0)

	//for _, resource := range application.Status.Resources {
	//	meta = append(meta, MetadataField{
	//		Key:   resourceSyncKey(&resource),
	//		Value: string(resource.Status),
	//	}, MetadataField{
	//		Key:   resourceHealthKey(&resource),
	//		Value: string(resource.Health.Status),
	//	})
	//}

	response := Response{
		Version:  version,
		Metadata: meta,
	}

	err = json.NewEncoder(task.stdout).Encode(response)
	if err != nil {
		return fmt.Errorf("could not serialize response: %s", err)
	}

	return nil
}
