package awsappmesh


// Experimental.
type TfRoute_SpecTcpRouteActionWeightedTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#virtual_node TfRoute#virtual_node}.
	// Experimental.
	VirtualNode *string `field:"required" json:"virtualNode" yaml:"virtualNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#weight TfRoute#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#port TfRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

