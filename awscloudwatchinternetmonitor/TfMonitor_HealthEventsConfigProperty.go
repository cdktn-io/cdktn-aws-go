package awscloudwatchinternetmonitor


// Experimental.
type TfMonitor_HealthEventsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internetmonitor_monitor#availability_score_threshold TfMonitor#availability_score_threshold}.
	// Experimental.
	AvailabilityScoreThreshold *float64 `field:"optional" json:"availabilityScoreThreshold" yaml:"availabilityScoreThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internetmonitor_monitor#performance_score_threshold TfMonitor#performance_score_threshold}.
	// Experimental.
	PerformanceScoreThreshold *float64 `field:"optional" json:"performanceScoreThreshold" yaml:"performanceScoreThreshold"`
}

