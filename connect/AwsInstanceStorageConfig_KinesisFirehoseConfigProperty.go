package connect


// Experimental.
type AwsInstanceStorageConfig_KinesisFirehoseConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance_storage_config#firehose_arn AwsInstanceStorageConfig#firehose_arn}.
	// Experimental.
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
}

