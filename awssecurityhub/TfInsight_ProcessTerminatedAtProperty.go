package awssecurityhub


// Experimental.
type TfInsight_ProcessTerminatedAtProperty struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#date_range TfInsight#date_range}
	// Experimental.
	DateRange *TfInsight_FiltersProcessTerminatedAtDateRangeProperty `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#end TfInsight#end}.
	// Experimental.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#start TfInsight#start}.
	// Experimental.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

