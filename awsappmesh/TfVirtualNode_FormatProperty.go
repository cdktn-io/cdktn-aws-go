package awsappmesh


// Experimental.
type TfVirtualNode_FormatProperty struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#json TfVirtualNode#json}
	// Experimental.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#text TfVirtualNode#text}.
	// Experimental.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

