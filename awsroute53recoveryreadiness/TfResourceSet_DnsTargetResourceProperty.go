package awsroute53recoveryreadiness


// Experimental.
type TfResourceSet_DnsTargetResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#domain_name TfResourceSet#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#hosted_zone_arn TfResourceSet#hosted_zone_arn}.
	// Experimental.
	HostedZoneArn *string `field:"optional" json:"hostedZoneArn" yaml:"hostedZoneArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#record_set_id TfResourceSet#record_set_id}.
	// Experimental.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#record_type TfResourceSet#record_type}.
	// Experimental.
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// target_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#target_resource TfResourceSet#target_resource}
	// Experimental.
	TargetResource *TfResourceSet_TargetResourceProperty `field:"optional" json:"targetResource" yaml:"targetResource"`
}

