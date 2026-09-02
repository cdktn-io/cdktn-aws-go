package awsroute53


// Experimental.
type DataTfTrafficPolicyDocument_ItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#endpoint_reference DataTfTrafficPolicyDocument#endpoint_reference}.
	// Experimental.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#health_check DataTfTrafficPolicyDocument#health_check}.
	// Experimental.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
}

