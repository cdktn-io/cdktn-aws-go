package awssagemakerai


// Experimental.
type TfMonitoringSchedule_MonitoringResourcesProperty struct {
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#cluster_config TfMonitoringSchedule#cluster_config}
	// Experimental.
	ClusterConfig *TfMonitoringSchedule_ClusterConfigProperty `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

