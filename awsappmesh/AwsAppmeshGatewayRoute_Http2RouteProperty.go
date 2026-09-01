package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_Http2RouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#action AwsAppmeshGatewayRoute#action}
	// Experimental.
	Action *AwsAppmeshGatewayRoute_SpecHttp2RouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match AwsAppmeshGatewayRoute#match}
	// Experimental.
	Match *AwsAppmeshGatewayRoute_SpecHttp2RouteMatchProperty `field:"required" json:"match" yaml:"match"`
}

