package awsmacie


// Experimental.
type AwsMacie2ClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndProperty struct {
	// simple_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#simple_criterion AwsMacie2ClassificationJob#simple_criterion}
	// Experimental.
	SimpleCriterion *AwsMacie2ClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndSimpleCriterionProperty `field:"optional" json:"simpleCriterion" yaml:"simpleCriterion"`
	// tag_criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_criterion AwsMacie2ClassificationJob#tag_criterion}
	// Experimental.
	TagCriterion *AwsMacie2ClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndTagCriterionProperty `field:"optional" json:"tagCriterion" yaml:"tagCriterion"`
}

