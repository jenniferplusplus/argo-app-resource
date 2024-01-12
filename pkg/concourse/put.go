package concourse

import (
	"argo-app-resource/pkg/argocd"
	"encoding/json"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"time"
)

func (task *Task) Put() error {

	req := DefaultPutRequest()
	decoder := json.NewDecoder(task.stdin)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", err)
	}
	setupLogging(task.stderr, req.Params.Debug)

	var application *v1alpha1.Application
	client, err := argocd.NewClient(&req.Source)
	if err != nil {
		return fmt.Errorf("can't create client: %s", err)
	}

	if req.Source.Project == "" {
		if req.Params.RollbackRevision == "" {
			application, err = client.SyncApplicationLatest(req.Source.App)
		} else {
			application, err = client.SyncApplicationRevision(req.Source.App, req.Params.RollbackRevision)
		}
	} else {
		if req.Params.RollbackRevision == "" {
			application, err = client.SyncApplicationLatestWithProject(req.Source.App, req.Source.Project)
		} else {
			application, err = client.SyncApplicationRevisionWithProject(req.Source.App, req.Params.RollbackRevision, req.Source.Project)
		}
	}

	if err != nil {
		return fmt.Errorf("can't sync the application: %s", err)
	}

	version := Version{
		Revision:   application.Status.Sync.Revision,
		DeployedAt: time.Now().UTC(),
		Health:     string(application.Status.Health.Status),
		SyncStatus: string(application.Status.Sync.Status),
	}
	meta := make([]MetadataField, 0)

	for _, resource := range application.Status.Resources {
		meta = append(meta, MetadataField{
			Name:  resourceMetaKey(&resource),
			Value: fmt.Sprintf("%s/%s", resource.Status, resource.Health.Status),
		})
	}

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
