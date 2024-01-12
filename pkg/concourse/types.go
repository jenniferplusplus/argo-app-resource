package concourse

import (
	"argo-app-resource/pkg/resource"
	"time"
)

type CheckRequest struct {
	Source  resource.Source `json:"source"`
	Version *Version        `json:"version"`
}

type CheckResponse []Version

type GetRequest struct {
	Source  resource.Source `json:"source"`
	Params  *GetParams      `json:"params"`
	Version *Version        `json:"version"`
}

type Response struct {
	Version  Version         `json:"version"`
	Metadata []MetadataField `json:"metadata"`
}

type PutRequest struct {
	Source resource.Source `json:"source"`
	Params PutParams       `json:"params"`
}

type Version struct {
	Revision   string    `json:"revision"`
	DeployedAt time.Time `json:"deployed_at"`
	Health     string    `json:"health"`
	SyncStatus string    `json:"sync_status"`
}

type MetadataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetParams struct {
	Health     string `json:"health"`
	SyncStatus string `json:"sync_status"`
	Debug      bool   `json:"debug,omitempty"`
}

type PutParams struct {
	RollbackRevision string `json:"rollback_revision,omitempty"`
	Debug            bool   `json:"debug,omitempty"`
}

func DefaultGetRequest() *GetRequest {
	req := GetRequest{Source: resource.DefaultSource()}
	return &req
}

func DefaultPutRequest() *PutRequest {
	req := PutRequest{Source: resource.DefaultSource()}
	return &req
}

func DefaultCheckRequest() *CheckRequest {
	req := CheckRequest{Source: resource.DefaultSource()}
	return &req
}
