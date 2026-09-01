package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_AccessLogProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file AwsAppmeshVirtualGateway#file}
	// Experimental.
	File *AwsAppmeshVirtualGateway_SpecLoggingAccessLogFileProperty `field:"optional" json:"file" yaml:"file"`
}

