package awsfis


// Experimental.
type AwsFisExperimentTemplate_OutputsProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#s3_configuration AwsFisExperimentTemplate#s3_configuration}
	// Experimental.
	S3Configuration *AwsFisExperimentTemplate_ExperimentReportConfigurationOutputsS3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

