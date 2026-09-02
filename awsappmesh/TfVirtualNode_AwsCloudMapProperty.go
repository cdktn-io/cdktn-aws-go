package awsappmesh


// Experimental.
type TfVirtualNode_AwsCloudMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#namespace_name TfVirtualNode#namespace_name}.
	// Experimental.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#service_name TfVirtualNode#service_name}.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#attributes TfVirtualNode#attributes}.
	// Experimental.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

