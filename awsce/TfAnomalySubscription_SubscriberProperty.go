package awsce


// Experimental.
type TfAnomalySubscription_SubscriberProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#address TfAnomalySubscription#address}.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#type TfAnomalySubscription#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

