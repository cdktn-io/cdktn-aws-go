package awscloudformation


// Experimental.
type TfStackSet_OperationPreferencesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#failure_tolerance_count TfStackSet#failure_tolerance_count}.
	// Experimental.
	FailureToleranceCount *float64 `field:"optional" json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#failure_tolerance_percentage TfStackSet#failure_tolerance_percentage}.
	// Experimental.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#max_concurrent_count TfStackSet#max_concurrent_count}.
	// Experimental.
	MaxConcurrentCount *float64 `field:"optional" json:"maxConcurrentCount" yaml:"maxConcurrentCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#max_concurrent_percentage TfStackSet#max_concurrent_percentage}.
	// Experimental.
	MaxConcurrentPercentage *float64 `field:"optional" json:"maxConcurrentPercentage" yaml:"maxConcurrentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#region_concurrency_type TfStackSet#region_concurrency_type}.
	// Experimental.
	RegionConcurrencyType *string `field:"optional" json:"regionConcurrencyType" yaml:"regionConcurrencyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#region_order TfStackSet#region_order}.
	// Experimental.
	RegionOrder *[]*string `field:"optional" json:"regionOrder" yaml:"regionOrder"`
}

