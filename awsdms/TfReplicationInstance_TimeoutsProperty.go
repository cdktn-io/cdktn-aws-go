package awsdms


// Experimental.
type TfReplicationInstance_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#create TfReplicationInstance#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#delete TfReplicationInstance#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#update TfReplicationInstance#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

