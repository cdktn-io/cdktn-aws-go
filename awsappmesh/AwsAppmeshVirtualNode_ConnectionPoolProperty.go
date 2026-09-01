package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_ConnectionPoolProperty struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#grpc AwsAppmeshVirtualNode#grpc}
	// Experimental.
	Grpc *AwsAppmeshVirtualNode_SpecListenerConnectionPoolGrpcProperty `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http AwsAppmeshVirtualNode#http}
	// Experimental.
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#http2 AwsAppmeshVirtualNode#http2}
	// Experimental.
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// tcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tcp AwsAppmeshVirtualNode#tcp}
	// Experimental.
	Tcp interface{} `field:"optional" json:"tcp" yaml:"tcp"`
}

