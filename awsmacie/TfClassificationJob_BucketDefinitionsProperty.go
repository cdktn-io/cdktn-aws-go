package awsmacie


// Experimental.
type TfClassificationJob_BucketDefinitionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#account_id TfClassificationJob#account_id}.
	// Experimental.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#buckets TfClassificationJob#buckets}.
	// Experimental.
	Buckets *[]*string `field:"required" json:"buckets" yaml:"buckets"`
}

