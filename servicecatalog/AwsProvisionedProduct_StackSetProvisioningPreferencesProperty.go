package servicecatalog


// Experimental.
type AwsProvisionedProduct_StackSetProvisioningPreferencesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#accounts AwsProvisionedProduct#accounts}.
	// Experimental.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#failure_tolerance_count AwsProvisionedProduct#failure_tolerance_count}.
	// Experimental.
	FailureToleranceCount *float64 `field:"optional" json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#failure_tolerance_percentage AwsProvisionedProduct#failure_tolerance_percentage}.
	// Experimental.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#max_concurrency_count AwsProvisionedProduct#max_concurrency_count}.
	// Experimental.
	MaxConcurrencyCount *float64 `field:"optional" json:"maxConcurrencyCount" yaml:"maxConcurrencyCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#max_concurrency_percentage AwsProvisionedProduct#max_concurrency_percentage}.
	// Experimental.
	MaxConcurrencyPercentage *float64 `field:"optional" json:"maxConcurrencyPercentage" yaml:"maxConcurrencyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#regions AwsProvisionedProduct#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

