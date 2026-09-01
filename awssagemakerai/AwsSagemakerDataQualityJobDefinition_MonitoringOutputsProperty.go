package awssagemakerai


// Experimental.
type AwsSagemakerDataQualityJobDefinition_MonitoringOutputsProperty struct {
	// s3_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#s3_output AwsSagemakerDataQualityJobDefinition#s3_output}
	// Experimental.
	S3Output *AwsSagemakerDataQualityJobDefinition_S3OutputProperty `field:"required" json:"s3Output" yaml:"s3Output"`
}

