package macie


// Experimental.
type AwsClassificationJob_S3JobDefinitionScopingExcludesAndProperty struct {
	// simple_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#simple_scope_term AwsClassificationJob#simple_scope_term}
	// Experimental.
	SimpleScopeTerm *AwsClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermProperty `field:"optional" json:"simpleScopeTerm" yaml:"simpleScopeTerm"`
	// tag_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_scope_term AwsClassificationJob#tag_scope_term}
	// Experimental.
	TagScopeTerm *AwsClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty `field:"optional" json:"tagScopeTerm" yaml:"tagScopeTerm"`
}

