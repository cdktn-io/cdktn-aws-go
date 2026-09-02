package awsappmesh


// Experimental.
type TfGatewayRoute_GrpcRouteProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#action TfGatewayRoute#action}
	// Experimental.
	Action *TfGatewayRoute_SpecGrpcRouteActionProperty `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#match TfGatewayRoute#match}
	// Experimental.
	Match *TfGatewayRoute_SpecGrpcRouteMatchProperty `field:"required" json:"match" yaml:"match"`
}

