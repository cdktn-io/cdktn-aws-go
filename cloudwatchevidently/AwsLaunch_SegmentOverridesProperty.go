package cloudwatchevidently


// Experimental.
type AwsLaunch_SegmentOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#evaluation_order AwsLaunch#evaluation_order}.
	// Experimental.
	EvaluationOrder *float64 `field:"required" json:"evaluationOrder" yaml:"evaluationOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#segment AwsLaunch#segment}.
	// Experimental.
	Segment *string `field:"required" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#weights AwsLaunch#weights}.
	// Experimental.
	Weights *map[string]*float64 `field:"required" json:"weights" yaml:"weights"`
}

