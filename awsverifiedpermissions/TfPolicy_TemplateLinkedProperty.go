package awsverifiedpermissions


// Experimental.
type TfPolicy_TemplateLinkedProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#policy_template_id TfPolicy#policy_template_id}.
	// Experimental.
	PolicyTemplateId *string `field:"required" json:"policyTemplateId" yaml:"policyTemplateId"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#principal TfPolicy#principal}
	// Experimental.
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#resource TfPolicy#resource}
	// Experimental.
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}

