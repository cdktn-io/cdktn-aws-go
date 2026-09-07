package webservicesbudgets


// Experimental.
type AwsBudget_NotificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#comparison_operator AwsBudget#comparison_operator}.
	// Experimental.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#notification_type AwsBudget#notification_type}.
	// Experimental.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#threshold AwsBudget#threshold}.
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#threshold_type AwsBudget#threshold_type}.
	// Experimental.
	ThresholdType *string `field:"required" json:"thresholdType" yaml:"thresholdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#subscriber_email_addresses AwsBudget#subscriber_email_addresses}.
	// Experimental.
	SubscriberEmailAddresses *[]*string `field:"optional" json:"subscriberEmailAddresses" yaml:"subscriberEmailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#subscriber_sns_topic_arns AwsBudget#subscriber_sns_topic_arns}.
	// Experimental.
	SubscriberSnsTopicArns *[]*string `field:"optional" json:"subscriberSnsTopicArns" yaml:"subscriberSnsTopicArns"`
}

