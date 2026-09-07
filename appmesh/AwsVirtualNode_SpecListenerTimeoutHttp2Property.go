package appmesh


// Experimental.
type AwsVirtualNode_SpecListenerTimeoutHttp2Property struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle AwsVirtualNode#idle}
	// Experimental.
	Idle *AwsVirtualNode_SpecListenerTimeoutHttp2IdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#per_request AwsVirtualNode#per_request}
	// Experimental.
	PerRequest *AwsVirtualNode_SpecListenerTimeoutHttp2PerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

