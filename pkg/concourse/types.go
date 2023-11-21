package concourse

type CheckRequest struct {
	Source  Source   `json:"source"`
	Version *Version `json:"version"`
}

type CheckResponse []Version

type GetRequest struct {
	Source  Source    `json:"source"`
	Params  GetParams `json:"params"`
	Version Version   `json:"version"`
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
	Username string `json:"username"`
	Token    string `json:"token"`
	Debug    bool   `json:"debug,omitempty"`
}

type Version struct {
	SyncWindow string `json:"syncWindow,omitempty"`
	Revision   string `json:"revision,omitempty"`
	Health     string `json:"health,omitempty"`
}

type MetadataField struct {
	Group     string `json:"group,omitempty"`
	Kind      string `json:"kind,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Status    string `json:"status,omitempty"`
	Health    string `json:"health,omitempty"`
	Hook      string `json:"hook,omitempty"`
	Message   string `json:"message,omitempty"`
}

// TODO
type GetParams struct {
	RawFormat    string `json:"format"`
	SkipDownload bool   `json:"skip_download"`
}

type PutParams struct {
	RollbackRevision string `json:"rollback_revision,omitempty"`
}
