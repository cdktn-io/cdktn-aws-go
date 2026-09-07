package appmesh


// Experimental.
type AwsGatewayRoute_SpecGrpcRouteMatchProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#service_name AwsGatewayRoute#service_name}.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#port AwsGatewayRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

