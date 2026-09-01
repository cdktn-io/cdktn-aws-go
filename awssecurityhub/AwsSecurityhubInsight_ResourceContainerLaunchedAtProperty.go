package awssecurityhub


// Experimental.
type AwsSecurityhubInsight_ResourceContainerLaunchedAtProperty struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#date_range AwsSecurityhubInsight#date_range}
	// Experimental.
	DateRange *AwsSecurityhubInsight_FiltersResourceContainerLaunchedAtDateRangeProperty `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#end AwsSecurityhubInsight#end}.
	// Experimental.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#start AwsSecurityhubInsight#start}.
	// Experimental.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

