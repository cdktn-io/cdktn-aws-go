package awsappmesh


// Experimental.
type TfRoute_SpecProperty struct {
	// grpc_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#grpc_route TfRoute#grpc_route}
	// Experimental.
	GrpcRoute *TfRoute_GrpcRouteProperty `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// http2_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#http2_route TfRoute#http2_route}
	// Experimental.
	Http2Route *TfRoute_Http2RouteProperty `field:"optional" json:"http2Route" yaml:"http2Route"`
	// http_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#http_route TfRoute#http_route}
	// Experimental.
	HttpRoute *TfRoute_HttpRouteProperty `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#priority TfRoute#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// tcp_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#tcp_route TfRoute#tcp_route}
	// Experimental.
	TcpRoute *TfRoute_TcpRouteProperty `field:"optional" json:"tcpRoute" yaml:"tcpRoute"`
}

