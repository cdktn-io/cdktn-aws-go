package awsconfig


// Experimental.
type TfConfigurationAggregator_OrganizationAggregationSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_aggregator#role_arn TfConfigurationAggregator#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_aggregator#all_regions TfConfigurationAggregator#all_regions}.
	// Experimental.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_aggregator#regions TfConfigurationAggregator#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

