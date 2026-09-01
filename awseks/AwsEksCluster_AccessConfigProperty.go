package awseks


// Experimental.
type AwsEksCluster_AccessConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#authentication_mode AwsEksCluster#authentication_mode}.
	// Experimental.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#bootstrap_cluster_creator_admin_permissions AwsEksCluster#bootstrap_cluster_creator_admin_permissions}.
	// Experimental.
	BootstrapClusterCreatorAdminPermissions interface{} `field:"optional" json:"bootstrapClusterCreatorAdminPermissions" yaml:"bootstrapClusterCreatorAdminPermissions"`
}

