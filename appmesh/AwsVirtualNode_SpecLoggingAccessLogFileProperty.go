package appmesh


// Experimental.
type AwsVirtualNode_SpecLoggingAccessLogFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#path AwsVirtualNode#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#format AwsVirtualNode#format}
	// Experimental.
	Format *AwsVirtualNode_FormatProperty `field:"optional" json:"format" yaml:"format"`
}

