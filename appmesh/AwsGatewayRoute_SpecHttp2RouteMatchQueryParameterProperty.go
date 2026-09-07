package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttp2RouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#name AwsGatewayRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match AwsGatewayRoute#match}
	// Experimental.
	Match *AwsGatewayRoute_SpecHttp2RouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

