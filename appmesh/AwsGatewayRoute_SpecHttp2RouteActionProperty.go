package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttp2RouteActionProperty struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#target AwsGatewayRoute#target}
	// Experimental.
	Target *AwsGatewayRoute_SpecHttp2RouteActionTargetProperty `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#rewrite AwsGatewayRoute#rewrite}
	// Experimental.
	Rewrite *AwsGatewayRoute_SpecHttp2RouteActionRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

