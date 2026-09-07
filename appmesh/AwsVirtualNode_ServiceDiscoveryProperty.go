package appmesh


// Experimental.
type AwsVirtualNode_ServiceDiscoveryProperty struct {
	// aws_cloud_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#aws_cloud_map AwsVirtualNode#aws_cloud_map}
	// Experimental.
	AwsCloudMap *AwsVirtualNode_AwsCloudMapProperty `field:"optional" json:"awsCloudMap" yaml:"awsCloudMap"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#dns AwsVirtualNode#dns}
	// Experimental.
	Dns *AwsVirtualNode_DnsProperty `field:"optional" json:"dns" yaml:"dns"`
}

