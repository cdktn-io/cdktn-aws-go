package msk


// Experimental.
type AwsCluster_BrokerLogsProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#cloudwatch_logs AwsCluster#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *AwsCluster_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#firehose AwsCluster#firehose}
	// Experimental.
	Firehose *AwsCluster_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#s3 AwsCluster#s3}
	// Experimental.
	S3 *AwsCluster_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

