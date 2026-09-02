package awsroute53resolver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResolverFirewallRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#action TfResolverFirewallRule#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#firewall_rule_group_id TfResolverFirewallRule#firewall_rule_group_id}.
	// Experimental.
	FirewallRuleGroupId *string `field:"required" json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#name TfResolverFirewallRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#priority TfResolverFirewallRule#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#block_override_dns_type TfResolverFirewallRule#block_override_dns_type}.
	// Experimental.
	BlockOverrideDnsType *string `field:"optional" json:"blockOverrideDnsType" yaml:"blockOverrideDnsType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#block_override_domain TfResolverFirewallRule#block_override_domain}.
	// Experimental.
	BlockOverrideDomain *string `field:"optional" json:"blockOverrideDomain" yaml:"blockOverrideDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#block_override_ttl TfResolverFirewallRule#block_override_ttl}.
	// Experimental.
	BlockOverrideTtl *float64 `field:"optional" json:"blockOverrideTtl" yaml:"blockOverrideTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#block_response TfResolverFirewallRule#block_response}.
	// Experimental.
	BlockResponse *string `field:"optional" json:"blockResponse" yaml:"blockResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#confidence_threshold TfResolverFirewallRule#confidence_threshold}.
	// Experimental.
	ConfidenceThreshold *string `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#dns_threat_protection TfResolverFirewallRule#dns_threat_protection}.
	// Experimental.
	DnsThreatProtection *string `field:"optional" json:"dnsThreatProtection" yaml:"dnsThreatProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#firewall_domain_list_id TfResolverFirewallRule#firewall_domain_list_id}.
	// Experimental.
	FirewallDomainListId *string `field:"optional" json:"firewallDomainListId" yaml:"firewallDomainListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#firewall_domain_redirection_action TfResolverFirewallRule#firewall_domain_redirection_action}.
	// Experimental.
	FirewallDomainRedirectionAction *string `field:"optional" json:"firewallDomainRedirectionAction" yaml:"firewallDomainRedirectionAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#id TfResolverFirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#q_type TfResolverFirewallRule#q_type}.
	// Experimental.
	QType *string `field:"optional" json:"qType" yaml:"qType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule#region TfResolverFirewallRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

