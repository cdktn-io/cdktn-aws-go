package kendra


// Experimental.
type AwsDataSource_InlineConfigurationsProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition AwsDataSource#condition}
	// Experimental.
	Condition *AwsDataSource_ConditionProperty `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#document_content_deletion AwsDataSource#document_content_deletion}.
	// Experimental.
	DocumentContentDeletion interface{} `field:"optional" json:"documentContentDeletion" yaml:"documentContentDeletion"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#target AwsDataSource#target}
	// Experimental.
	Target *AwsDataSource_TargetProperty `field:"optional" json:"target" yaml:"target"`
}

