package awscloudformation


// Experimental.
type TfStackSetInstance_OperationPreferencesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#concurrency_mode TfStackSetInstance#concurrency_mode}.
	// Experimental.
	ConcurrencyMode *string `field:"optional" json:"concurrencyMode" yaml:"concurrencyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#failure_tolerance_count TfStackSetInstance#failure_tolerance_count}.
	// Experimental.
	FailureToleranceCount *float64 `field:"optional" json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#failure_tolerance_percentage TfStackSetInstance#failure_tolerance_percentage}.
	// Experimental.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#max_concurrent_count TfStackSetInstance#max_concurrent_count}.
	// Experimental.
	MaxConcurrentCount *float64 `field:"optional" json:"maxConcurrentCount" yaml:"maxConcurrentCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#max_concurrent_percentage TfStackSetInstance#max_concurrent_percentage}.
	// Experimental.
	MaxConcurrentPercentage *float64 `field:"optional" json:"maxConcurrentPercentage" yaml:"maxConcurrentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#region_concurrency_type TfStackSetInstance#region_concurrency_type}.
	// Experimental.
	RegionConcurrencyType *string `field:"optional" json:"regionConcurrencyType" yaml:"regionConcurrencyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#region_order TfStackSetInstance#region_order}.
	// Experimental.
	RegionOrder *[]*string `field:"optional" json:"regionOrder" yaml:"regionOrder"`
}

