package awsredshift


// Experimental.
type TfClusterIamRoles_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster_iam_roles#create TfClusterIamRoles#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster_iam_roles#delete TfClusterIamRoles#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster_iam_roles#update TfClusterIamRoles#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

