package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepEc2AsgCapacityIncreaseConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#capacity_monitoring_approach TfPlan#capacity_monitoring_approach}.
	// Experimental.
	CapacityMonitoringApproach *string `field:"required" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// asg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#asg TfPlan#asg}
	// Experimental.
	Asg interface{} `field:"optional" json:"asg" yaml:"asg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#target_percent TfPlan#target_percent}.
	// Experimental.
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes TfPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ungraceful TfPlan#ungraceful}
	// Experimental.
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

