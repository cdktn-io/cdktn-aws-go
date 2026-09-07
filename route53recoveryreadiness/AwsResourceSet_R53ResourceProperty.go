package route53recoveryreadiness


// Experimental.
type AwsResourceSet_R53ResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#domain_name AwsResourceSet#domain_name}.
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoveryreadiness_resource_set#record_set_id AwsResourceSet#record_set_id}.
	// Experimental.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
}

