package awsmsk


// Experimental.
type AwsMskCluster_BrokerLogsProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#cloudwatch_logs AwsMskCluster#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsMskCluster_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#firehose AwsMskCluster#firehose}
	// Experimental.
	Firehose *AwsMskCluster_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#s3 AwsMskCluster#s3}
	// Experimental.
	S3 *AwsMskCluster_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

