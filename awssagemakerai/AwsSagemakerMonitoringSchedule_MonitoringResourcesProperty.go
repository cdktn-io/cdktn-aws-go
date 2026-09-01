package awssagemakerai


// Experimental.
type AwsSagemakerMonitoringSchedule_MonitoringResourcesProperty struct {
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#cluster_config AwsSagemakerMonitoringSchedule#cluster_config}
	// Experimental.
	ClusterConfig *AwsSagemakerMonitoringSchedule_ClusterConfigProperty `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

