package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecListenerTimeoutTcpIdleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#unit AwsAppmeshVirtualNode#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#value AwsAppmeshVirtualNode#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

