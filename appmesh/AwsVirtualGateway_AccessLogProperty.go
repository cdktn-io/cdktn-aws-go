package appmesh


// Experimental.
type AwsVirtualGateway_AccessLogProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file AwsVirtualGateway#file}
	// Experimental.
	File *AwsVirtualGateway_SpecLoggingAccessLogFileProperty `field:"optional" json:"file" yaml:"file"`
}

