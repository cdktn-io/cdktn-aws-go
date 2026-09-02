package awsroute53


// Experimental.
type TfRecordsExclusive_AliasTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#dns_name TfRecordsExclusive#dns_name}.
	// Experimental.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#evaluate_target_health TfRecordsExclusive#evaluate_target_health}.
	// Experimental.
	EvaluateTargetHealth interface{} `field:"required" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#hosted_zone_id TfRecordsExclusive#hosted_zone_id}.
	// Experimental.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

