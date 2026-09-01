package awssagemakerai


// Experimental.
type AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty struct {
	// monitoring_app_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_app_specification AwsSagemakerMonitoringSchedule#monitoring_app_specification}
	// Experimental.
	MonitoringAppSpecification *AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationProperty `field:"required" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// monitoring_inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_inputs AwsSagemakerMonitoringSchedule#monitoring_inputs}
	// Experimental.
	MonitoringInputs *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty `field:"required" json:"monitoringInputs" yaml:"monitoringInputs"`
	// monitoring_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_output_config AwsSagemakerMonitoringSchedule#monitoring_output_config}
	// Experimental.
	MonitoringOutputConfig *AwsSagemakerMonitoringSchedule_MonitoringOutputConfigProperty `field:"required" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// monitoring_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_resources AwsSagemakerMonitoringSchedule#monitoring_resources}
	// Experimental.
	MonitoringResources *AwsSagemakerMonitoringSchedule_MonitoringResourcesProperty `field:"required" json:"monitoringResources" yaml:"monitoringResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#role_arn AwsSagemakerMonitoringSchedule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// baseline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#baseline AwsSagemakerMonitoringSchedule#baseline}
	// Experimental.
	Baseline *AwsSagemakerMonitoringSchedule_BaselineProperty `field:"optional" json:"baseline" yaml:"baseline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#environment AwsSagemakerMonitoringSchedule#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#network_config AwsSagemakerMonitoringSchedule#network_config}
	// Experimental.
	NetworkConfig *AwsSagemakerMonitoringSchedule_NetworkConfigProperty `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#stopping_condition AwsSagemakerMonitoringSchedule#stopping_condition}
	// Experimental.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

