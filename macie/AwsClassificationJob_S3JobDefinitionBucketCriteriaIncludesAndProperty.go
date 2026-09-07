package macie


// Experimental.
type AwsClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndProperty struct {
	// simple_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#simple_criterion AwsClassificationJob#simple_criterion}
	// Experimental.
	SimpleCriterion *AwsClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty `field:"optional" json:"simpleCriterion" yaml:"simpleCriterion"`
	// tag_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_criterion AwsClassificationJob#tag_criterion}
	// Experimental.
	TagCriterion *AwsClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionProperty `field:"optional" json:"tagCriterion" yaml:"tagCriterion"`
}

