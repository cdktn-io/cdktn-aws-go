package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_GrpcRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#action AwsAppmeshGatewayRoute#action}
	// Experimental.
	Action *AwsAppmeshGatewayRoute_SpecGrpcRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match AwsAppmeshGatewayRoute#match}
	// Experimental.
	Match *AwsAppmeshGatewayRoute_SpecGrpcRouteMatchProperty `field:"required" json:"match" yaml:"match"`
}

