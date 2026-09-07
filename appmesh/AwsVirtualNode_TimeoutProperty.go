package appmesh


// Experimental.
type AwsVirtualNode_TimeoutProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#grpc AwsVirtualNode#grpc}
	// Experimental.
	Grpc *AwsVirtualNode_SpecListenerTimeoutGrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http AwsVirtualNode#http}
	// Experimental.
	Http *AwsVirtualNode_SpecListenerTimeoutHttpProperty `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http2 AwsVirtualNode#http2}
	// Experimental.
	Http2 *AwsVirtualNode_SpecListenerTimeoutHttp2Property `field:"optional" json:"http2" yaml:"http2"`
	// tcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tcp AwsVirtualNode#tcp}
	// Experimental.
	Tcp *AwsVirtualNode_SpecListenerTimeoutTcpProperty `field:"optional" json:"tcp" yaml:"tcp"`
}

