package awssns

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTopicSubscriptionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#endpoint TfTopicSubscription#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#protocol TfTopicSubscription#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#topic_arn TfTopicSubscription#topic_arn}.
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#confirmation_timeout_in_minutes TfTopicSubscription#confirmation_timeout_in_minutes}.
	// Experimental.
	ConfirmationTimeoutInMinutes *float64 `field:"optional" json:"confirmationTimeoutInMinutes" yaml:"confirmationTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#delivery_policy TfTopicSubscription#delivery_policy}.
	// Experimental.
	DeliveryPolicy *string `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#endpoint_auto_confirms TfTopicSubscription#endpoint_auto_confirms}.
	// Experimental.
	EndpointAutoConfirms interface{} `field:"optional" json:"endpointAutoConfirms" yaml:"endpointAutoConfirms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#filter_policy TfTopicSubscription#filter_policy}.
	// Experimental.
	FilterPolicy *string `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#filter_policy_scope TfTopicSubscription#filter_policy_scope}.
	// Experimental.
	FilterPolicyScope *string `field:"optional" json:"filterPolicyScope" yaml:"filterPolicyScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#id TfTopicSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#raw_message_delivery TfTopicSubscription#raw_message_delivery}.
	// Experimental.
	RawMessageDelivery interface{} `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#redrive_policy TfTopicSubscription#redrive_policy}.
	// Experimental.
	RedrivePolicy *string `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#region TfTopicSubscription#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#replay_policy TfTopicSubscription#replay_policy}.
	// Experimental.
	ReplayPolicy *string `field:"optional" json:"replayPolicy" yaml:"replayPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic_subscription#subscription_role_arn TfTopicSubscription#subscription_role_arn}.
	// Experimental.
	SubscriptionRoleArn *string `field:"optional" json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
}

