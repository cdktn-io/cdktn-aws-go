package lambda


// Experimental.
type AwsFunction_DurableConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#execution_timeout AwsFunction#execution_timeout}.
	// Experimental.
	ExecutionTimeout *float64 `field:"required" json:"executionTimeout" yaml:"executionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#retention_period AwsFunction#retention_period}.
	// Experimental.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

