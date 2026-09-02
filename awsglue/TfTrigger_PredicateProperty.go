package awsglue


// Experimental.
type TfTrigger_PredicateProperty struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#conditions TfTrigger#conditions}
	// Experimental.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#logical TfTrigger#logical}.
	// Experimental.
	Logical *string `field:"optional" json:"logical" yaml:"logical"`
}

