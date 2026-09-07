package cloudhsm


// Experimental.
type AwsCluster_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudhsm_v2_cluster#create AwsCluster#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudhsm_v2_cluster#delete AwsCluster#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudhsm_v2_cluster#update AwsCluster#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

