package finspace


// Experimental.
type AwsKxCluster_AutoScalingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#auto_scaling_metric AwsKxCluster#auto_scaling_metric}.
	// Experimental.
	AutoScalingMetric *string `field:"required" json:"autoScalingMetric" yaml:"autoScalingMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#max_node_count AwsKxCluster#max_node_count}.
	// Experimental.
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#metric_target AwsKxCluster#metric_target}.
	// Experimental.
	MetricTarget *float64 `field:"required" json:"metricTarget" yaml:"metricTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#min_node_count AwsKxCluster#min_node_count}.
	// Experimental.
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#scale_in_cooldown_seconds AwsKxCluster#scale_in_cooldown_seconds}.
	// Experimental.
	ScaleInCooldownSeconds *float64 `field:"required" json:"scaleInCooldownSeconds" yaml:"scaleInCooldownSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#scale_out_cooldown_seconds AwsKxCluster#scale_out_cooldown_seconds}.
	// Experimental.
	ScaleOutCooldownSeconds *float64 `field:"required" json:"scaleOutCooldownSeconds" yaml:"scaleOutCooldownSeconds"`
}

