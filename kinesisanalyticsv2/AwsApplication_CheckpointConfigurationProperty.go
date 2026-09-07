package kinesisanalyticsv2


// Experimental.
type AwsApplication_CheckpointConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#configuration_type AwsApplication#configuration_type}.
	// Experimental.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#checkpointing_enabled AwsApplication#checkpointing_enabled}.
	// Experimental.
	CheckpointingEnabled interface{} `field:"optional" json:"checkpointingEnabled" yaml:"checkpointingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#checkpoint_interval AwsApplication#checkpoint_interval}.
	// Experimental.
	CheckpointInterval *float64 `field:"optional" json:"checkpointInterval" yaml:"checkpointInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#min_pause_between_checkpoints AwsApplication#min_pause_between_checkpoints}.
	// Experimental.
	MinPauseBetweenCheckpoints *float64 `field:"optional" json:"minPauseBetweenCheckpoints" yaml:"minPauseBetweenCheckpoints"`
}

