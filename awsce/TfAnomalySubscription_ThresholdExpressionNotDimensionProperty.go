package awsce


// Experimental.
type TfAnomalySubscription_ThresholdExpressionNotDimensionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#key TfAnomalySubscription#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#match_options TfAnomalySubscription#match_options}.
	// Experimental.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#values TfAnomalySubscription#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

