package awsmacie


// Experimental.
type TfClassificationJob_ScopingProperty struct {
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#excludes TfClassificationJob#excludes}
	// Experimental.
	Excludes *TfClassificationJob_S3JobDefinitionScopingExcludesProperty `field:"optional" json:"excludes" yaml:"excludes"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#includes TfClassificationJob#includes}
	// Experimental.
	Includes *TfClassificationJob_S3JobDefinitionScopingIncludesProperty `field:"optional" json:"includes" yaml:"includes"`
}

