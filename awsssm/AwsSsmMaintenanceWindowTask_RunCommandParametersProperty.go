package awsssm


// Experimental.
type AwsSsmMaintenanceWindowTask_RunCommandParametersProperty struct {
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#cloudwatch_config AwsSsmMaintenanceWindowTask#cloudwatch_config}
	// Experimental.
	CloudwatchConfig *AwsSsmMaintenanceWindowTask_CloudwatchConfigProperty `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#comment AwsSsmMaintenanceWindowTask#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_hash AwsSsmMaintenanceWindowTask#document_hash}.
	// Experimental.
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_hash_type AwsSsmMaintenanceWindowTask#document_hash_type}.
	// Experimental.
	DocumentHashType *string `field:"optional" json:"documentHashType" yaml:"documentHashType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_version AwsSsmMaintenanceWindowTask#document_version}.
	// Experimental.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_config AwsSsmMaintenanceWindowTask#notification_config}
	// Experimental.
	NotificationConfig *AwsSsmMaintenanceWindowTask_NotificationConfigProperty `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#output_s3_bucket AwsSsmMaintenanceWindowTask#output_s3_bucket}.
	// Experimental.
	OutputS3Bucket *string `field:"optional" json:"outputS3Bucket" yaml:"outputS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#output_s3_key_prefix AwsSsmMaintenanceWindowTask#output_s3_key_prefix}.
	// Experimental.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#parameter AwsSsmMaintenanceWindowTask#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#service_role_arn AwsSsmMaintenanceWindowTask#service_role_arn}.
	// Experimental.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#timeout_seconds AwsSsmMaintenanceWindowTask#timeout_seconds}.
	// Experimental.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

