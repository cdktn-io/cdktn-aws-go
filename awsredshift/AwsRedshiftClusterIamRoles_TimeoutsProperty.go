package awsredshift


// Experimental.
type AwsRedshiftClusterIamRoles_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster_iam_roles#create AwsRedshiftClusterIamRoles#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster_iam_roles#delete AwsRedshiftClusterIamRoles#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster_iam_roles#update AwsRedshiftClusterIamRoles#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

