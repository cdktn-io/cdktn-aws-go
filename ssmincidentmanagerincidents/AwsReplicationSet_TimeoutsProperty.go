package ssmincidentmanagerincidents


// Experimental.
type AwsReplicationSet_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_replication_set#create AwsReplicationSet#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_replication_set#delete AwsReplicationSet#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_replication_set#update AwsReplicationSet#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

