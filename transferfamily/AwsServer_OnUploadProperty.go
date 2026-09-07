package transferfamily


// Experimental.
type AwsServer_OnUploadProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#execution_role AwsServer#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#workflow_id AwsServer#workflow_id}.
	// Experimental.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

