package awstransferfamily


// Experimental.
type TfServer_OnPartialUploadProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#execution_role TfServer#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#workflow_id TfServer#workflow_id}.
	// Experimental.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

