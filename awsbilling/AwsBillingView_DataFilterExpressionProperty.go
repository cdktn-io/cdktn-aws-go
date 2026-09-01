package awsbilling


// Experimental.
type AwsBillingView_DataFilterExpressionProperty struct {
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#dimensions AwsBillingView#dimensions}
	// Experimental.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#tags AwsBillingView#tags}
	// Experimental.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// time_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#time_range AwsBillingView#time_range}
	// Experimental.
	TimeRange interface{} `field:"optional" json:"timeRange" yaml:"timeRange"`
}

