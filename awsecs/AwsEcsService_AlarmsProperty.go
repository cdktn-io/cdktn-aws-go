package awsecs


// Experimental.
type AwsEcsService_AlarmsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#alarm_names AwsEcsService#alarm_names}.
	// Experimental.
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#enable AwsEcsService#enable}.
	// Experimental.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#rollback AwsEcsService#rollback}.
	// Experimental.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

