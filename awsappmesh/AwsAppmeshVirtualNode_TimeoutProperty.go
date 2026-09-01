package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_TimeoutProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#grpc AwsAppmeshVirtualNode#grpc}
	// Experimental.
	Grpc *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http AwsAppmeshVirtualNode#http}
	// Experimental.
	Http *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpProperty `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http2 AwsAppmeshVirtualNode#http2}
	// Experimental.
	Http2 *AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2Property `field:"optional" json:"http2" yaml:"http2"`
	// tcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tcp AwsAppmeshVirtualNode#tcp}
	// Experimental.
	Tcp *AwsAppmeshVirtualNode_SpecListenerTimeoutTcpProperty `field:"optional" json:"tcp" yaml:"tcp"`
}

