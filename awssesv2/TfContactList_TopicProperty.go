package awssesv2


// Experimental.
type TfContactList_TopicProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#default_subscription_status TfContactList#default_subscription_status}.
	// Experimental.
	DefaultSubscriptionStatus *string `field:"required" json:"defaultSubscriptionStatus" yaml:"defaultSubscriptionStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#display_name TfContactList#display_name}.
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#topic_name TfContactList#topic_name}.
	// Experimental.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_contact_list#description TfContactList#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

