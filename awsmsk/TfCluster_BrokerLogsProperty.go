package awsmsk


// Experimental.
type TfCluster_BrokerLogsProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#cloudwatch_logs TfCluster#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs *TfCluster_CloudwatchLogsProperty `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#firehose TfCluster#firehose}
	// Experimental.
	Firehose *TfCluster_FirehoseProperty `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#s3 TfCluster#s3}
	// Experimental.
	S3 *TfCluster_S3Property `field:"optional" json:"s3" yaml:"s3"`
}

