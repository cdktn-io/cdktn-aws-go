package awsappmesh


// Experimental.
type TfVirtualGateway_FormatProperty struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#json TfVirtualGateway#json}
	// Experimental.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#text TfVirtualGateway#text}.
	// Experimental.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

