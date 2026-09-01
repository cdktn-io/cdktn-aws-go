package awsce


// Experimental.
type AwsCeAnomalySubscription_SubscriberProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#address AwsCeAnomalySubscription#address}.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#type AwsCeAnomalySubscription#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

