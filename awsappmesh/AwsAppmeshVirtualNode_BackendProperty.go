package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_BackendProperty struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#virtual_service AwsAppmeshVirtualNode#virtual_service}
	// Experimental.
	VirtualService *AwsAppmeshVirtualNode_VirtualServiceProperty `field:"required" json:"virtualService" yaml:"virtualService"`
}

