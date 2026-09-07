package sagemakerai


// Experimental.
type AwsDataQualityJobDefinition_DataQualityBaselineConfigProperty struct {
	// constraints_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#constraints_resource AwsDataQualityJobDefinition#constraints_resource}
	// Experimental.
	ConstraintsResource *AwsDataQualityJobDefinition_ConstraintsResourceProperty `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// statistics_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#statistics_resource AwsDataQualityJobDefinition#statistics_resource}
	// Experimental.
	StatisticsResource *AwsDataQualityJobDefinition_StatisticsResourceProperty `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

