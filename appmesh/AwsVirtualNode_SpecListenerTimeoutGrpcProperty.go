package appmesh


// Experimental.
type AwsVirtualNode_SpecListenerTimeoutGrpcProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle AwsVirtualNode#idle}
	// Experimental.
	Idle *AwsVirtualNode_SpecListenerTimeoutGrpcIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#per_request AwsVirtualNode#per_request}
	// Experimental.
	PerRequest *AwsVirtualNode_SpecListenerTimeoutGrpcPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

