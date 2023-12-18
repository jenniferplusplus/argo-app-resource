package resource

type Source struct {
	Host       string `json:"host"`
	Insecure   bool   `json:"insecure,omitempty"`
	App        string `json:"app"`
	Project    string `json:"project,omitempty"`
	Token      string `json:"token"`
	Debug      bool   `json:"debug,omitempty"`
	UseGrpcWeb bool   `json:"use_grpc_web,omitempty"`
}

func DefaultSource() Source {
	source := Source{
		Insecure:   false,
		Debug:      false,
		UseGrpcWeb: false,
	}
	return source
}
