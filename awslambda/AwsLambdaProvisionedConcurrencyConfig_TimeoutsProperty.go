package awslambda


// Experimental.
type AwsLambdaProvisionedConcurrencyConfig_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_provisioned_concurrency_config#create AwsLambdaProvisionedConcurrencyConfig#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_provisioned_concurrency_config#update AwsLambdaProvisionedConcurrencyConfig#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

