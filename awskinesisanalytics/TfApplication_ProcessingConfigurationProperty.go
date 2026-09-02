package awskinesisanalytics


// Experimental.
type TfApplication_ProcessingConfigurationProperty struct {
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#lambda TfApplication#lambda}
	// Experimental.
	Lambda *TfApplication_InputsProcessingConfigurationLambdaProperty `field:"required" json:"lambda" yaml:"lambda"`
}

