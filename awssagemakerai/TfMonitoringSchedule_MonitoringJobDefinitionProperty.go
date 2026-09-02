package awssagemakerai


// Experimental.
type TfMonitoringSchedule_MonitoringJobDefinitionProperty struct {
	// monitoring_app_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_app_specification TfMonitoringSchedule#monitoring_app_specification}
	// Experimental.
	MonitoringAppSpecification *TfMonitoringSchedule_MonitoringAppSpecificationProperty `field:"required" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// monitoring_inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_inputs TfMonitoringSchedule#monitoring_inputs}
	// Experimental.
	MonitoringInputs *TfMonitoringSchedule_MonitoringInputsProperty `field:"required" json:"monitoringInputs" yaml:"monitoringInputs"`
	// monitoring_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_output_config TfMonitoringSchedule#monitoring_output_config}
	// Experimental.
	MonitoringOutputConfig *TfMonitoringSchedule_MonitoringOutputConfigProperty `field:"required" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// monitoring_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_resources TfMonitoringSchedule#monitoring_resources}
	// Experimental.
	MonitoringResources *TfMonitoringSchedule_MonitoringResourcesProperty `field:"required" json:"monitoringResources" yaml:"monitoringResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#role_arn TfMonitoringSchedule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// baseline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#baseline TfMonitoringSchedule#baseline}
	// Experimental.
	Baseline *TfMonitoringSchedule_BaselineProperty `field:"optional" json:"baseline" yaml:"baseline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#environment TfMonitoringSchedule#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#network_config TfMonitoringSchedule#network_config}
	// Experimental.
	NetworkConfig *TfMonitoringSchedule_NetworkConfigProperty `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#stopping_condition TfMonitoringSchedule#stopping_condition}
	// Experimental.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

