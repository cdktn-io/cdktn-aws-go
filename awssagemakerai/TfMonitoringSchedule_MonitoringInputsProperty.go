package awssagemakerai


// Experimental.
type TfMonitoringSchedule_MonitoringInputsProperty struct {
	// batch_transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#batch_transform_input TfMonitoringSchedule#batch_transform_input}
	// Experimental.
	BatchTransformInput *TfMonitoringSchedule_BatchTransformInputProperty `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// endpoint_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#endpoint_input TfMonitoringSchedule#endpoint_input}
	// Experimental.
	EndpointInput *TfMonitoringSchedule_EndpointInputProperty `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

