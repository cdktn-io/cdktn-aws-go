package awsappmesh


// Experimental.
type AwsAppmeshRoute_TcpRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#action AwsAppmeshRoute#action}
	// Experimental.
	Action *AwsAppmeshRoute_SpecTcpRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#match AwsAppmeshRoute#match}
	// Experimental.
	Match *AwsAppmeshRoute_SpecTcpRouteMatchProperty `field:"optional" json:"match" yaml:"match"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#timeout AwsAppmeshRoute#timeout}
	// Experimental.
	Timeout *AwsAppmeshRoute_SpecTcpRouteTimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
}

