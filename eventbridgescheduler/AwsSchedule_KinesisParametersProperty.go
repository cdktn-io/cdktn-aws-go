package eventbridgescheduler


// Experimental.
type AwsSchedule_KinesisParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#partition_key AwsSchedule#partition_key}.
	// Experimental.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

