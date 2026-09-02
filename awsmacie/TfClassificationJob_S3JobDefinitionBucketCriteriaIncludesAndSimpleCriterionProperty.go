package awsmacie


// Experimental.
type TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#comparator TfClassificationJob#comparator}.
	// Experimental.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#key TfClassificationJob#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#values TfClassificationJob#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

