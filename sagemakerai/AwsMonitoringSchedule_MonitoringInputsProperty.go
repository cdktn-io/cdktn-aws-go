package sagemakerai


// Experimental.
type AwsMonitoringSchedule_MonitoringInputsProperty struct {
	// batch_transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#batch_transform_input AwsMonitoringSchedule#batch_transform_input}
	// Experimental.
	BatchTransformInput *AwsMonitoringSchedule_BatchTransformInputProperty `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// endpoint_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#endpoint_input AwsMonitoringSchedule#endpoint_input}
	// Experimental.
	EndpointInput *AwsMonitoringSchedule_EndpointInputProperty `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

