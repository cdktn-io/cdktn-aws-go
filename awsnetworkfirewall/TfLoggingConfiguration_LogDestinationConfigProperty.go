package awsnetworkfirewall


// Experimental.
type TfLoggingConfiguration_LogDestinationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_logging_configuration#log_destination TfLoggingConfiguration#log_destination}.
	// Experimental.
	LogDestination *map[string]*string `field:"required" json:"logDestination" yaml:"logDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_logging_configuration#log_destination_type TfLoggingConfiguration#log_destination_type}.
	// Experimental.
	LogDestinationType *string `field:"required" json:"logDestinationType" yaml:"logDestinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_logging_configuration#log_type TfLoggingConfiguration#log_type}.
	// Experimental.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

