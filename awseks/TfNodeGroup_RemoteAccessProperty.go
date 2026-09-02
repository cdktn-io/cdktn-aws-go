package awseks


// Experimental.
type TfNodeGroup_RemoteAccessProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#ec2_ssh_key TfNodeGroup#ec2_ssh_key}.
	// Experimental.
	Ec2SshKey *string `field:"optional" json:"ec2SshKey" yaml:"ec2SshKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#source_security_group_ids TfNodeGroup#source_security_group_ids}.
	// Experimental.
	SourceSecurityGroupIds *[]*string `field:"optional" json:"sourceSecurityGroupIds" yaml:"sourceSecurityGroupIds"`
}

