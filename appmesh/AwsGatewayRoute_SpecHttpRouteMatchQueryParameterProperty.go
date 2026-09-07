package appmesh


// Experimental.
type AwsGatewayRoute_SpecHttpRouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#name AwsGatewayRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match AwsGatewayRoute#match}
	// Experimental.
	Match *AwsGatewayRoute_SpecHttpRouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

