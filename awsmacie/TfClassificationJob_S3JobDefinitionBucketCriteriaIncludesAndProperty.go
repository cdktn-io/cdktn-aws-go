package awsmacie


// Experimental.
type TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndProperty struct {
	// simple_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#simple_criterion TfClassificationJob#simple_criterion}
	// Experimental.
	SimpleCriterion *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty `field:"optional" json:"simpleCriterion" yaml:"simpleCriterion"`
	// tag_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_criterion TfClassificationJob#tag_criterion}
	// Experimental.
	TagCriterion *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionProperty `field:"optional" json:"tagCriterion" yaml:"tagCriterion"`
}

