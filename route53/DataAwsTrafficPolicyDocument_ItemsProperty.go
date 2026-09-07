package route53


// Experimental.
type DataAwsTrafficPolicyDocument_ItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#endpoint_reference DataAwsTrafficPolicyDocument#endpoint_reference}.
	// Experimental.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#health_check DataAwsTrafficPolicyDocument#health_check}.
	// Experimental.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
}

