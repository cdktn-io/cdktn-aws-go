package awsmsk


// Experimental.
type AwsMskReplicator_ConsumerGroupReplicationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#consumer_groups_to_replicate AwsMskReplicator#consumer_groups_to_replicate}.
	// Experimental.
	ConsumerGroupsToReplicate *[]*string `field:"required" json:"consumerGroupsToReplicate" yaml:"consumerGroupsToReplicate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#consumer_group_offset_sync_mode AwsMskReplicator#consumer_group_offset_sync_mode}.
	// Experimental.
	ConsumerGroupOffsetSyncMode *string `field:"optional" json:"consumerGroupOffsetSyncMode" yaml:"consumerGroupOffsetSyncMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#consumer_groups_to_exclude AwsMskReplicator#consumer_groups_to_exclude}.
	// Experimental.
	ConsumerGroupsToExclude *[]*string `field:"optional" json:"consumerGroupsToExclude" yaml:"consumerGroupsToExclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#detect_and_copy_new_consumer_groups AwsMskReplicator#detect_and_copy_new_consumer_groups}.
	// Experimental.
	DetectAndCopyNewConsumerGroups interface{} `field:"optional" json:"detectAndCopyNewConsumerGroups" yaml:"detectAndCopyNewConsumerGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#synchronise_consumer_group_offsets AwsMskReplicator#synchronise_consumer_group_offsets}.
	// Experimental.
	SynchroniseConsumerGroupOffsets interface{} `field:"optional" json:"synchroniseConsumerGroupOffsets" yaml:"synchroniseConsumerGroupOffsets"`
}

