package config


// Experimental.
type AwsConfigRule_SourceDetailProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#event_source AwsConfigRule#event_source}.
	// Experimental.
	EventSource *string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#maximum_execution_frequency AwsConfigRule#maximum_execution_frequency}.
	// Experimental.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#message_type AwsConfigRule#message_type}.
	// Experimental.
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
}

