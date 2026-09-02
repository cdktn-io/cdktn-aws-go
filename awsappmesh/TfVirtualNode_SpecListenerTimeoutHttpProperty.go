package awsappmesh


// Experimental.
type TfVirtualNode_SpecListenerTimeoutHttpProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle TfVirtualNode#idle}
	// Experimental.
	Idle *TfVirtualNode_SpecListenerTimeoutHttpIdleProperty `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#per_request TfVirtualNode#per_request}
	// Experimental.
	PerRequest *TfVirtualNode_SpecListenerTimeoutHttpPerRequestProperty `field:"optional" json:"perRequest" yaml:"perRequest"`
}

