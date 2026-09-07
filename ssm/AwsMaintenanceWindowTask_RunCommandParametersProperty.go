package ssm


// Experimental.
type AwsMaintenanceWindowTask_RunCommandParametersProperty struct {
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#cloudwatch_config AwsMaintenanceWindowTask#cloudwatch_config}
	// Experimental.
	CloudwatchConfig *AwsMaintenanceWindowTask_CloudwatchConfigProperty `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#comment AwsMaintenanceWindowTask#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_hash AwsMaintenanceWindowTask#document_hash}.
	// Experimental.
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_hash_type AwsMaintenanceWindowTask#document_hash_type}.
	// Experimental.
	DocumentHashType *string `field:"optional" json:"documentHashType" yaml:"documentHashType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_version AwsMaintenanceWindowTask#document_version}.
	// Experimental.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_config AwsMaintenanceWindowTask#notification_config}
	// Experimental.
	NotificationConfig *AwsMaintenanceWindowTask_NotificationConfigProperty `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#output_s3_bucket AwsMaintenanceWindowTask#output_s3_bucket}.
	// Experimental.
	OutputS3Bucket *string `field:"optional" json:"outputS3Bucket" yaml:"outputS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#output_s3_key_prefix AwsMaintenanceWindowTask#output_s3_key_prefix}.
	// Experimental.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#parameter AwsMaintenanceWindowTask#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#service_role_arn AwsMaintenanceWindowTask#service_role_arn}.
	// Experimental.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#timeout_seconds AwsMaintenanceWindowTask#timeout_seconds}.
	// Experimental.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

