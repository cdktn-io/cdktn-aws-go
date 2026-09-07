package appmesh


// Experimental.
type AwsGatewayRoute_HttpRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#action AwsGatewayRoute#action}
	// Experimental.
	Action *AwsGatewayRoute_SpecHttpRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match AwsGatewayRoute#match}
	// Experimental.
	Match *AwsGatewayRoute_SpecHttpRouteMatchProperty `field:"required" json:"match" yaml:"match"`
}

