package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle AwsAppmeshVirtualNode#idle}
	// Experimental.
	Idle *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#per_request AwsAppmeshVirtualNode#per_request}
	// Experimental.
	PerRequest *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

