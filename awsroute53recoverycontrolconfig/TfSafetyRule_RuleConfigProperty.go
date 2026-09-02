package awsroute53recoverycontrolconfig


// Experimental.
type TfSafetyRule_RuleConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#inverted TfSafetyRule#inverted}.
	// Experimental.
	Inverted interface{} `field:"required" json:"inverted" yaml:"inverted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#threshold TfSafetyRule#threshold}.
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#type TfSafetyRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

