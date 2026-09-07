package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepParallelConfigStepExecutionApprovalConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#approval_role AwsPlan#approval_role}.
	// Experimental.
	ApprovalRole *string `field:"required" json:"approvalRole" yaml:"approvalRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes AwsPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

