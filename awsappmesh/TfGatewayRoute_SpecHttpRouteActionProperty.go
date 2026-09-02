package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttpRouteActionProperty struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#target TfGatewayRoute#target}
	// Experimental.
	Target *TfGatewayRoute_SpecHttpRouteActionTargetProperty `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#rewrite TfGatewayRoute#rewrite}
	// Experimental.
	Rewrite *TfGatewayRoute_SpecHttpRouteActionRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

