package awsappmesh


// Experimental.
type TfVirtualService_ProviderProperty struct {
	// virtual_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_service#virtual_node TfVirtualService#virtual_node}
	// Experimental.
	VirtualNode *TfVirtualService_VirtualNodeProperty `field:"optional" json:"virtualNode" yaml:"virtualNode"`
	// virtual_router block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_service#virtual_router TfVirtualService#virtual_router}
	// Experimental.
	VirtualRouter *TfVirtualService_VirtualRouterProperty `field:"optional" json:"virtualRouter" yaml:"virtualRouter"`
}

