package awsssm


// Experimental.
type AwsSsmMaintenanceWindowTask_NotificationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_arn AwsSsmMaintenanceWindowTask#notification_arn}.
	// Experimental.
	NotificationArn *string `field:"optional" json:"notificationArn" yaml:"notificationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_events AwsSsmMaintenanceWindowTask#notification_events}.
	// Experimental.
	NotificationEvents *[]*string `field:"optional" json:"notificationEvents" yaml:"notificationEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#notification_type AwsSsmMaintenanceWindowTask#notification_type}.
	// Experimental.
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
}

