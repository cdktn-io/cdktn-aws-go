package awsoam


// Experimental.
type AwsOamLink_LinkConfigurationProperty struct {
	// log_group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/oam_link#log_group_configuration AwsOamLink#log_group_configuration}
	// Experimental.
	LogGroupConfiguration *AwsOamLink_LogGroupConfigurationProperty `field:"optional" json:"logGroupConfiguration" yaml:"logGroupConfiguration"`
	// metric_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/oam_link#metric_configuration AwsOamLink#metric_configuration}
	// Experimental.
	MetricConfiguration *AwsOamLink_MetricConfigurationProperty `field:"optional" json:"metricConfiguration" yaml:"metricConfiguration"`
}

