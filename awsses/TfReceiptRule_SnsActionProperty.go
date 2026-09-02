package awsses


// Experimental.
type TfReceiptRule_SnsActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#position TfReceiptRule#position}.
	// Experimental.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#topic_arn TfReceiptRule#topic_arn}.
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#encoding TfReceiptRule#encoding}.
	// Experimental.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
}

