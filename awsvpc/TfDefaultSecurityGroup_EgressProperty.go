package awsvpc


// Experimental.
type TfDefaultSecurityGroup_EgressProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#cidr_blocks TfDefaultSecurityGroup#cidr_blocks}.
	// Experimental.
	CidrBlocks *[]*string `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#description TfDefaultSecurityGroup#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#from_port TfDefaultSecurityGroup#from_port}.
	// Experimental.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#ipv6_cidr_blocks TfDefaultSecurityGroup#ipv6_cidr_blocks}.
	// Experimental.
	Ipv6CidrBlocks *[]*string `field:"optional" json:"ipv6CidrBlocks" yaml:"ipv6CidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#prefix_list_ids TfDefaultSecurityGroup#prefix_list_ids}.
	// Experimental.
	PrefixListIds *[]*string `field:"optional" json:"prefixListIds" yaml:"prefixListIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#protocol TfDefaultSecurityGroup#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#security_groups TfDefaultSecurityGroup#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#self TfDefaultSecurityGroup#self}.
	// Experimental.
	SelfAttribute interface{} `field:"optional" json:"selfAttribute" yaml:"selfAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_security_group#to_port TfDefaultSecurityGroup#to_port}.
	// Experimental.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

