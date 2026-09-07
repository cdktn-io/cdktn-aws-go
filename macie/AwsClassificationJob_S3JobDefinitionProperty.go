package macie


// Experimental.
type AwsClassificationJob_S3JobDefinitionProperty struct {
	// bucket_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#bucket_criteria AwsClassificationJob#bucket_criteria}
	// Experimental.
	BucketCriteria *AwsClassificationJob_BucketCriteriaProperty `field:"optional" json:"bucketCriteria" yaml:"bucketCriteria"`
	// bucket_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#bucket_definitions AwsClassificationJob#bucket_definitions}
	// Experimental.
	BucketDefinitions interface{} `field:"optional" json:"bucketDefinitions" yaml:"bucketDefinitions"`
	// scoping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#scoping AwsClassificationJob#scoping}
	// Experimental.
	Scoping *AwsClassificationJob_ScopingProperty `field:"optional" json:"scoping" yaml:"scoping"`
}

