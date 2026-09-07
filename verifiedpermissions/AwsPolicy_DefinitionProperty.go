package verifiedpermissions


// Experimental.
type AwsPolicy_DefinitionProperty struct {
	// static block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#static AwsPolicy#static}
	// Experimental.
	Static interface{} `field:"optional" json:"static" yaml:"static"`
	// template_linked block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#template_linked AwsPolicy#template_linked}
	// Experimental.
	TemplateLinked interface{} `field:"optional" json:"templateLinked" yaml:"templateLinked"`
}

