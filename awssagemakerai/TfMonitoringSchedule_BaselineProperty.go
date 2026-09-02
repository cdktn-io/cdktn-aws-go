package awssagemakerai


// Experimental.
type TfMonitoringSchedule_BaselineProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#baselining_job_name TfMonitoringSchedule#baselining_job_name}.
	// Experimental.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// constraints_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#constraints_resource TfMonitoringSchedule#constraints_resource}
	// Experimental.
	ConstraintsResource *TfMonitoringSchedule_ConstraintsResourceProperty `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// statistics_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#statistics_resource TfMonitoringSchedule#statistics_resource}
	// Experimental.
	StatisticsResource *TfMonitoringSchedule_StatisticsResourceProperty `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

