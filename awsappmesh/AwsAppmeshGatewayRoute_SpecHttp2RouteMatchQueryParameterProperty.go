package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecHttp2RouteMatchQueryParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#name AwsAppmeshGatewayRoute#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match AwsAppmeshGatewayRoute#match}
	// Experimental.
	Match *AwsAppmeshGatewayRoute_SpecHttp2RouteMatchQueryParameterMatchProperty `field:"optional" json:"match" yaml:"match"`
}

