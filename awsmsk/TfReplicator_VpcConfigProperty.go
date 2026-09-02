package awsmsk


// Experimental.
type TfReplicator_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#subnet_ids TfReplicator#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#security_groups_ids TfReplicator#security_groups_ids}.
	// Experimental.
	SecurityGroupsIds *[]*string `field:"optional" json:"securityGroupsIds" yaml:"securityGroupsIds"`
}

