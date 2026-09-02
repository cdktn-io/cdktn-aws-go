package awscodestarconnections


// Experimental.
type TfHost_VpcConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarconnections_host#security_group_ids TfHost#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarconnections_host#subnet_ids TfHost#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarconnections_host#vpc_id TfHost#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codestarconnections_host#tls_certificate TfHost#tls_certificate}.
	// Experimental.
	TlsCertificate *string `field:"optional" json:"tlsCertificate" yaml:"tlsCertificate"`
}

