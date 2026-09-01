package awssagemakerai


// Experimental.
type AwsSagemakerDataQualityJobDefinition_NetworkConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#enable_inter_container_traffic_encryption AwsSagemakerDataQualityJobDefinition#enable_inter_container_traffic_encryption}.
	// Experimental.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#enable_network_isolation AwsSagemakerDataQualityJobDefinition#enable_network_isolation}.
	// Experimental.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#vpc_config AwsSagemakerDataQualityJobDefinition#vpc_config}
	// Experimental.
	VpcConfig *AwsSagemakerDataQualityJobDefinition_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

