package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_ProcessingConfigurationProperty struct {
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#lambda AwsKinesisAnalyticsApplication#lambda}
	// Experimental.
	Lambda *AwsKinesisAnalyticsApplication_InputsProcessingConfigurationLambdaProperty `field:"required" json:"lambda" yaml:"lambda"`
}

