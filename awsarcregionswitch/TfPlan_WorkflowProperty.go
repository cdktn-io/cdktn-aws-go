package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#workflow_target_action TfPlan#workflow_target_action}.
	// Experimental.
	WorkflowTargetAction *string `field:"required" json:"workflowTargetAction" yaml:"workflowTargetAction"`
	// step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#step TfPlan#step}
	// Experimental.
	Step interface{} `field:"optional" json:"step" yaml:"step"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#workflow_description TfPlan#workflow_description}.
	// Experimental.
	WorkflowDescription *string `field:"optional" json:"workflowDescription" yaml:"workflowDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#workflow_target_region TfPlan#workflow_target_region}.
	// Experimental.
	WorkflowTargetRegion *string `field:"optional" json:"workflowTargetRegion" yaml:"workflowTargetRegion"`
}

