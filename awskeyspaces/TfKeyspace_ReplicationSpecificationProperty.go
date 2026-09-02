package awskeyspaces


// Experimental.
type TfKeyspace_ReplicationSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_keyspace#region_list TfKeyspace#region_list}.
	// Experimental.
	RegionList *[]*string `field:"optional" json:"regionList" yaml:"regionList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_keyspace#replication_strategy TfKeyspace#replication_strategy}.
	// Experimental.
	ReplicationStrategy *string `field:"optional" json:"replicationStrategy" yaml:"replicationStrategy"`
}

