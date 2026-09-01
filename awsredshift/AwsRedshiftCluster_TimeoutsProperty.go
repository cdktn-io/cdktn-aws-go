package awsredshift


// Experimental.
type AwsRedshiftCluster_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#create AwsRedshiftCluster#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#delete AwsRedshiftCluster#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#update AwsRedshiftCluster#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

