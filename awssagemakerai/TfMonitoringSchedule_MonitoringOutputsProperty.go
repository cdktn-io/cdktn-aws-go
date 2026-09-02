package awssagemakerai


// Experimental.
type TfMonitoringSchedule_MonitoringOutputsProperty struct {
	// s3_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#s3_output TfMonitoringSchedule#s3_output}
	// Experimental.
	S3Output *TfMonitoringSchedule_S3OutputProperty `field:"required" json:"s3Output" yaml:"s3Output"`
}

