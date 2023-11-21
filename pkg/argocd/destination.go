package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func (c *Client) AddDestination(projectName, server, namespace, name string) error {
	p, err := c.GetProject(projectName)
	if err != nil {
		return err
	}

	p.Spec.Destinations = []v1alpha1.ApplicationDestination{
		{
			Server:    server,
			Namespace: namespace,
			Name:      name,
		},
	}

	_, err = c.projectClient.Update(context.Background(), &project.ProjectUpdateRequest{
		Project: p,
	})
	if err != nil {
		return err
	}

	return nil
}
