package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecGrpcRouteActionTargetProperty struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#virtual_service AwsAppmeshGatewayRoute#virtual_service}
	// Experimental.
	VirtualService *AwsAppmeshGatewayRoute_SpecGrpcRouteActionTargetVirtualServiceProperty `field:"required" json:"virtualService" yaml:"virtualService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#port AwsAppmeshGatewayRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

