package awsmwaa


// Experimental.
type TfEnvironment_WorkerLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#enabled TfEnvironment#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#log_level TfEnvironment#log_level}.
	// Experimental.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

