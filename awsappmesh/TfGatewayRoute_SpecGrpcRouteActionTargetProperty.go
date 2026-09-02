package awsappmesh


// Experimental.
type TfGatewayRoute_SpecGrpcRouteActionTargetProperty struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#virtual_service TfGatewayRoute#virtual_service}
	// Experimental.
	VirtualService *TfGatewayRoute_SpecGrpcRouteActionTargetVirtualServiceProperty `field:"required" json:"virtualService" yaml:"virtualService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#port TfGatewayRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

