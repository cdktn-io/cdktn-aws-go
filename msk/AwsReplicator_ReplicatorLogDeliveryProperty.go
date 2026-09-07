package msk


// Experimental.
type AwsReplicator_ReplicatorLogDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#cloudwatch_logs AwsReplicator#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsReplicator_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#firehose AwsReplicator#firehose}
	// Experimental.
	Firehose *AwsReplicator_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#s3 AwsReplicator#s3}
	// Experimental.
	S3 *AwsReplicator_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

