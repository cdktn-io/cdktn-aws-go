package awscloudfront


// Experimental.
type AwsCloudfrontMonitoringSubscription_MonitoringSubscriptionProperty struct {
	// realtime_metrics_subscription_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_monitoring_subscription#realtime_metrics_subscription_config AwsCloudfrontMonitoringSubscription#realtime_metrics_subscription_config}
	// Experimental.
	RealtimeMetricsSubscriptionConfig *AwsCloudfrontMonitoringSubscription_RealtimeMetricsSubscriptionConfigProperty `field:"required" json:"realtimeMetricsSubscriptionConfig" yaml:"realtimeMetricsSubscriptionConfig"`
}

