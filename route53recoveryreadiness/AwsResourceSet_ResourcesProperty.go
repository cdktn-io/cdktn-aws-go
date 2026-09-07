package route53recoveryreadiness


// Experimental.
type AwsResourceSet_ResourcesProperty struct {
	// dns_target_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#dns_target_resource AwsResourceSet#dns_target_resource}
	// Experimental.
	DnsTargetResource *AwsResourceSet_DnsTargetResourceProperty `field:"optional" json:"dnsTargetResource" yaml:"dnsTargetResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#readiness_scopes AwsResourceSet#readiness_scopes}.
	// Experimental.
	ReadinessScopes *[]*string `field:"optional" json:"readinessScopes" yaml:"readinessScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#resource_arn AwsResourceSet#resource_arn}.
	// Experimental.
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

