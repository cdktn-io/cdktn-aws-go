package mwaa


// Experimental.
type AwsEnvironment_SchedulerLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#enabled AwsEnvironment#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#log_level AwsEnvironment#log_level}.
	// Experimental.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

