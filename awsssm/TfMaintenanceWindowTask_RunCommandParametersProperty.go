package awsssm


// Experimental.
type TfMaintenanceWindowTask_RunCommandParametersProperty struct {
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#cloudwatch_config TfMaintenanceWindowTask#cloudwatch_config}
	// Experimental.
	CloudwatchConfig *TfMaintenanceWindowTask_CloudwatchConfigProperty `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#comment TfMaintenanceWindowTask#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_hash TfMaintenanceWindowTask#document_hash}.
	// Experimental.
	DocumentHash *string `field:"optional" json:"documentHash" yaml:"documentHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_hash_type TfMaintenanceWindowTask#document_hash_type}.
	// Experimental.
	DocumentHashType *string `field:"optional" json:"documentHashType" yaml:"documentHashType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_version TfMaintenanceWindowTask#document_version}.
	// Experimental.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_config TfMaintenanceWindowTask#notification_config}
	// Experimental.
	NotificationConfig *TfMaintenanceWindowTask_NotificationConfigProperty `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#output_s3_bucket TfMaintenanceWindowTask#output_s3_bucket}.
	// Experimental.
	OutputS3Bucket *string `field:"optional" json:"outputS3Bucket" yaml:"outputS3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#output_s3_key_prefix TfMaintenanceWindowTask#output_s3_key_prefix}.
	// Experimental.
	OutputS3KeyPrefix *string `field:"optional" json:"outputS3KeyPrefix" yaml:"outputS3KeyPrefix"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#parameter TfMaintenanceWindowTask#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#service_role_arn TfMaintenanceWindowTask#service_role_arn}.
	// Experimental.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#timeout_seconds TfMaintenanceWindowTask#timeout_seconds}.
	// Experimental.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

