package awscloudwatchevidently


// Experimental.
type AwsEvidentlyLaunch_StepsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#group_weights AwsEvidentlyLaunch#group_weights}.
	// Experimental.
	GroupWeights *map[string]*float64 `field:"required" json:"groupWeights" yaml:"groupWeights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#start_time AwsEvidentlyLaunch#start_time}.
	// Experimental.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// segment_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#segment_overrides AwsEvidentlyLaunch#segment_overrides}
	// Experimental.
	SegmentOverrides interface{} `field:"optional" json:"segmentOverrides" yaml:"segmentOverrides"`
}

