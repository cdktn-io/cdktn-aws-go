package awsmacie


// Experimental.
type TfClassificationJob_S3JobDefinitionScopingIncludesAndProperty struct {
	// simple_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#simple_scope_term TfClassificationJob#simple_scope_term}
	// Experimental.
	SimpleScopeTerm *TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermProperty `field:"optional" json:"simpleScopeTerm" yaml:"simpleScopeTerm"`
	// tag_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_scope_term TfClassificationJob#tag_scope_term}
	// Experimental.
	TagScopeTerm *TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermProperty `field:"optional" json:"tagScopeTerm" yaml:"tagScopeTerm"`
}

