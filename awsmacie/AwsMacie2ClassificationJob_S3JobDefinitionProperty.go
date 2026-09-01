package awsmacie


// Experimental.
type AwsMacie2ClassificationJob_S3JobDefinitionProperty struct {
	// bucket_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#bucket_criteria AwsMacie2ClassificationJob#bucket_criteria}
	// Experimental.
	BucketCriteria *AwsMacie2ClassificationJob_BucketCriteriaProperty `field:"optional" json:"bucketCriteria" yaml:"bucketCriteria"`
	// bucket_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#bucket_definitions AwsMacie2ClassificationJob#bucket_definitions}
	// Experimental.
	BucketDefinitions interface{} `field:"optional" json:"bucketDefinitions" yaml:"bucketDefinitions"`
	// scoping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#scoping AwsMacie2ClassificationJob#scoping}
	// Experimental.
	Scoping *AwsMacie2ClassificationJob_ScopingProperty `field:"optional" json:"scoping" yaml:"scoping"`
}

