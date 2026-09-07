package ce


// Experimental.
type AwsAnomalySubscription_ThresholdExpressionProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#and AwsAnomalySubscription#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#cost_category AwsAnomalySubscription#cost_category}
	// Experimental.
	CostCategory *AwsAnomalySubscription_ThresholdExpressionCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#dimension AwsAnomalySubscription#dimension}
	// Experimental.
	Dimension *AwsAnomalySubscription_ThresholdExpressionDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#not AwsAnomalySubscription#not}
	// Experimental.
	Not *AwsAnomalySubscription_NotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#or AwsAnomalySubscription#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#tags AwsAnomalySubscription#tags}
	// Experimental.
	Tags *AwsAnomalySubscription_ThresholdExpressionTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

