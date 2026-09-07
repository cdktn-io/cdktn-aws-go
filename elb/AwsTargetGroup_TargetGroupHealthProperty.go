package elb


// Experimental.
type AwsTargetGroup_TargetGroupHealthProperty struct {
	// dns_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#dns_failover AwsTargetGroup#dns_failover}
	// Experimental.
	DnsFailover *AwsTargetGroup_DnsFailoverProperty `field:"optional" json:"dnsFailover" yaml:"dnsFailover"`
	// unhealthy_state_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#unhealthy_state_routing AwsTargetGroup#unhealthy_state_routing}
	// Experimental.
	UnhealthyStateRouting *AwsTargetGroup_UnhealthyStateRoutingProperty `field:"optional" json:"unhealthyStateRouting" yaml:"unhealthyStateRouting"`
}

