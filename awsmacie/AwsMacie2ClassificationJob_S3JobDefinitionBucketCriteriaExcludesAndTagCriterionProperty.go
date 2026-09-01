package awsmacie


// Experimental.
type AwsMacie2ClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndTagCriterionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#comparator AwsMacie2ClassificationJob#comparator}.
	// Experimental.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_values AwsMacie2ClassificationJob#tag_values}
	// Experimental.
	TagValues interface{} `field:"optional" json:"tagValues" yaml:"tagValues"`
}

