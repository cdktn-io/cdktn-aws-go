package config


// Experimental.
type AwsRemediationConfiguration_SsmControlsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#concurrent_execution_rate_percentage AwsRemediationConfiguration#concurrent_execution_rate_percentage}.
	// Experimental.
	ConcurrentExecutionRatePercentage *float64 `field:"optional" json:"concurrentExecutionRatePercentage" yaml:"concurrentExecutionRatePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_remediation_configuration#error_percentage AwsRemediationConfiguration#error_percentage}.
	// Experimental.
	ErrorPercentage *float64 `field:"optional" json:"errorPercentage" yaml:"errorPercentage"`
}

