package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region_to_run AwsPlan#region_to_run}.
	// Experimental.
	RegionToRun *string `field:"required" json:"regionToRun" yaml:"regionToRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#retry_interval_minutes AwsPlan#retry_interval_minutes}.
	// Experimental.
	RetryIntervalMinutes *float64 `field:"required" json:"retryIntervalMinutes" yaml:"retryIntervalMinutes"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#lambda AwsPlan#lambda}
	// Experimental.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes AwsPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ungraceful AwsPlan#ungraceful}
	// Experimental.
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

