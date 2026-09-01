package awsmacie


// Experimental.
type AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndProperty struct {
	// simple_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#simple_scope_term AwsMacie2ClassificationJob#simple_scope_term}
	// Experimental.
	SimpleScopeTerm *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndSimpleScopeTermProperty `field:"optional" json:"simpleScopeTerm" yaml:"simpleScopeTerm"`
	// tag_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_scope_term AwsMacie2ClassificationJob#tag_scope_term}
	// Experimental.
	TagScopeTerm *AwsMacie2ClassificationJob_S3JobDefinitionScopingExcludesAndTagScopeTermProperty `field:"optional" json:"tagScopeTerm" yaml:"tagScopeTerm"`
}

