package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttp2RouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#name TfGatewayRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match TfGatewayRoute#match}
	// Experimental.
	Match *TfGatewayRoute_SpecHttp2RouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

