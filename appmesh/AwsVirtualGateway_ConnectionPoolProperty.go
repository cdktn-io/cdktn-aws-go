package appmesh


// Experimental.
type AwsVirtualGateway_ConnectionPoolProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#grpc AwsVirtualGateway#grpc}
	// Experimental.
	Grpc *AwsVirtualGateway_GrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#http AwsVirtualGateway#http}
	// Experimental.
	Http *AwsVirtualGateway_HttpProperty `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#http2 AwsVirtualGateway#http2}
	// Experimental.
	Http2 *AwsVirtualGateway_Http2Property `field:"optional" json:"http2" yaml:"http2"`
}

