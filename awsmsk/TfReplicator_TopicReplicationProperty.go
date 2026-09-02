package awsmsk


// Experimental.
type TfReplicator_TopicReplicationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#topics_to_replicate TfReplicator#topics_to_replicate}.
	// Experimental.
	TopicsToReplicate *[]*string `field:"required" json:"topicsToReplicate" yaml:"topicsToReplicate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#copy_access_control_lists_for_topics TfReplicator#copy_access_control_lists_for_topics}.
	// Experimental.
	CopyAccessControlListsForTopics interface{} `field:"optional" json:"copyAccessControlListsForTopics" yaml:"copyAccessControlListsForTopics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#copy_topic_configurations TfReplicator#copy_topic_configurations}.
	// Experimental.
	CopyTopicConfigurations interface{} `field:"optional" json:"copyTopicConfigurations" yaml:"copyTopicConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#detect_and_copy_new_topics TfReplicator#detect_and_copy_new_topics}.
	// Experimental.
	DetectAndCopyNewTopics interface{} `field:"optional" json:"detectAndCopyNewTopics" yaml:"detectAndCopyNewTopics"`
	// starting_position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#starting_position TfReplicator#starting_position}
	// Experimental.
	StartingPosition *TfReplicator_StartingPositionProperty `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// topic_name_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#topic_name_configuration TfReplicator#topic_name_configuration}
	// Experimental.
	TopicNameConfiguration *TfReplicator_TopicNameConfigurationProperty `field:"optional" json:"topicNameConfiguration" yaml:"topicNameConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#topics_to_exclude TfReplicator#topics_to_exclude}.
	// Experimental.
	TopicsToExclude *[]*string `field:"optional" json:"topicsToExclude" yaml:"topicsToExclude"`
}

