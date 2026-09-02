package awsopensearch


// Experimental.
type TfDomain_MaintenanceScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#cron_expression_for_recurrence TfDomain#cron_expression_for_recurrence}.
	// Experimental.
	CronExpressionForRecurrence *string `field:"required" json:"cronExpressionForRecurrence" yaml:"cronExpressionForRecurrence"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#duration TfDomain#duration}
	// Experimental.
	Duration *TfDomain_DurationProperty `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#start_at TfDomain#start_at}.
	// Experimental.
	StartAt *string `field:"required" json:"startAt" yaml:"startAt"`
}

