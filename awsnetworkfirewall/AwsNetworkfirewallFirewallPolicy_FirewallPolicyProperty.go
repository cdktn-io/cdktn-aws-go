package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallFirewallPolicy_FirewallPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stateless_default_actions AwsNetworkfirewallFirewallPolicy#stateless_default_actions}.
	// Experimental.
	StatelessDefaultActions *[]*string `field:"required" json:"statelessDefaultActions" yaml:"statelessDefaultActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stateless_fragment_default_actions AwsNetworkfirewallFirewallPolicy#stateless_fragment_default_actions}.
	// Experimental.
	StatelessFragmentDefaultActions *[]*string `field:"required" json:"statelessFragmentDefaultActions" yaml:"statelessFragmentDefaultActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#enable_tls_session_holding AwsNetworkfirewallFirewallPolicy#enable_tls_session_holding}.
	// Experimental.
	EnableTlsSessionHolding interface{} `field:"optional" json:"enableTlsSessionHolding" yaml:"enableTlsSessionHolding"`
	// policy_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#policy_variables AwsNetworkfirewallFirewallPolicy#policy_variables}
	// Experimental.
	PolicyVariables *AwsNetworkfirewallFirewallPolicy_PolicyVariablesProperty `field:"optional" json:"policyVariables" yaml:"policyVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stateful_default_actions AwsNetworkfirewallFirewallPolicy#stateful_default_actions}.
	// Experimental.
	StatefulDefaultActions *[]*string `field:"optional" json:"statefulDefaultActions" yaml:"statefulDefaultActions"`
	// stateful_engine_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stateful_engine_options AwsNetworkfirewallFirewallPolicy#stateful_engine_options}
	// Experimental.
	StatefulEngineOptions *AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsProperty `field:"optional" json:"statefulEngineOptions" yaml:"statefulEngineOptions"`
	// stateful_rule_group_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stateful_rule_group_reference AwsNetworkfirewallFirewallPolicy#stateful_rule_group_reference}
	// Experimental.
	StatefulRuleGroupReference interface{} `field:"optional" json:"statefulRuleGroupReference" yaml:"statefulRuleGroupReference"`
	// stateless_custom_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stateless_custom_action AwsNetworkfirewallFirewallPolicy#stateless_custom_action}
	// Experimental.
	StatelessCustomAction interface{} `field:"optional" json:"statelessCustomAction" yaml:"statelessCustomAction"`
	// stateless_rule_group_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#stateless_rule_group_reference AwsNetworkfirewallFirewallPolicy#stateless_rule_group_reference}
	// Experimental.
	StatelessRuleGroupReference interface{} `field:"optional" json:"statelessRuleGroupReference" yaml:"statelessRuleGroupReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall_policy#tls_inspection_configuration_arn AwsNetworkfirewallFirewallPolicy#tls_inspection_configuration_arn}.
	// Experimental.
	TlsInspectionConfigurationArn *string `field:"optional" json:"tlsInspectionConfigurationArn" yaml:"tlsInspectionConfigurationArn"`
}

