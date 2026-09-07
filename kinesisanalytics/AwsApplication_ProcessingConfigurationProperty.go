package kinesisanalytics


// Experimental.
type AwsApplication_ProcessingConfigurationProperty struct {
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#lambda AwsApplication#lambda}
	// Experimental.
	Lambda *AwsApplication_InputsProcessingConfigurationLambdaProperty `field:"required" json:"lambda" yaml:"lambda"`
}

