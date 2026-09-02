package awsappmesh


// Experimental.
type TfGatewayRoute_SpecHttpRouteActionTargetProperty struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#virtual_service TfGatewayRoute#virtual_service}
	// Experimental.
	VirtualService *TfGatewayRoute_SpecHttpRouteActionTargetVirtualServiceProperty `field:"required" json:"virtualService" yaml:"virtualService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#port TfGatewayRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

