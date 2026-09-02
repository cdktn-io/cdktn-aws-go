package awsverifiedpermissions


// Experimental.
type TfPolicy_StaticProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#statement TfPolicy#statement}.
	// Experimental.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#description TfPolicy#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

