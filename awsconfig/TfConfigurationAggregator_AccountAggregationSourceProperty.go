package awsconfig


// Experimental.
type TfConfigurationAggregator_AccountAggregationSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_aggregator#account_ids TfConfigurationAggregator#account_ids}.
	// Experimental.
	AccountIds *[]*string `field:"required" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_aggregator#all_regions TfConfigurationAggregator#all_regions}.
	// Experimental.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_aggregator#regions TfConfigurationAggregator#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

