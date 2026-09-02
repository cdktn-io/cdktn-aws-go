package awsapplicationautoscaling


// Experimental.
type TfPolicy_PredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#id TfPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#expression TfPolicy#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#label TfPolicy#label}.
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric_stat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_stat TfPolicy#metric_stat}
	// Experimental.
	MetricStat *TfPolicy_PredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryMetricStatProperty `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#return_data TfPolicy#return_data}.
	// Experimental.
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

