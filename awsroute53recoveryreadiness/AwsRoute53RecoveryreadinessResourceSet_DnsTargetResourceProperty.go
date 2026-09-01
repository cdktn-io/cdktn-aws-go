package awsroute53recoveryreadiness


// Experimental.
type AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#domain_name AwsRoute53RecoveryreadinessResourceSet#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#hosted_zone_arn AwsRoute53RecoveryreadinessResourceSet#hosted_zone_arn}.
	// Experimental.
	HostedZoneArn *string `field:"optional" json:"hostedZoneArn" yaml:"hostedZoneArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#record_set_id AwsRoute53RecoveryreadinessResourceSet#record_set_id}.
	// Experimental.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#record_type AwsRoute53RecoveryreadinessResourceSet#record_type}.
	// Experimental.
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// target_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#target_resource AwsRoute53RecoveryreadinessResourceSet#target_resource}
	// Experimental.
	TargetResource *AwsRoute53RecoveryreadinessResourceSet_TargetResourceProperty `field:"optional" json:"targetResource" yaml:"targetResource"`
}

