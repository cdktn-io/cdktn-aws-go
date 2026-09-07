package sagemakerai


// Experimental.
type AwsMonitoringSchedule_BaselineProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#baselining_job_name AwsMonitoringSchedule#baselining_job_name}.
	// Experimental.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// constraints_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#constraints_resource AwsMonitoringSchedule#constraints_resource}
	// Experimental.
	ConstraintsResource *AwsMonitoringSchedule_ConstraintsResourceProperty `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// statistics_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#statistics_resource AwsMonitoringSchedule#statistics_resource}
	// Experimental.
	StatisticsResource *AwsMonitoringSchedule_StatisticsResourceProperty `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

