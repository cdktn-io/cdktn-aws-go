package timestreamwrite


// Experimental.
type AwsTable_CompositePartitionKeyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#type AwsTable#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#enforcement_in_record AwsTable#enforcement_in_record}.
	// Experimental.
	EnforcementInRecord *string `field:"optional" json:"enforcementInRecord" yaml:"enforcementInRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#name AwsTable#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

