package cloudfront


// Experimental.
type AwsMonitoringSubscription_MonitoringSubscriptionProperty struct {
	// realtime_metrics_subscription_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_monitoring_subscription#realtime_metrics_subscription_config AwsMonitoringSubscription#realtime_metrics_subscription_config}
	// Experimental.
	RealtimeMetricsSubscriptionConfig *AwsMonitoringSubscription_RealtimeMetricsSubscriptionConfigProperty `field:"required" json:"realtimeMetricsSubscriptionConfig" yaml:"realtimeMetricsSubscriptionConfig"`
}

