package ce


// Experimental.
type AwsAnomalySubscription_AndProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#cost_category AwsAnomalySubscription#cost_category}
	// Experimental.
	CostCategory *AwsAnomalySubscription_ThresholdExpressionAndCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#dimension AwsAnomalySubscription#dimension}
	// Experimental.
	Dimension *AwsAnomalySubscription_ThresholdExpressionAndDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#tags AwsAnomalySubscription#tags}
	// Experimental.
	Tags *AwsAnomalySubscription_ThresholdExpressionAndTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

