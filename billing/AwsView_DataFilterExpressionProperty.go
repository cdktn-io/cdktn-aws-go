package billing


// Experimental.
type AwsView_DataFilterExpressionProperty struct {
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#dimensions AwsView#dimensions}
	// Experimental.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#tags AwsView#tags}
	// Experimental.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// time_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#time_range AwsView#time_range}
	// Experimental.
	TimeRange interface{} `field:"optional" json:"timeRange" yaml:"timeRange"`
}

