package concourse

import (
	"argo-app-resource/pkg/argocd"
	"encoding/json"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func (task *Task) Get() error {
	req := DefaultGetRequest()
	decoder := json.NewDecoder(task.stdin)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(req)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", err)
	}
	setupLogging(task.stderr, req.Params.Debug)
	dest := task.args[1]

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

	if req.Params != nil && req.Params.Health != "" && version.Health != req.Params.Health {
		return fmt.Errorf("current application health is %s, required %s", version.Health, req.Params.Health)
	}

	if req.Params != nil && req.Params.SyncStatus != "" && version.SyncStatus != req.Params.SyncStatus {
		return fmt.Errorf("current application state is %s, required %s", version.SyncStatus, req.Params.SyncStatus)
	}

	meta := make([]MetadataField, 0)

	for _, res := range application.Status.Resources {
		meta = append(meta, MetadataField{
			Name:  resourceMetaKey(&res),
			Value: fmt.Sprintf("%s", res.Status),
		})
	}

	response := Response{
		Version:  version,
		Metadata: meta,
	}

	err = saveVersion(dest, &version)
	if err != nil {
		logrus.Error(err)
	}

	err = saveRevisionHistory(dest, application)
	if err != nil {
		logrus.Error(err)
	}

	err = json.NewEncoder(task.stdout).Encode(response)
	if err != nil {
		return fmt.Errorf("could not serialize response: %s", err)
	}

	return nil
}

func resourceMetaKey(status *v1alpha1.ResourceStatus) string {
	var ns string
	if status.Namespace != "" {
		ns = status.Namespace
	} else {
		ns = "_"
	}
	return fmt.Sprintf(
		"%s/%s.%s.%s",
		status.Kind,
		status.Version,
		ns,
		status.Name,
	)
}

func saveVersion(dest string, version *Version) (err error) {
	f, err := os.Create(filepath.Join(dest, "version.json"))
	if err != nil {
		return fmt.Errorf("cannot write to file %s", filepath.Join(dest, "version.json"))
	}
	defer closeFile(f, &err)

	err = json.NewEncoder(f).Encode(version)
	if err != nil {
		return fmt.Errorf("could not serialize version to file: %w", err)
	}

	return err
}

func saveRevisionHistory(dest string, application *v1alpha1.Application) error {
	f, err := os.Create(filepath.Join(dest, "history.json"))
	if err != nil {
		return fmt.Errorf("cannot write to file %s", filepath.Join(dest, "history.json"))
	}
	defer closeFile(f, &err)

	err = json.NewEncoder(f).Encode(application.Status.History)
	if err != nil {
		return fmt.Errorf("could not serialize history to file: %w", err)
	}

	return nil
}

func closeFile(f *os.File, err *error) {
	c := f.Close()
	if err == nil {
		err = &c
	}
}
