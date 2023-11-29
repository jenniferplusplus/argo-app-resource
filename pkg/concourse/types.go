package concourse

import (
	"time"
)

type CheckRequest struct {
	Source  Source   `json:"source"`
	Version *Version `json:"version"`
}

type CheckResponse []Version

type GetRequest struct {
	Source  Source     `json:"source"`
	Params  *GetParams `json:"params"`
	Version *Version   `json:"version"`
}

type Response struct {
	Version  Version         `json:"version"`
	Metadata []MetadataField `json:"metadata"`
}

type PutRequest struct {
	Source Source    `json:"source"`
	Params PutParams `json:"params"`
}

type Source struct {
	Host     string `json:"host"`
	Insecure bool   `json:"insecure,omitempty"`
	App      string `json:"app"`
	Project  string `json:"project,omitempty"`
	Token    string `json:"token"`
	Debug    bool   `json:"debug,omitempty"`
}

type Version struct {
	Revision   string    `json:"revision"`
	DeployedAt time.Time `json:"deployed_at"`
	Health     string    `json:"health"`
	SyncStatus string    `json:"sync_status"`
}

type MetadataField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetParams struct {
	Health     string `json:"health"`
	SyncStatus string `json:"sync_status"`
}

type PutParams struct {
	RollbackRevision string `json:"rollback_revision,omitempty"`
}
