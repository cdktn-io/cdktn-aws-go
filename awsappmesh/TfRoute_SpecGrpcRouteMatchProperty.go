package awsappmesh


// Experimental.
type TfRoute_SpecGrpcRouteMatchProperty struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#metadata TfRoute#metadata}
	// Experimental.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#method_name TfRoute#method_name}.
	// Experimental.
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#port TfRoute#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#prefix TfRoute#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_route#service_name TfRoute#service_name}.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

