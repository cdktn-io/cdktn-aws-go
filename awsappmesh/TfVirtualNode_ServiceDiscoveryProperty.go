package awsappmesh


// Experimental.
type TfVirtualNode_ServiceDiscoveryProperty struct {
	// aws_cloud_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#aws_cloud_map TfVirtualNode#aws_cloud_map}
	// Experimental.
	AwsCloudMap *TfVirtualNode_AwsCloudMapProperty `field:"optional" json:"awsCloudMap" yaml:"awsCloudMap"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#dns TfVirtualNode#dns}
	// Experimental.
	Dns *TfVirtualNode_DnsProperty `field:"optional" json:"dns" yaml:"dns"`
}

