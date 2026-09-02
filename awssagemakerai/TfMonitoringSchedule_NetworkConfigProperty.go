package awssagemakerai


// Experimental.
type TfMonitoringSchedule_NetworkConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#enable_inter_container_traffic_encryption TfMonitoringSchedule#enable_inter_container_traffic_encryption}.
	// Experimental.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#enable_network_isolation TfMonitoringSchedule#enable_network_isolation}.
	// Experimental.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#vpc_config TfMonitoringSchedule#vpc_config}
	// Experimental.
	VpcConfig *TfMonitoringSchedule_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

