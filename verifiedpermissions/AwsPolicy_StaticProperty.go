package verifiedpermissions


// Experimental.
type AwsPolicy_StaticProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#statement AwsPolicy#statement}.
	// Experimental.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#description AwsPolicy#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

