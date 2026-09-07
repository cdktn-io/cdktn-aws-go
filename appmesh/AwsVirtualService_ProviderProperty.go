package appmesh


// Experimental.
type AwsVirtualService_ProviderProperty struct {
	// virtual_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_service#virtual_node AwsVirtualService#virtual_node}
	// Experimental.
	VirtualNode *AwsVirtualService_VirtualNodeProperty `field:"optional" json:"virtualNode" yaml:"virtualNode"`
	// virtual_router block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_service#virtual_router AwsVirtualService#virtual_router}
	// Experimental.
	VirtualRouter *AwsVirtualService_VirtualRouterProperty `field:"optional" json:"virtualRouter" yaml:"virtualRouter"`
}

