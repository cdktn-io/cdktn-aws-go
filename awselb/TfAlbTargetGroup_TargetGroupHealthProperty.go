package awselb


// Experimental.
type TfAlbTargetGroup_TargetGroupHealthProperty struct {
	// dns_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#dns_failover TfAlbTargetGroup#dns_failover}
	// Experimental.
	DnsFailover *TfAlbTargetGroup_DnsFailoverProperty `field:"optional" json:"dnsFailover" yaml:"dnsFailover"`
	// unhealthy_state_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#unhealthy_state_routing TfAlbTargetGroup#unhealthy_state_routing}
	// Experimental.
	UnhealthyStateRouting *TfAlbTargetGroup_UnhealthyStateRoutingProperty `field:"optional" json:"unhealthyStateRouting" yaml:"unhealthyStateRouting"`
}

