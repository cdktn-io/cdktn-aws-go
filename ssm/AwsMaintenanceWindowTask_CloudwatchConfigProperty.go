package ssm


// Experimental.
type AwsMaintenanceWindowTask_CloudwatchConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#cloudwatch_log_group_name AwsMaintenanceWindowTask#cloudwatch_log_group_name}.
	// Experimental.
	CloudwatchLogGroupName *string `field:"optional" json:"cloudwatchLogGroupName" yaml:"cloudwatchLogGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#cloudwatch_output_enabled AwsMaintenanceWindowTask#cloudwatch_output_enabled}.
	// Experimental.
	CloudwatchOutputEnabled interface{} `field:"optional" json:"cloudwatchOutputEnabled" yaml:"cloudwatchOutputEnabled"`
}

