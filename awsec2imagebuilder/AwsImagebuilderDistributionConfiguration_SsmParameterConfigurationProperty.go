package awsec2imagebuilder


// Experimental.
type AwsImagebuilderDistributionConfiguration_SsmParameterConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#parameter_name AwsImagebuilderDistributionConfiguration#parameter_name}.
	// Experimental.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#ami_account_id AwsImagebuilderDistributionConfiguration#ami_account_id}.
	// Experimental.
	AmiAccountId *string `field:"optional" json:"amiAccountId" yaml:"amiAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#data_type AwsImagebuilderDistributionConfiguration#data_type}.
	// Experimental.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
}

