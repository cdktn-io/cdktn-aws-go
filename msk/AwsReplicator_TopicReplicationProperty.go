package msk


// Experimental.
type AwsReplicator_TopicReplicationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#topics_to_replicate AwsReplicator#topics_to_replicate}.
	// Experimental.
	TopicsToReplicate *[]*string `field:"required" json:"topicsToReplicate" yaml:"topicsToReplicate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#copy_access_control_lists_for_topics AwsReplicator#copy_access_control_lists_for_topics}.
	// Experimental.
	CopyAccessControlListsForTopics interface{} `field:"optional" json:"copyAccessControlListsForTopics" yaml:"copyAccessControlListsForTopics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#copy_topic_configurations AwsReplicator#copy_topic_configurations}.
	// Experimental.
	CopyTopicConfigurations interface{} `field:"optional" json:"copyTopicConfigurations" yaml:"copyTopicConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#detect_and_copy_new_topics AwsReplicator#detect_and_copy_new_topics}.
	// Experimental.
	DetectAndCopyNewTopics interface{} `field:"optional" json:"detectAndCopyNewTopics" yaml:"detectAndCopyNewTopics"`
	// starting_position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#starting_position AwsReplicator#starting_position}
	// Experimental.
	StartingPosition *AwsReplicator_StartingPositionProperty `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// topic_name_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#topic_name_configuration AwsReplicator#topic_name_configuration}
	// Experimental.
	TopicNameConfiguration *AwsReplicator_TopicNameConfigurationProperty `field:"optional" json:"topicNameConfiguration" yaml:"topicNameConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#topics_to_exclude AwsReplicator#topics_to_exclude}.
	// Experimental.
	TopicsToExclude *[]*string `field:"optional" json:"topicsToExclude" yaml:"topicsToExclude"`
}

