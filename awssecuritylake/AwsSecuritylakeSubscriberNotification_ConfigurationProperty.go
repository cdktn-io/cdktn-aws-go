package awssecuritylake


// Experimental.
type AwsSecuritylakeSubscriberNotification_ConfigurationProperty struct {
	// https_notification_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber_notification#https_notification_configuration AwsSecuritylakeSubscriberNotification#https_notification_configuration}
	// Experimental.
	HttpsNotificationConfiguration interface{} `field:"optional" json:"httpsNotificationConfiguration" yaml:"httpsNotificationConfiguration"`
	// sqs_notification_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber_notification#sqs_notification_configuration AwsSecuritylakeSubscriberNotification#sqs_notification_configuration}
	// Experimental.
	SqsNotificationConfiguration interface{} `field:"optional" json:"sqsNotificationConfiguration" yaml:"sqsNotificationConfiguration"`
}

