package awseventbridge


// Experimental.
type AwsCloudwatchEventTarget_KinesisTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#partition_key_path AwsCloudwatchEventTarget#partition_key_path}.
	// Experimental.
	PartitionKeyPath *string `field:"optional" json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

