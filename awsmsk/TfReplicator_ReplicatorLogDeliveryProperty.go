package awsmsk


// Experimental.
type TfReplicator_ReplicatorLogDeliveryProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#cloudwatch_logs TfReplicator#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfReplicator_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#firehose TfReplicator#firehose}
	// Experimental.
	Firehose *TfReplicator_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#s3 TfReplicator#s3}
	// Experimental.
	S3 *TfReplicator_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

