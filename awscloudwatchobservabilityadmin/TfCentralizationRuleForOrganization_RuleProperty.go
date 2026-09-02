package awscloudwatchobservabilityadmin


// Experimental.
type TfCentralizationRuleForOrganization_RuleProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#destination TfCentralizationRuleForOrganization#destination}
	// Experimental.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#source TfCentralizationRuleForOrganization#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

