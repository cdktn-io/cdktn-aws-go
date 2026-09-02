package awsses


// Experimental.
type TfReceiptRule_LambdaActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#function_arn TfReceiptRule#function_arn}.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#position TfReceiptRule#position}.
	// Experimental.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#invocation_type TfReceiptRule#invocation_type}.
	// Experimental.
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#topic_arn TfReceiptRule#topic_arn}.
	// Experimental.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

