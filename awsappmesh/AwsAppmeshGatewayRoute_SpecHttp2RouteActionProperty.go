package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecHttp2RouteActionProperty struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#target AwsAppmeshGatewayRoute#target}
	// Experimental.
	Target *AwsAppmeshGatewayRoute_SpecHttp2RouteActionTargetProperty `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#rewrite AwsAppmeshGatewayRoute#rewrite}
	// Experimental.
	Rewrite *AwsAppmeshGatewayRoute_SpecHttp2RouteActionRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

