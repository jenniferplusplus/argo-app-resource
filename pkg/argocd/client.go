package argocd

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

type Connection struct {
	Address string
	Token   string
}

type Client struct {
	projectClient project.ProjectServiceClient
	clusterClient cluster.ClusterServiceClient
	appClient     application.ApplicationServiceClient
}

func NewClient(conn *Connection) (*Client, error) {
	apiClient, err := apiclient.NewClient(&apiclient.ClientOptions{
		ServerAddr: fmt.Sprintf(conn.Address),
		Insecure:   true,
		AuthToken:  conn.Token,
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
	return c.GetApplicationWithProject(name, "default")
}

func (c *Client) GetApplicationWithProject(name string, project string) (*v1alpha1.Application, error) {
	return c.appClient.Get(context.Background(), &application.ApplicationQuery{
		Name:    &name,
		Project: []string{project},
	})
}
