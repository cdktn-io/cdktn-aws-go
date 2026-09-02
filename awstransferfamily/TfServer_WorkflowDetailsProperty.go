package awstransferfamily


// Experimental.
type TfServer_WorkflowDetailsProperty struct {
	// on_partial_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#on_partial_upload TfServer#on_partial_upload}
	// Experimental.
	OnPartialUpload *TfServer_OnPartialUploadProperty `field:"optional" json:"onPartialUpload" yaml:"onPartialUpload"`
	// on_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server#on_upload TfServer#on_upload}
	// Experimental.
	OnUpload *TfServer_OnUploadProperty `field:"optional" json:"onUpload" yaml:"onUpload"`
}

