package awsappmesh


// Experimental.
type AwsAppmeshRoute_SpecProperty struct {
	// grpc_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#grpc_route AwsAppmeshRoute#grpc_route}
	// Experimental.
	GrpcRoute *AwsAppmeshRoute_GrpcRouteProperty `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// http2_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#http2_route AwsAppmeshRoute#http2_route}
	// Experimental.
	Http2Route *AwsAppmeshRoute_Http2RouteProperty `field:"optional" json:"http2Route" yaml:"http2Route"`
	// http_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#http_route AwsAppmeshRoute#http_route}
	// Experimental.
	HttpRoute *AwsAppmeshRoute_HttpRouteProperty `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#priority AwsAppmeshRoute#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// tcp_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#tcp_route AwsAppmeshRoute#tcp_route}
	// Experimental.
	TcpRoute *AwsAppmeshRoute_TcpRouteProperty `field:"optional" json:"tcpRoute" yaml:"tcpRoute"`
}

