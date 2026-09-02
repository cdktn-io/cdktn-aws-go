package awssagemakerai


// Experimental.
type TfDataQualityJobDefinition_DataQualityBaselineConfigProperty struct {
	// constraints_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#constraints_resource TfDataQualityJobDefinition#constraints_resource}
	// Experimental.
	ConstraintsResource *TfDataQualityJobDefinition_ConstraintsResourceProperty `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// statistics_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#statistics_resource TfDataQualityJobDefinition#statistics_resource}
	// Experimental.
	StatisticsResource *TfDataQualityJobDefinition_StatisticsResourceProperty `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

