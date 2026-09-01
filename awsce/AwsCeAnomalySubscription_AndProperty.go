package awsce


// Experimental.
type AwsCeAnomalySubscription_AndProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#cost_category AwsCeAnomalySubscription#cost_category}
	// Experimental.
	CostCategory *AwsCeAnomalySubscription_ThresholdExpressionAndCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#dimension AwsCeAnomalySubscription#dimension}
	// Experimental.
	Dimension *AwsCeAnomalySubscription_ThresholdExpressionAndDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#tags AwsCeAnomalySubscription#tags}
	// Experimental.
	Tags *AwsCeAnomalySubscription_ThresholdExpressionAndTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

