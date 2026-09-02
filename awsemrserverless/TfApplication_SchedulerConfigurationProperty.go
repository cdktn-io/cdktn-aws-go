package awsemrserverless


// Experimental.
type TfApplication_SchedulerConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#max_concurrent_runs TfApplication#max_concurrent_runs}.
	// Experimental.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#queue_timeout_minutes TfApplication#queue_timeout_minutes}.
	// Experimental.
	QueueTimeoutMinutes *float64 `field:"optional" json:"queueTimeoutMinutes" yaml:"queueTimeoutMinutes"`
}

