package awsappmesh


// Experimental.
type TfRoute_TcpRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action TfRoute#action}
	// Experimental.
	Action *TfRoute_SpecTcpRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match TfRoute#match}
	// Experimental.
	Match *TfRoute_SpecTcpRouteMatchProperty `field:"optional" json:"match" yaml:"match"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout TfRoute#timeout}
	// Experimental.
	Timeout *TfRoute_SpecTcpRouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

