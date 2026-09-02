package awsmskconnect


// Experimental.
type TfConnector_CloudwatchLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#enabled TfConnector#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#log_group TfConnector#log_group}.
	// Experimental.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

