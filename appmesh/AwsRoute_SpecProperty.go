package appmesh


// Experimental.
type AwsRoute_SpecProperty struct {
	// grpc_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#grpc_route AwsRoute#grpc_route}
	// Experimental.
	GrpcRoute *AwsRoute_GrpcRouteProperty `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// http2_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#http2_route AwsRoute#http2_route}
	// Experimental.
	Http2Route *AwsRoute_Http2RouteProperty `field:"optional" json:"http2Route" yaml:"http2Route"`
	// http_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#http_route AwsRoute#http_route}
	// Experimental.
	HttpRoute *AwsRoute_HttpRouteProperty `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#priority AwsRoute#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// tcp_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#tcp_route AwsRoute#tcp_route}
	// Experimental.
	TcpRoute *AwsRoute_TcpRouteProperty `field:"optional" json:"tcpRoute" yaml:"tcpRoute"`
}

