package awsnetworkfirewall


// Experimental.
type TfRuleGroup_ReferenceSetsProperty struct {
	// ip_set_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#ip_set_references TfRuleGroup#ip_set_references}
	// Experimental.
	IpSetReferences interface{} `field:"optional" json:"ipSetReferences" yaml:"ipSetReferences"`
}

