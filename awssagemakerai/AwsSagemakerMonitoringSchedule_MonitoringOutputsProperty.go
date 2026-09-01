package awssagemakerai


// Experimental.
type AwsSagemakerMonitoringSchedule_MonitoringOutputsProperty struct {
	// s3_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#s3_output AwsSagemakerMonitoringSchedule#s3_output}
	// Experimental.
	S3Output *AwsSagemakerMonitoringSchedule_S3OutputProperty `field:"required" json:"s3Output" yaml:"s3Output"`
}

