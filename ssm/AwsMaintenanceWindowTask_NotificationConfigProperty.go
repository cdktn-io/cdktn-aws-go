package ssm


// Experimental.
type AwsMaintenanceWindowTask_NotificationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_arn AwsMaintenanceWindowTask#notification_arn}.
	// Experimental.
	NotificationArn *string `field:"optional" json:"notificationArn" yaml:"notificationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_events AwsMaintenanceWindowTask#notification_events}.
	// Experimental.
	NotificationEvents *[]*string `field:"optional" json:"notificationEvents" yaml:"notificationEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_type AwsMaintenanceWindowTask#notification_type}.
	// Experimental.
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
}

