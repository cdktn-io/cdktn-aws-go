package awssesv2


// Experimental.
type AwsSesv2ContactList_TopicProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#default_subscription_status AwsSesv2ContactList#default_subscription_status}.
	// Experimental.
	DefaultSubscriptionStatus *string `field:"required" json:"defaultSubscriptionStatus" yaml:"defaultSubscriptionStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#display_name AwsSesv2ContactList#display_name}.
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#topic_name AwsSesv2ContactList#topic_name}.
	// Experimental.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#description AwsSesv2ContactList#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

