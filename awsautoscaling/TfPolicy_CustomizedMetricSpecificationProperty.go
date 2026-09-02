package awsautoscaling


// Experimental.
type TfPolicy_CustomizedMetricSpecificationProperty struct {
	// metric_dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric_dimension TfPolicy#metric_dimension}
	// Experimental.
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric_name TfPolicy#metric_name}.
	// Experimental.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metrics TfPolicy#metrics}
	// Experimental.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#namespace TfPolicy#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#period TfPolicy#period}.
	// Experimental.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#statistic TfPolicy#statistic}.
	// Experimental.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#unit TfPolicy#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

