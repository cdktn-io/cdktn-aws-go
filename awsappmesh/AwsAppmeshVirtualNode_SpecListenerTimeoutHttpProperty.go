package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecListenerTimeoutHttpProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle AwsAppmeshVirtualNode#idle}
	// Experimental.
	Idle *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#per_request AwsAppmeshVirtualNode#per_request}
	// Experimental.
	PerRequest *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

