package awsmacie


// Experimental.
type TfClassificationJob_S3JobDefinitionBucketCriteriaExcludesAndTagCriterionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#comparator TfClassificationJob#comparator}.
	// Experimental.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// tag_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#tag_values TfClassificationJob#tag_values}
	// Experimental.
	TagValues interface{} `field:"optional" json:"tagValues" yaml:"tagValues"`
}

