package awskendra


// Experimental.
type AwsKendraDataSource_InlineConfigurationsProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition AwsKendraDataSource#condition}
	// Experimental.
	Condition *AwsKendraDataSource_ConditionProperty `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#document_content_deletion AwsKendraDataSource#document_content_deletion}.
	// Experimental.
	DocumentContentDeletion interface{} `field:"optional" json:"documentContentDeletion" yaml:"documentContentDeletion"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#target AwsKendraDataSource#target}
	// Experimental.
	Target *AwsKendraDataSource_TargetProperty `field:"optional" json:"target" yaml:"target"`
}

