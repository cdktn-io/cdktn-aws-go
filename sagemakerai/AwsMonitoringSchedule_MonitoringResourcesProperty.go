package sagemakerai


// Experimental.
type AwsMonitoringSchedule_MonitoringResourcesProperty struct {
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#cluster_config AwsMonitoringSchedule#cluster_config}
	// Experimental.
	ClusterConfig *AwsMonitoringSchedule_ClusterConfigProperty `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

