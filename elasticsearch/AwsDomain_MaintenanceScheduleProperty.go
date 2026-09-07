package elasticsearch


// Experimental.
type AwsDomain_MaintenanceScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#cron_expression_for_recurrence AwsDomain#cron_expression_for_recurrence}.
	// Experimental.
	CronExpressionForRecurrence *string `field:"required" json:"cronExpressionForRecurrence" yaml:"cronExpressionForRecurrence"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#duration AwsDomain#duration}
	// Experimental.
	Duration *AwsDomain_DurationProperty `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#start_at AwsDomain#start_at}.
	// Experimental.
	StartAt *string `field:"required" json:"startAt" yaml:"startAt"`
}

