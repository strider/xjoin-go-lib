package pkg

// +k8s:deepcopy-gen=true
type ValidationResponse struct {
	Result  ValidationResult `json:"result,omitempty"`
	Reason  string           `json:"reason,omitempty"`
	Message string           `json:"message,omitempty"`
	Details ResponseDetails  `json:"details,omitempty"`
}

// ValidationResult from the validation run.
// +enum
type ValidationResult string

// These are the possible validation results.
const (
	ValidationValid   = "valid"
	ValidationInvalid = "invalid"
)

// +k8s:deepcopy-gen=true
type ResponseDetails struct {
	TotalMismatch int `json:"totalMismatch,omitempty"`

	IdsMissingFromElasticsearch      []string `json:"idsMissingFromElasticsearch,omitempty"`
	IdsMissingFromElasticsearchCount int      `json:"idsMissingFromElasticsearchCount,omitempty"`

	IdsOnlyInElasticsearch      []string `json:"idsOnlyInElasticsearch,omitempty"`
	IdsOnlyInElasticsearchCount int      `json:"idsOnlyInElasticsearchCount,omitempty"`

	IdsWithMismatchContent []string `json:"idsWithMismatchContent,omitempty"`

	MismatchContentDetails []MismatchContentDetails `json:"mismatchContentDetails,omitempty"`
}

// +k8s:deepcopy-gen=true
type MismatchContentDetails struct {
	ID                   string `json:"id,omitempty"`
	ElasticsearchContent string `json:"elasticsearchContent,omitempty"`
	DatabaseContent      string `json:"databaseContent,omitempty"`
}

// +k8s:deepcopy-gen=true
type DatasourceDBConnection struct {
	SslMode  string `json:"sslMode,omitempty"`
	Name     string `json:"name,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
	Table    string `json:"table,omitempty"`
	Port     string `json:"port,omitempty"`
}
