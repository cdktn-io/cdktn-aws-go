package awssagemakerai


// Experimental.
type AwsSagemakerMonitoringSchedule_NetworkConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#enable_inter_container_traffic_encryption AwsSagemakerMonitoringSchedule#enable_inter_container_traffic_encryption}.
	// Experimental.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#enable_network_isolation AwsSagemakerMonitoringSchedule#enable_network_isolation}.
	// Experimental.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#vpc_config AwsSagemakerMonitoringSchedule#vpc_config}
	// Experimental.
	VpcConfig *AwsSagemakerMonitoringSchedule_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

