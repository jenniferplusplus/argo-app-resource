package argocd

import (
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
)

type Connection struct {
	Address string
	Token   string
}

type Client struct {
	projectClient project.ProjectServiceClient
	clusterClient cluster.ClusterServiceClient
}

func NewClient(c *Connection) (*Client, error) {
	apiClient, err := apiclient.NewClient(&apiclient.ClientOptions{
		ServerAddr: fmt.Sprintf(c.Address),
		Insecure:   true,
		AuthToken:  c.Token,
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

	return &Client{
		projectClient: projectClient,
		clusterClient: clusterClient,
	}, nil
}
