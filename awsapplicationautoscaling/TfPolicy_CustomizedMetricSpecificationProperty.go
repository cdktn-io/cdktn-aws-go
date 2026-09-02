package awsapplicationautoscaling


// Experimental.
type TfPolicy_CustomizedMetricSpecificationProperty struct {
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#dimensions TfPolicy#dimensions}
	// Experimental.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_name TfPolicy#metric_name}.
	// Experimental.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metrics TfPolicy#metrics}
	// Experimental.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#namespace TfPolicy#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#statistic TfPolicy#statistic}.
	// Experimental.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#unit TfPolicy#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

