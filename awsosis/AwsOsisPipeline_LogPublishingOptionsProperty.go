package awsosis


// Experimental.
type AwsOsisPipeline_LogPublishingOptionsProperty struct {
	// cloudwatch_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#cloudwatch_log_destination AwsOsisPipeline#cloudwatch_log_destination}
	// Experimental.
	CloudwatchLogDestination interface{} `field:"optional" json:"cloudwatchLogDestination" yaml:"cloudwatchLogDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#is_logging_enabled AwsOsisPipeline#is_logging_enabled}.
	// Experimental.
	IsLoggingEnabled interface{} `field:"optional" json:"isLoggingEnabled" yaml:"isLoggingEnabled"`
}

