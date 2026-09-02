package awssecretsmanager


// Experimental.
type TfSecretRotation_RotationRulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#automatically_after_days TfSecretRotation#automatically_after_days}.
	// Experimental.
	AutomaticallyAfterDays *float64 `field:"optional" json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#duration TfSecretRotation#duration}.
	// Experimental.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#schedule_expression TfSecretRotation#schedule_expression}.
	// Experimental.
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

