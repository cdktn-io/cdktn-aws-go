package awsappmesh


// Experimental.
type TfVirtualGateway_AccessLogProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#file TfVirtualGateway#file}
	// Experimental.
	File *TfVirtualGateway_SpecLoggingAccessLogFileProperty `field:"optional" json:"file" yaml:"file"`
}

