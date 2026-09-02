package awsmsk


// Experimental.
type TfReplicator_KafkaClusterProperty struct {
	// amazon_msk_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#amazon_msk_cluster TfReplicator#amazon_msk_cluster}
	// Experimental.
	AmazonMskCluster *TfReplicator_AmazonMskClusterProperty `field:"required" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#vpc_config TfReplicator#vpc_config}
	// Experimental.
	VpcConfig *TfReplicator_VpcConfigProperty `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
}

