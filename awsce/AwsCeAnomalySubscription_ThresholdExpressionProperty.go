package awsce


// Experimental.
type AwsCeAnomalySubscription_ThresholdExpressionProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#and AwsCeAnomalySubscription#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#cost_category AwsCeAnomalySubscription#cost_category}
	// Experimental.
	CostCategory *AwsCeAnomalySubscription_ThresholdExpressionCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#dimension AwsCeAnomalySubscription#dimension}
	// Experimental.
	Dimension *AwsCeAnomalySubscription_ThresholdExpressionDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#not AwsCeAnomalySubscription#not}
	// Experimental.
	Not *AwsCeAnomalySubscription_NotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#or AwsCeAnomalySubscription#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#tags AwsCeAnomalySubscription#tags}
	// Experimental.
	Tags *AwsCeAnomalySubscription_ThresholdExpressionTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

