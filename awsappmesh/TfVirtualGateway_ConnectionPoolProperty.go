package awsappmesh


// Experimental.
type TfVirtualGateway_ConnectionPoolProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#grpc TfVirtualGateway#grpc}
	// Experimental.
	Grpc *TfVirtualGateway_GrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#http TfVirtualGateway#http}
	// Experimental.
	Http *TfVirtualGateway_HttpProperty `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#http2 TfVirtualGateway#http2}
	// Experimental.
	Http2 *TfVirtualGateway_Http2Property `field:"optional" json:"http2" yaml:"http2"`
}

