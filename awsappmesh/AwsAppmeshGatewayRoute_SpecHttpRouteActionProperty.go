package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecHttpRouteActionProperty struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#target AwsAppmeshGatewayRoute#target}
	// Experimental.
	Target *AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetProperty `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#rewrite AwsAppmeshGatewayRoute#rewrite}
	// Experimental.
	Rewrite *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

