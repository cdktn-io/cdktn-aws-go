package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallContainerAssociation_ContainerMonitoringConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#cluster_arn AwsNetworkfirewallContainerAssociation#cluster_arn}.
	// Experimental.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// attribute_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#attribute_filter AwsNetworkfirewallContainerAssociation#attribute_filter}
	// Experimental.
	AttributeFilter interface{} `field:"optional" json:"attributeFilter" yaml:"attributeFilter"`
}

