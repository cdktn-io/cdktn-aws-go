package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_DnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#hostname AwsAppmeshVirtualNode#hostname}.
	// Experimental.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#ip_preference AwsAppmeshVirtualNode#ip_preference}.
	// Experimental.
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#response_type AwsAppmeshVirtualNode#response_type}.
	// Experimental.
	ResponseType *string `field:"optional" json:"responseType" yaml:"responseType"`
}

