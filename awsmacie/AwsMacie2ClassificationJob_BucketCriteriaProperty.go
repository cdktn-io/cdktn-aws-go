package awsmacie


// Experimental.
type AwsMacie2ClassificationJob_BucketCriteriaProperty struct {
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#excludes AwsMacie2ClassificationJob#excludes}
	// Experimental.
	Excludes *AwsMacie2ClassificationJob_S3JobDefinitionBucketCriteriaExcludesProperty `field:"optional" json:"excludes" yaml:"excludes"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#includes AwsMacie2ClassificationJob#includes}
	// Experimental.
	Includes *AwsMacie2ClassificationJob_S3JobDefinitionBucketCriteriaIncludesProperty `field:"optional" json:"includes" yaml:"includes"`
}

