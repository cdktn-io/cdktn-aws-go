package rds


// Experimental.
type AwsDbOptionGroup_OptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#option_name AwsDbOptionGroup#option_name}.
	// Experimental.
	OptionName *string `field:"required" json:"optionName" yaml:"optionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#db_security_group_memberships AwsDbOptionGroup#db_security_group_memberships}.
	// Experimental.
	DbSecurityGroupMemberships *[]*string `field:"optional" json:"dbSecurityGroupMemberships" yaml:"dbSecurityGroupMemberships"`
	// option_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#option_settings AwsDbOptionGroup#option_settings}
	// Experimental.
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#port AwsDbOptionGroup#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#version AwsDbOptionGroup#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#vpc_security_group_memberships AwsDbOptionGroup#vpc_security_group_memberships}.
	// Experimental.
	VpcSecurityGroupMemberships *[]*string `field:"optional" json:"vpcSecurityGroupMemberships" yaml:"vpcSecurityGroupMemberships"`
}

