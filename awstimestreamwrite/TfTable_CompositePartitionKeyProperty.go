package awstimestreamwrite


// Experimental.
type TfTable_CompositePartitionKeyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#type TfTable#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#enforcement_in_record TfTable#enforcement_in_record}.
	// Experimental.
	EnforcementInRecord *string `field:"optional" json:"enforcementInRecord" yaml:"enforcementInRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#name TfTable#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

