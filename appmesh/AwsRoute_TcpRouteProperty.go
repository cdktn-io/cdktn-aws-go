package appmesh


// Experimental.
type AwsRoute_TcpRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action AwsRoute#action}
	// Experimental.
	Action *AwsRoute_SpecTcpRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsRoute#match}
	// Experimental.
	Match *AwsRoute_SpecTcpRouteMatchProperty `field:"optional" json:"match" yaml:"match"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout AwsRoute#timeout}
	// Experimental.
	Timeout *AwsRoute_SpecTcpRouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

