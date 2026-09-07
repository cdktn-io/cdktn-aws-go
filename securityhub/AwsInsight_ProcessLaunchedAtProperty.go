package securityhub


// Experimental.
type AwsInsight_ProcessLaunchedAtProperty struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#date_range AwsInsight#date_range}
	// Experimental.
	DateRange *AwsInsight_FiltersProcessLaunchedAtDateRangeProperty `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#end AwsInsight#end}.
	// Experimental.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#start AwsInsight#start}.
	// Experimental.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

