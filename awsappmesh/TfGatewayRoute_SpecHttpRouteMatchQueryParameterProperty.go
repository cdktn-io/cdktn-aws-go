package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttpRouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#name TfGatewayRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match TfGatewayRoute#match}
	// Experimental.
	Match *TfGatewayRoute_SpecHttpRouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

