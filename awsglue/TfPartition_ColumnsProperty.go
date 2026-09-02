package awsglue


// Experimental.
type TfPartition_ColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#name TfPartition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#comment TfPartition#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#type TfPartition#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

