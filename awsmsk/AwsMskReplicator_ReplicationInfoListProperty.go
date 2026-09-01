package awsmsk


// Experimental.
type AwsMskReplicator_ReplicationInfoListProperty struct {
	// consumer_group_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#consumer_group_replication AwsMskReplicator#consumer_group_replication}
	// Experimental.
	ConsumerGroupReplication interface{} `field:"required" json:"consumerGroupReplication" yaml:"consumerGroupReplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#source_kafka_cluster_arn AwsMskReplicator#source_kafka_cluster_arn}.
	// Experimental.
	SourceKafkaClusterArn *string `field:"required" json:"sourceKafkaClusterArn" yaml:"sourceKafkaClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#target_compression_type AwsMskReplicator#target_compression_type}.
	// Experimental.
	TargetCompressionType *string `field:"required" json:"targetCompressionType" yaml:"targetCompressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#target_kafka_cluster_arn AwsMskReplicator#target_kafka_cluster_arn}.
	// Experimental.
	TargetKafkaClusterArn *string `field:"required" json:"targetKafkaClusterArn" yaml:"targetKafkaClusterArn"`
	// topic_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#topic_replication AwsMskReplicator#topic_replication}
	// Experimental.
	TopicReplication interface{} `field:"required" json:"topicReplication" yaml:"topicReplication"`
}

