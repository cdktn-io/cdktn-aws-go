package endusermessaging


// Experimental.
type AwsApp_LimitsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_app#daily AwsApp#daily}.
	// Experimental.
	Daily *float64 `field:"optional" json:"daily" yaml:"daily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_app#maximum_duration AwsApp#maximum_duration}.
	// Experimental.
	MaximumDuration *float64 `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_app#messages_per_second AwsApp#messages_per_second}.
	// Experimental.
	MessagesPerSecond *float64 `field:"optional" json:"messagesPerSecond" yaml:"messagesPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpoint_app#total AwsApp#total}.
	// Experimental.
	Total *float64 `field:"optional" json:"total" yaml:"total"`
}

