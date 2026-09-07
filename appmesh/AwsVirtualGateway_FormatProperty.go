package appmesh


// Experimental.
type AwsVirtualGateway_FormatProperty struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#json AwsVirtualGateway#json}
	// Experimental.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#text AwsVirtualGateway#text}.
	// Experimental.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

