package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttp2RouteActionProperty struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#target TfGatewayRoute#target}
	// Experimental.
	Target *TfGatewayRoute_SpecHttp2RouteActionTargetProperty `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#rewrite TfGatewayRoute#rewrite}
	// Experimental.
	Rewrite *TfGatewayRoute_SpecHttp2RouteActionRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

