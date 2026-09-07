package cloudwatchevidently


// Experimental.
type AwsLaunch_StepsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#group_weights AwsLaunch#group_weights}.
	// Experimental.
	GroupWeights *map[string]*float64 `field:"required" json:"groupWeights" yaml:"groupWeights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#start_time AwsLaunch#start_time}.
	// Experimental.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// segment_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#segment_overrides AwsLaunch#segment_overrides}
	// Experimental.
	SegmentOverrides interface{} `field:"optional" json:"segmentOverrides" yaml:"segmentOverrides"`
}

