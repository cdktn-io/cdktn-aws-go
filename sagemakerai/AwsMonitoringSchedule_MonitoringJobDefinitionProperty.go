package sagemakerai


// Experimental.
type AwsMonitoringSchedule_MonitoringJobDefinitionProperty struct {
	// monitoring_app_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_app_specification AwsMonitoringSchedule#monitoring_app_specification}
	// Experimental.
	MonitoringAppSpecification *AwsMonitoringSchedule_MonitoringAppSpecificationProperty `field:"required" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// monitoring_inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_inputs AwsMonitoringSchedule#monitoring_inputs}
	// Experimental.
	MonitoringInputs *AwsMonitoringSchedule_MonitoringInputsProperty `field:"required" json:"monitoringInputs" yaml:"monitoringInputs"`
	// monitoring_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_output_config AwsMonitoringSchedule#monitoring_output_config}
	// Experimental.
	MonitoringOutputConfig *AwsMonitoringSchedule_MonitoringOutputConfigProperty `field:"required" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// monitoring_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_resources AwsMonitoringSchedule#monitoring_resources}
	// Experimental.
	MonitoringResources *AwsMonitoringSchedule_MonitoringResourcesProperty `field:"required" json:"monitoringResources" yaml:"monitoringResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#role_arn AwsMonitoringSchedule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// baseline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#baseline AwsMonitoringSchedule#baseline}
	// Experimental.
	Baseline *AwsMonitoringSchedule_BaselineProperty `field:"optional" json:"baseline" yaml:"baseline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#environment AwsMonitoringSchedule#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#network_config AwsMonitoringSchedule#network_config}
	// Experimental.
	NetworkConfig *AwsMonitoringSchedule_NetworkConfigProperty `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#stopping_condition AwsMonitoringSchedule#stopping_condition}
	// Experimental.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

