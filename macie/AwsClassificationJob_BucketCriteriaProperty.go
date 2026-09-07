package macie


// Experimental.
type AwsClassificationJob_BucketCriteriaProperty struct {
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#excludes AwsClassificationJob#excludes}
	// Experimental.
	Excludes *AwsClassificationJob_S3JobDefinitionBucketCriteriaExcludesProperty `field:"optional" json:"excludes" yaml:"excludes"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#includes AwsClassificationJob#includes}
	// Experimental.
	Includes *AwsClassificationJob_S3JobDefinitionBucketCriteriaIncludesProperty `field:"optional" json:"includes" yaml:"includes"`
}

