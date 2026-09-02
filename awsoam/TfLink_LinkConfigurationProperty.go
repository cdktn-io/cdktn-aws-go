package awsoam


// Experimental.
type TfLink_LinkConfigurationProperty struct {
	// log_group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/oam_link#log_group_configuration TfLink#log_group_configuration}
	// Experimental.
	LogGroupConfiguration *TfLink_LogGroupConfigurationProperty `field:"optional" json:"logGroupConfiguration" yaml:"logGroupConfiguration"`
	// metric_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/oam_link#metric_configuration TfLink#metric_configuration}
	// Experimental.
	MetricConfiguration *TfLink_MetricConfigurationProperty `field:"optional" json:"metricConfiguration" yaml:"metricConfiguration"`
}

