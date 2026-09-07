package msk


// Experimental.
type AwsReplicator_KafkaClusterProperty struct {
	// amazon_msk_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#amazon_msk_cluster AwsReplicator#amazon_msk_cluster}
	// Experimental.
	AmazonMskCluster *AwsReplicator_AmazonMskClusterProperty `field:"required" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#vpc_config AwsReplicator#vpc_config}
	// Experimental.
	VpcConfig *AwsReplicator_VpcConfigProperty `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
}

