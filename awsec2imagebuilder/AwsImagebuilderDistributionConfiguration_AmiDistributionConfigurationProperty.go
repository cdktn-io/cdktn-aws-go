package awsec2imagebuilder


// Experimental.
type AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#ami_tags AwsImagebuilderDistributionConfiguration#ami_tags}.
	// Experimental.
	AmiTags *map[string]*string `field:"optional" json:"amiTags" yaml:"amiTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#description AwsImagebuilderDistributionConfiguration#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#kms_key_id AwsImagebuilderDistributionConfiguration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// launch_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#launch_permission AwsImagebuilderDistributionConfiguration#launch_permission}
	// Experimental.
	LaunchPermission *AwsImagebuilderDistributionConfiguration_LaunchPermissionProperty `field:"optional" json:"launchPermission" yaml:"launchPermission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#name AwsImagebuilderDistributionConfiguration#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#target_account_ids AwsImagebuilderDistributionConfiguration#target_account_ids}.
	// Experimental.
	TargetAccountIds *[]*string `field:"optional" json:"targetAccountIds" yaml:"targetAccountIds"`
}

