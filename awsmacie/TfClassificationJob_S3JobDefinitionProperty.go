package awsmacie


// Experimental.
type TfClassificationJob_S3JobDefinitionProperty struct {
	// bucket_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#bucket_criteria TfClassificationJob#bucket_criteria}
	// Experimental.
	BucketCriteria *TfClassificationJob_BucketCriteriaProperty `field:"optional" json:"bucketCriteria" yaml:"bucketCriteria"`
	// bucket_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#bucket_definitions TfClassificationJob#bucket_definitions}
	// Experimental.
	BucketDefinitions interface{} `field:"optional" json:"bucketDefinitions" yaml:"bucketDefinitions"`
	// scoping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#scoping TfClassificationJob#scoping}
	// Experimental.
	Scoping *TfClassificationJob_ScopingProperty `field:"optional" json:"scoping" yaml:"scoping"`
}

