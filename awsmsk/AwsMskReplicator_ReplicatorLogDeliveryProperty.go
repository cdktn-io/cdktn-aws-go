package awsmsk


// Experimental.
type AwsMskReplicator_ReplicatorLogDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#cloudwatch_logs AwsMskReplicator#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsMskReplicator_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#firehose AwsMskReplicator#firehose}
	// Experimental.
	Firehose *AwsMskReplicator_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#s3 AwsMskReplicator#s3}
	// Experimental.
	S3 *AwsMskReplicator_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

