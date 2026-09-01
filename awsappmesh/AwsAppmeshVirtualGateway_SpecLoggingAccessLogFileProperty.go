package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecLoggingAccessLogFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#path AwsAppmeshVirtualGateway#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#format AwsAppmeshVirtualGateway#format}
	// Experimental.
	Format *AwsAppmeshVirtualGateway_FormatProperty `field:"optional" json:"format" yaml:"format"`
}

