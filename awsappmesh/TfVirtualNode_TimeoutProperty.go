package awsappmesh


// Experimental.
type TfVirtualNode_TimeoutProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#grpc TfVirtualNode#grpc}
	// Experimental.
	Grpc *TfVirtualNode_SpecListenerTimeoutGrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http TfVirtualNode#http}
	// Experimental.
	Http *TfVirtualNode_SpecListenerTimeoutHttpProperty `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http2 TfVirtualNode#http2}
	// Experimental.
	Http2 *TfVirtualNode_SpecListenerTimeoutHttp2Property `field:"optional" json:"http2" yaml:"http2"`
	// tcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tcp TfVirtualNode#tcp}
	// Experimental.
	Tcp *TfVirtualNode_SpecListenerTimeoutTcpProperty `field:"optional" json:"tcp" yaml:"tcp"`
}

