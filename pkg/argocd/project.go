package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/project"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func (c *Client) GetProject(name string) (*v1alpha1.AppProject, error) {
	return c.projectClient.Get(context.Background(), &project.ProjectQuery{
		Name: name,
	})
}
