package transferfamily


// Experimental.
type AwsServer_WorkflowDetailsProperty struct {
	// on_partial_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#on_partial_upload AwsServer#on_partial_upload}
	// Experimental.
	OnPartialUpload *AwsServer_OnPartialUploadProperty `field:"optional" json:"onPartialUpload" yaml:"onPartialUpload"`
	// on_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#on_upload AwsServer#on_upload}
	// Experimental.
	OnUpload *AwsServer_OnUploadProperty `field:"optional" json:"onUpload" yaml:"onUpload"`
}

