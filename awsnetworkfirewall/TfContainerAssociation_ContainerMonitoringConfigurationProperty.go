package awsnetworkfirewall


// Experimental.
type TfContainerAssociation_ContainerMonitoringConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#cluster_arn TfContainerAssociation#cluster_arn}.
	// Experimental.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// attribute_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_container_association#attribute_filter TfContainerAssociation#attribute_filter}
	// Experimental.
	AttributeFilter interface{} `field:"optional" json:"attributeFilter" yaml:"attributeFilter"`
}

