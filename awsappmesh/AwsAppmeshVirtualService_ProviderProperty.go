package awsappmesh


// Experimental.
type AwsAppmeshVirtualService_ProviderProperty struct {
	// virtual_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_service#virtual_node AwsAppmeshVirtualService#virtual_node}
	// Experimental.
	VirtualNode *AwsAppmeshVirtualService_VirtualNodeProperty `field:"optional" json:"virtualNode" yaml:"virtualNode"`
	// virtual_router block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_service#virtual_router AwsAppmeshVirtualService#virtual_router}
	// Experimental.
	VirtualRouter *AwsAppmeshVirtualService_VirtualRouterProperty `field:"optional" json:"virtualRouter" yaml:"virtualRouter"`
}

