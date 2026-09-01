package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallLoggingConfiguration_LoggingConfigurationProperty struct {
	// log_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_logging_configuration#log_destination_config AwsNetworkfirewallLoggingConfiguration#log_destination_config}
	// Experimental.
	LogDestinationConfig interface{} `field:"required" json:"logDestinationConfig" yaml:"logDestinationConfig"`
}

