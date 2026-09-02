package awsfis


// Experimental.
type TfExperimentTemplate_OutputsProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#s3_configuration TfExperimentTemplate#s3_configuration}
	// Experimental.
	S3Configuration *TfExperimentTemplate_ExperimentReportConfigurationOutputsS3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

