package awslambda


// Experimental.
type TfProvisionedConcurrencyConfig_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_provisioned_concurrency_config#create TfProvisionedConcurrencyConfig#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_provisioned_concurrency_config#update TfProvisionedConcurrencyConfig#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

