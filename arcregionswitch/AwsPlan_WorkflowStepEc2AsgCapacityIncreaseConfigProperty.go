package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepEc2AsgCapacityIncreaseConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#capacity_monitoring_approach AwsPlan#capacity_monitoring_approach}.
	// Experimental.
	CapacityMonitoringApproach *string `field:"required" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// asg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#asg AwsPlan#asg}
	// Experimental.
	Asg interface{} `field:"optional" json:"asg" yaml:"asg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#target_percent AwsPlan#target_percent}.
	// Experimental.
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes AwsPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ungraceful AwsPlan#ungraceful}
	// Experimental.
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

