package awsendusermessagingsms


// Experimental.
type TfEventDestination_CloudwatchLogsDestinationProperty struct {
	// ARN of the IAM role that End User Messaging SMS assumes to write to the log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_event_destination#iam_role_arn TfEventDestination#iam_role_arn}
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// ARN of the Amazon CloudWatch log group that receives the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_event_destination#log_group_arn TfEventDestination#log_group_arn}
	// Experimental.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

