package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_FormatProperty struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#json AwsAppmeshVirtualNode#json}
	// Experimental.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#text AwsAppmeshVirtualNode#text}.
	// Experimental.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

