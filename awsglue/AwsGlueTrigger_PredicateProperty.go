package awsglue


// Experimental.
type AwsGlueTrigger_PredicateProperty struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#conditions AwsGlueTrigger#conditions}
	// Experimental.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#logical AwsGlueTrigger#logical}.
	// Experimental.
	Logical *string `field:"optional" json:"logical" yaml:"logical"`
}

