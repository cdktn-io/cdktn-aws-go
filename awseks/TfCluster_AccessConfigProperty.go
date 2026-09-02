package awseks


// Experimental.
type TfCluster_AccessConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#authentication_mode TfCluster#authentication_mode}.
	// Experimental.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#bootstrap_cluster_creator_admin_permissions TfCluster#bootstrap_cluster_creator_admin_permissions}.
	// Experimental.
	BootstrapClusterCreatorAdminPermissions interface{} `field:"optional" json:"bootstrapClusterCreatorAdminPermissions" yaml:"bootstrapClusterCreatorAdminPermissions"`
}

