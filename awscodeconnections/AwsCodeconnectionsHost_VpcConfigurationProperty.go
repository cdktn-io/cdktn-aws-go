package awscodeconnections


// Experimental.
type AwsCodeconnectionsHost_VpcConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeconnections_host#security_group_ids AwsCodeconnectionsHost#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeconnections_host#subnet_ids AwsCodeconnectionsHost#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeconnections_host#vpc_id AwsCodeconnectionsHost#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codeconnections_host#tls_certificate AwsCodeconnectionsHost#tls_certificate}.
	// Experimental.
	TlsCertificate *string `field:"optional" json:"tlsCertificate" yaml:"tlsCertificate"`
}

