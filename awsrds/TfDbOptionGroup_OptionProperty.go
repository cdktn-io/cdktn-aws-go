package awsrds


// Experimental.
type TfDbOptionGroup_OptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#option_name TfDbOptionGroup#option_name}.
	// Experimental.
	OptionName *string `field:"required" json:"optionName" yaml:"optionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#db_security_group_memberships TfDbOptionGroup#db_security_group_memberships}.
	// Experimental.
	DbSecurityGroupMemberships *[]*string `field:"optional" json:"dbSecurityGroupMemberships" yaml:"dbSecurityGroupMemberships"`
	// option_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#option_settings TfDbOptionGroup#option_settings}
	// Experimental.
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#port TfDbOptionGroup#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#version TfDbOptionGroup#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#vpc_security_group_memberships TfDbOptionGroup#vpc_security_group_memberships}.
	// Experimental.
	VpcSecurityGroupMemberships *[]*string `field:"optional" json:"vpcSecurityGroupMemberships" yaml:"vpcSecurityGroupMemberships"`
}

