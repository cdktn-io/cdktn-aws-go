package awsfis


// Experimental.
type TfExperimentTemplate_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#log_schema_version TfExperimentTemplate#log_schema_version}.
	// Experimental.
	LogSchemaVersion *float64 `field:"required" json:"logSchemaVersion" yaml:"logSchemaVersion"`
	// cloudwatch_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#cloudwatch_logs_configuration TfExperimentTemplate#cloudwatch_logs_configuration}
	// Experimental.
	CloudwatchLogsConfiguration *TfExperimentTemplate_CloudwatchLogsConfigurationProperty `field:"optional" json:"cloudwatchLogsConfiguration" yaml:"cloudwatchLogsConfiguration"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#s3_configuration TfExperimentTemplate#s3_configuration}
	// Experimental.
	S3Configuration *TfExperimentTemplate_LogConfigurationS3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

