package awsses


// Experimental.
type TfReceiptRule_S3ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#bucket_name TfReceiptRule#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#position TfReceiptRule#position}.
	// Experimental.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#iam_role_arn TfReceiptRule#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#kms_key_arn TfReceiptRule#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#object_key_prefix TfReceiptRule#object_key_prefix}.
	// Experimental.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#topic_arn TfReceiptRule#topic_arn}.
	// Experimental.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

