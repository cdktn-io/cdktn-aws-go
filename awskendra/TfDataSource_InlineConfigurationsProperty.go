package awskendra


// Experimental.
type TfDataSource_InlineConfigurationsProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition TfDataSource#condition}
	// Experimental.
	Condition *TfDataSource_ConditionProperty `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#document_content_deletion TfDataSource#document_content_deletion}.
	// Experimental.
	DocumentContentDeletion interface{} `field:"optional" json:"documentContentDeletion" yaml:"documentContentDeletion"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#target TfDataSource#target}
	// Experimental.
	Target *TfDataSource_TargetProperty `field:"optional" json:"target" yaml:"target"`
}

