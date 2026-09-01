package awsec2imagebuilder


// Experimental.
type AwsImagebuilderDistributionConfiguration_LaunchTemplateConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#launch_template_id AwsImagebuilderDistributionConfiguration#launch_template_id}.
	// Experimental.
	LaunchTemplateId *string `field:"required" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#account_id AwsImagebuilderDistributionConfiguration#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#default AwsImagebuilderDistributionConfiguration#default}.
	// Experimental.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
}

