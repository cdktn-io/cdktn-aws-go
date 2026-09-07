package recyclebin


// Experimental.
type AwsRule_RetentionPeriodProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#retention_period_unit AwsRule#retention_period_unit}.
	// Experimental.
	RetentionPeriodUnit *string `field:"required" json:"retentionPeriodUnit" yaml:"retentionPeriodUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#retention_period_value AwsRule#retention_period_value}.
	// Experimental.
	RetentionPeriodValue *float64 `field:"required" json:"retentionPeriodValue" yaml:"retentionPeriodValue"`
}

