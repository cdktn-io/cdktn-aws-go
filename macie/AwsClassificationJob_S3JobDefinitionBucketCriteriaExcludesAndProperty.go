package macie


// Experimental.
type AwsClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndProperty struct {
	// simple_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#simple_criterion AwsClassificationJob#simple_criterion}
	// Experimental.
	SimpleCriterion *AwsClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndSimpleCriterionProperty `field:"optional" json:"simpleCriterion" yaml:"simpleCriterion"`
	// tag_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_criterion AwsClassificationJob#tag_criterion}
	// Experimental.
	TagCriterion *AwsClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndTagCriterionProperty `field:"optional" json:"tagCriterion" yaml:"tagCriterion"`
}

