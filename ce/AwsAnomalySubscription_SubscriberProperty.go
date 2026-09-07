package ce


// Experimental.
type AwsAnomalySubscription_SubscriberProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#address AwsAnomalySubscription#address}.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#type AwsAnomalySubscription#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

