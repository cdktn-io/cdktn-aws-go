package arczonalshift


// Experimental.
type AwsZonalAutoshiftConfiguration_BlockingAlarmsProperty struct {
	// ARN of the CloudWatch alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arczonalshift_zonal_autoshift_configuration#alarm_identifier AwsZonalAutoshiftConfiguration#alarm_identifier}
	// Experimental.
	AlarmIdentifier *string `field:"required" json:"alarmIdentifier" yaml:"alarmIdentifier"`
	// Type of control condition. Valid value: `CLOUDWATCH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arczonalshift_zonal_autoshift_configuration#type AwsZonalAutoshiftConfiguration#type}
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

