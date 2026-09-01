package awselb


// Experimental.
type AwsAlbTargetGroup_TargetGroupHealthProperty struct {
	// dns_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#dns_failover AwsAlbTargetGroup#dns_failover}
	// Experimental.
	DnsFailover *AwsAlbTargetGroup_DnsFailoverProperty `field:"optional" json:"dnsFailover" yaml:"dnsFailover"`
	// unhealthy_state_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#unhealthy_state_routing AwsAlbTargetGroup#unhealthy_state_routing}
	// Experimental.
	UnhealthyStateRouting *AwsAlbTargetGroup_UnhealthyStateRoutingProperty `field:"optional" json:"unhealthyStateRouting" yaml:"unhealthyStateRouting"`
}

