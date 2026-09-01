package awsappmesh


// Experimental.
type AwsAppmeshGatewayRoute_SpecProperty struct {
	// grpc_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#grpc_route AwsAppmeshGatewayRoute#grpc_route}
	// Experimental.
	GrpcRoute *AwsAppmeshGatewayRoute_GrpcRouteProperty `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// http2_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#http2_route AwsAppmeshGatewayRoute#http2_route}
	// Experimental.
	Http2Route *AwsAppmeshGatewayRoute_Http2RouteProperty `field:"optional" json:"http2Route" yaml:"http2Route"`
	// http_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#http_route AwsAppmeshGatewayRoute#http_route}
	// Experimental.
	HttpRoute *AwsAppmeshGatewayRoute_HttpRouteProperty `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_gateway_route#priority AwsAppmeshGatewayRoute#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

