package awselbclassic


// Experimental.
type TfElb_HealthCheckProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#healthy_threshold TfElb#healthy_threshold}.
	// Experimental.
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#interval TfElb#interval}.
	// Experimental.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#target TfElb#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#timeout TfElb#timeout}.
	// Experimental.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#unhealthy_threshold TfElb#unhealthy_threshold}.
	// Experimental.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

