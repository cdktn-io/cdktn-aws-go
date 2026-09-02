package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region_to_run TfPlan#region_to_run}.
	// Experimental.
	RegionToRun *string `field:"required" json:"regionToRun" yaml:"regionToRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#retry_interval_minutes TfPlan#retry_interval_minutes}.
	// Experimental.
	RetryIntervalMinutes *float64 `field:"required" json:"retryIntervalMinutes" yaml:"retryIntervalMinutes"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#lambda TfPlan#lambda}
	// Experimental.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes TfPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ungraceful TfPlan#ungraceful}
	// Experimental.
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

