package awssagemakerai


// Experimental.
type AwsSagemakerDataQualityJobDefinition_DataQualityJobOutputConfigProperty struct {
	// monitoring_outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#monitoring_outputs AwsSagemakerDataQualityJobDefinition#monitoring_outputs}
	// Experimental.
	MonitoringOutputs *AwsSagemakerDataQualityJobDefinition_MonitoringOutputsProperty `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#kms_key_id AwsSagemakerDataQualityJobDefinition#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

