package rds


// Experimental.
type AwsDbOptionGroup_OptionSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#name AwsDbOptionGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_option_group#value AwsDbOptionGroup#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

