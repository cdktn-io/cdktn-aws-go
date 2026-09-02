package awssagemakerai


// Experimental.
type TfMonitoringSchedule_MonitoringOutputConfigProperty struct {
	// monitoring_outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_outputs TfMonitoringSchedule#monitoring_outputs}
	// Experimental.
	MonitoringOutputs *TfMonitoringSchedule_MonitoringOutputsProperty `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#kms_key_id TfMonitoringSchedule#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

