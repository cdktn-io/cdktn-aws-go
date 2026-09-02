package awsce


// Experimental.
type TfAnomalySubscription_ThresholdExpressionProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#and TfAnomalySubscription#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#cost_category TfAnomalySubscription#cost_category}
	// Experimental.
	CostCategory *TfAnomalySubscription_ThresholdExpressionCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#dimension TfAnomalySubscription#dimension}
	// Experimental.
	Dimension *TfAnomalySubscription_ThresholdExpressionDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#not TfAnomalySubscription#not}
	// Experimental.
	Not *TfAnomalySubscription_NotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#or TfAnomalySubscription#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_anomaly_subscription#tags TfAnomalySubscription#tags}
	// Experimental.
	Tags *TfAnomalySubscription_ThresholdExpressionTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

