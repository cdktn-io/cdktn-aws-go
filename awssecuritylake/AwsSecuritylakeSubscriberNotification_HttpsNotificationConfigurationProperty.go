package awssecuritylake


// Experimental.
type AwsSecuritylakeSubscriberNotification_HttpsNotificationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber_notification#endpoint AwsSecuritylakeSubscriberNotification#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber_notification#target_role_arn AwsSecuritylakeSubscriberNotification#target_role_arn}.
	// Experimental.
	TargetRoleArn *string `field:"required" json:"targetRoleArn" yaml:"targetRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber_notification#authorization_api_key_name AwsSecuritylakeSubscriberNotification#authorization_api_key_name}.
	// Experimental.
	AuthorizationApiKeyName *string `field:"optional" json:"authorizationApiKeyName" yaml:"authorizationApiKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber_notification#authorization_api_key_value AwsSecuritylakeSubscriberNotification#authorization_api_key_value}.
	// Experimental.
	AuthorizationApiKeyValue *string `field:"optional" json:"authorizationApiKeyValue" yaml:"authorizationApiKeyValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber_notification#http_method AwsSecuritylakeSubscriberNotification#http_method}.
	// Experimental.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
}

