package awsec2imagebuilder


// Experimental.
type TfDistributionConfiguration_DistributionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#region TfDistributionConfiguration#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// ami_distribution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#ami_distribution_configuration TfDistributionConfiguration#ami_distribution_configuration}
	// Experimental.
	AmiDistributionConfiguration *TfDistributionConfiguration_AmiDistributionConfigurationProperty `field:"optional" json:"amiDistributionConfiguration" yaml:"amiDistributionConfiguration"`
	// container_distribution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#container_distribution_configuration TfDistributionConfiguration#container_distribution_configuration}
	// Experimental.
	ContainerDistributionConfiguration *TfDistributionConfiguration_ContainerDistributionConfigurationProperty `field:"optional" json:"containerDistributionConfiguration" yaml:"containerDistributionConfiguration"`
	// fast_launch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#fast_launch_configuration TfDistributionConfiguration#fast_launch_configuration}
	// Experimental.
	FastLaunchConfiguration interface{} `field:"optional" json:"fastLaunchConfiguration" yaml:"fastLaunchConfiguration"`
	// launch_template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#launch_template_configuration TfDistributionConfiguration#launch_template_configuration}
	// Experimental.
	LaunchTemplateConfiguration interface{} `field:"optional" json:"launchTemplateConfiguration" yaml:"launchTemplateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#license_configuration_arns TfDistributionConfiguration#license_configuration_arns}.
	// Experimental.
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// s3_export_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#s3_export_configuration TfDistributionConfiguration#s3_export_configuration}
	// Experimental.
	S3ExportConfiguration *TfDistributionConfiguration_S3ExportConfigurationProperty `field:"optional" json:"s3ExportConfiguration" yaml:"s3ExportConfiguration"`
	// ssm_parameter_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#ssm_parameter_configuration TfDistributionConfiguration#ssm_parameter_configuration}
	// Experimental.
	SsmParameterConfiguration interface{} `field:"optional" json:"ssmParameterConfiguration" yaml:"ssmParameterConfiguration"`
}

