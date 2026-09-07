package appmesh


// Experimental.
type AwsVirtualNode_ConnectionPoolProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#grpc AwsVirtualNode#grpc}
	// Experimental.
	Grpc *AwsVirtualNode_SpecListenerConnectionPoolGrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http AwsVirtualNode#http}
	// Experimental.
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http2 AwsVirtualNode#http2}
	// Experimental.
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// tcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tcp AwsVirtualNode#tcp}
	// Experimental.
	Tcp interface{} `field:"optional" json:"tcp" yaml:"tcp"`
}

