package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_VirtualServiceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#virtual_service_name AwsAppmeshVirtualNode#virtual_service_name}.
	// Experimental.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#client_policy AwsAppmeshVirtualNode#client_policy}
	// Experimental.
	ClientPolicy *AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyProperty `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

