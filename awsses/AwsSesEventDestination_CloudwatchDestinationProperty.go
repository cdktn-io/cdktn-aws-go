package awsses


// Experimental.
type AwsSesEventDestination_CloudwatchDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#default_value AwsSesEventDestination#default_value}.
	// Experimental.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#dimension_name AwsSesEventDestination#dimension_name}.
	// Experimental.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#value_source AwsSesEventDestination#value_source}.
	// Experimental.
	ValueSource *string `field:"required" json:"valueSource" yaml:"valueSource"`
}

