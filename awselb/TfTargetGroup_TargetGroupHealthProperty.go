package awselb


// Experimental.
type TfTargetGroup_TargetGroupHealthProperty struct {
	// dns_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#dns_failover TfTargetGroup#dns_failover}
	// Experimental.
	DnsFailover *TfTargetGroup_DnsFailoverProperty `field:"optional" json:"dnsFailover" yaml:"dnsFailover"`
	// unhealthy_state_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#unhealthy_state_routing TfTargetGroup#unhealthy_state_routing}
	// Experimental.
	UnhealthyStateRouting *TfTargetGroup_UnhealthyStateRoutingProperty `field:"optional" json:"unhealthyStateRouting" yaml:"unhealthyStateRouting"`
}

