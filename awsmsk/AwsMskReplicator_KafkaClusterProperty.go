package awsmsk


// Experimental.
type AwsMskReplicator_KafkaClusterProperty struct {
	// amazon_msk_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#amazon_msk_cluster AwsMskReplicator#amazon_msk_cluster}
	// Experimental.
	AmazonMskCluster *AwsMskReplicator_AmazonMskClusterProperty `field:"required" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#vpc_config AwsMskReplicator#vpc_config}
	// Experimental.
	VpcConfig *AwsMskReplicator_VpcConfigProperty `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
}

