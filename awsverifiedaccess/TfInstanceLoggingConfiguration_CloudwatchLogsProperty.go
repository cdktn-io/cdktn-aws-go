package awsverifiedaccess


// Experimental.
type TfInstanceLoggingConfiguration_CloudwatchLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#enabled TfInstanceLoggingConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#log_group TfInstanceLoggingConfiguration#log_group}.
	// Experimental.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

