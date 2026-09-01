package awssagemakerai


// Experimental.
type AwsSagemakerMonitoringSchedule_MonitoringInputsProperty struct {
	// batch_transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#batch_transform_input AwsSagemakerMonitoringSchedule#batch_transform_input}
	// Experimental.
	BatchTransformInput *AwsSagemakerMonitoringSchedule_BatchTransformInputProperty `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// endpoint_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#endpoint_input AwsSagemakerMonitoringSchedule#endpoint_input}
	// Experimental.
	EndpointInput *AwsSagemakerMonitoringSchedule_EndpointInputProperty `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

