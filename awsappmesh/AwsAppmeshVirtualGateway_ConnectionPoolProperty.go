package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_ConnectionPoolProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#grpc AwsAppmeshVirtualGateway#grpc}
	// Experimental.
	Grpc *AwsAppmeshVirtualGateway_GrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#http AwsAppmeshVirtualGateway#http}
	// Experimental.
	Http *AwsAppmeshVirtualGateway_HttpProperty `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#http2 AwsAppmeshVirtualGateway#http2}
	// Experimental.
	Http2 *AwsAppmeshVirtualGateway_Http2Property `field:"optional" json:"http2" yaml:"http2"`
}

