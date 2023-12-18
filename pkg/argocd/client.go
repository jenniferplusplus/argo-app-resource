package argocd

import (
	"argo-app-resource/pkg/resource"
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

type Client struct {
	projectClient project.ProjectServiceClient
	clusterClient cluster.ClusterServiceClient
	appClient     application.ApplicationServiceClient
}

func NewClient(conn *resource.Source) (*Client, error) {
	apiClient, err := apiclient.NewClient(&apiclient.ClientOptions{
		ServerAddr: conn.Host,
		Insecure:   conn.Insecure,
		AuthToken:  conn.Token,
		GRPCWeb:    conn.UseGrpcWeb,
	})
	if err != nil {
		return nil, err
	}

	_, projectClient, err := apiClient.NewProjectClient()
	if err != nil {
		return nil, err
	}

	_, clusterClient, err := apiClient.NewClusterClient()
	if err != nil {
		return nil, err
	}

	_, appClient, err := apiClient.NewApplicationClient()

	return &Client{
		projectClient: projectClient,
		clusterClient: clusterClient,
		appClient:     appClient,
	}, nil
}

func (c *Client) GetApplication(name string) (*v1alpha1.Application, error) {
	return c.appClient.Get(context.Background(), &application.ApplicationQuery{
		Name: &name,
	})
}

func (c *Client) GetApplicationWithProject(name string, project string) (*v1alpha1.Application, error) {
	return c.appClient.Get(context.Background(), &application.ApplicationQuery{
		Name:    &name,
		Project: []string{project},
	})
}

func (c *Client) SyncApplicationRevision(name string, revision string) (*v1alpha1.Application, error) {
	return c.appClient.Sync(context.Background(), &application.ApplicationSyncRequest{
		Name:     &name,
		Revision: &revision,
	})
}

func (c *Client) SyncApplicationLatest(name string) (*v1alpha1.Application, error) {
	return c.appClient.Sync(context.Background(), &application.ApplicationSyncRequest{
		Name: &name,
	})
}

func (c *Client) SyncApplicationRevisionWithProject(name string, revision string, project string) (*v1alpha1.Application, error) {
	return c.appClient.Sync(context.Background(), &application.ApplicationSyncRequest{
		Name:     &name,
		Revision: &revision,
		Project:  &project,
	})
}

func (c *Client) SyncApplicationLatestWithProject(name string, project string) (*v1alpha1.Application, error) {
	return c.appClient.Sync(context.Background(), &application.ApplicationSyncRequest{
		Name:    &name,
		Project: &project,
	})
}
