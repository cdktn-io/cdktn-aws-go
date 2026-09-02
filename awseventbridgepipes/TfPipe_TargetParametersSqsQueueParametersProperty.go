package awseventbridgepipes


// Experimental.
type TfPipe_TargetParametersSqsQueueParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#message_deduplication_id TfPipe#message_deduplication_id}.
	// Experimental.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#message_group_id TfPipe#message_group_id}.
	// Experimental.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

