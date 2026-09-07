package ec2imagebuilder


// Experimental.
type AwsDistributionConfiguration_DistributionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#region AwsDistributionConfiguration#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// ami_distribution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#ami_distribution_configuration AwsDistributionConfiguration#ami_distribution_configuration}
	// Experimental.
	AmiDistributionConfiguration *AwsDistributionConfiguration_AmiDistributionConfigurationProperty `field:"optional" json:"amiDistributionConfiguration" yaml:"amiDistributionConfiguration"`
	// container_distribution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#container_distribution_configuration AwsDistributionConfiguration#container_distribution_configuration}
	// Experimental.
	ContainerDistributionConfiguration *AwsDistributionConfiguration_ContainerDistributionConfigurationProperty `field:"optional" json:"containerDistributionConfiguration" yaml:"containerDistributionConfiguration"`
	// fast_launch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#fast_launch_configuration AwsDistributionConfiguration#fast_launch_configuration}
	// Experimental.
	FastLaunchConfiguration interface{} `field:"optional" json:"fastLaunchConfiguration" yaml:"fastLaunchConfiguration"`
	// launch_template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#launch_template_configuration AwsDistributionConfiguration#launch_template_configuration}
	// Experimental.
	LaunchTemplateConfiguration interface{} `field:"optional" json:"launchTemplateConfiguration" yaml:"launchTemplateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#license_configuration_arns AwsDistributionConfiguration#license_configuration_arns}.
	// Experimental.
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// s3_export_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#s3_export_configuration AwsDistributionConfiguration#s3_export_configuration}
	// Experimental.
	S3ExportConfiguration *AwsDistributionConfiguration_S3ExportConfigurationProperty `field:"optional" json:"s3ExportConfiguration" yaml:"s3ExportConfiguration"`
	// ssm_parameter_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#ssm_parameter_configuration AwsDistributionConfiguration#ssm_parameter_configuration}
	// Experimental.
	SsmParameterConfiguration interface{} `field:"optional" json:"ssmParameterConfiguration" yaml:"ssmParameterConfiguration"`
}

