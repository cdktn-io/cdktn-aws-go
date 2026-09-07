package fis


// Experimental.
type AwsExperimentTemplate_ExperimentReportConfigurationProperty struct {
	// data_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#data_sources AwsExperimentTemplate#data_sources}
	// Experimental.
	DataSources *AwsExperimentTemplate_DataSourcesProperty `field:"optional" json:"dataSources" yaml:"dataSources"`
	// outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#outputs AwsExperimentTemplate#outputs}
	// Experimental.
	Outputs *AwsExperimentTemplate_OutputsProperty `field:"optional" json:"outputs" yaml:"outputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#post_experiment_duration AwsExperimentTemplate#post_experiment_duration}.
	// Experimental.
	PostExperimentDuration *string `field:"optional" json:"postExperimentDuration" yaml:"postExperimentDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#pre_experiment_duration AwsExperimentTemplate#pre_experiment_duration}.
	// Experimental.
	PreExperimentDuration *string `field:"optional" json:"preExperimentDuration" yaml:"preExperimentDuration"`
}

