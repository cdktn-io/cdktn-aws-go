package awssagemakerai


// Experimental.
type TfMonitoringSchedule_MonitoringScheduleConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_type TfMonitoringSchedule#monitoring_type}.
	// Experimental.
	MonitoringType *string `field:"required" json:"monitoringType" yaml:"monitoringType"`
	// monitoring_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_job_definition TfMonitoringSchedule#monitoring_job_definition}
	// Experimental.
	MonitoringJobDefinition *TfMonitoringSchedule_MonitoringJobDefinitionProperty `field:"optional" json:"monitoringJobDefinition" yaml:"monitoringJobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_job_definition_name TfMonitoringSchedule#monitoring_job_definition_name}.
	// Experimental.
	MonitoringJobDefinitionName *string `field:"optional" json:"monitoringJobDefinitionName" yaml:"monitoringJobDefinitionName"`
	// schedule_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#schedule_config TfMonitoringSchedule#schedule_config}
	// Experimental.
	ScheduleConfig *TfMonitoringSchedule_ScheduleConfigProperty `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
}

