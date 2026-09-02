package awsce


// Experimental.
type TfAnomalySubscription_OrProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#cost_category TfAnomalySubscription#cost_category}
	// Experimental.
	CostCategory *TfAnomalySubscription_ThresholdExpressionOrCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#dimension TfAnomalySubscription#dimension}
	// Experimental.
	Dimension *TfAnomalySubscription_ThresholdExpressionOrDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#tags TfAnomalySubscription#tags}
	// Experimental.
	Tags *TfAnomalySubscription_ThresholdExpressionOrTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

