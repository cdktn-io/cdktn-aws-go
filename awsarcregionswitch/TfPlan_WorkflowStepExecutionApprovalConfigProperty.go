package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepExecutionApprovalConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#approval_role TfPlan#approval_role}.
	// Experimental.
	ApprovalRole *string `field:"required" json:"approvalRole" yaml:"approvalRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes TfPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

