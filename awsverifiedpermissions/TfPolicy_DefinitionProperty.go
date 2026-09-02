package awsverifiedpermissions


// Experimental.
type TfPolicy_DefinitionProperty struct {
	// static block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#static TfPolicy#static}
	// Experimental.
	Static interface{} `field:"optional" json:"static" yaml:"static"`
	// template_linked block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#template_linked TfPolicy#template_linked}
	// Experimental.
	TemplateLinked interface{} `field:"optional" json:"templateLinked" yaml:"templateLinked"`
}

