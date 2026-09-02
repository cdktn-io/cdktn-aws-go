package awsappmesh


// Experimental.
type TfVirtualNode_SpecLoggingAccessLogFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#path TfVirtualNode#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#format TfVirtualNode#format}
	// Experimental.
	Format *TfVirtualNode_FormatProperty `field:"optional" json:"format" yaml:"format"`
}

