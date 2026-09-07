package ce


// Experimental.
type AwsAnomalySubscription_NotProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#cost_category AwsAnomalySubscription#cost_category}
	// Experimental.
	CostCategory *AwsAnomalySubscription_ThresholdExpressionNotCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#dimension AwsAnomalySubscription#dimension}
	// Experimental.
	Dimension *AwsAnomalySubscription_ThresholdExpressionNotDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#tags AwsAnomalySubscription#tags}
	// Experimental.
	Tags *AwsAnomalySubscription_ThresholdExpressionNotTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

