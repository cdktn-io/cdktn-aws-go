package awsappmesh


// Experimental.
type TfVirtualGateway_SpecLoggingAccessLogFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#path TfVirtualGateway#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#format TfVirtualGateway#format}
	// Experimental.
	Format *TfVirtualGateway_FormatProperty `field:"optional" json:"format" yaml:"format"`
}

