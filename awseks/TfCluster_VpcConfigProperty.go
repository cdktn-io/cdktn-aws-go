package awseks


// Experimental.
type TfCluster_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#subnet_ids TfCluster#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_egress_mode TfCluster#control_plane_egress_mode}.
	// Experimental.
	ControlPlaneEgressMode *string `field:"optional" json:"controlPlaneEgressMode" yaml:"controlPlaneEgressMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#endpoint_private_access TfCluster#endpoint_private_access}.
	// Experimental.
	EndpointPrivateAccess interface{} `field:"optional" json:"endpointPrivateAccess" yaml:"endpointPrivateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#endpoint_public_access TfCluster#endpoint_public_access}.
	// Experimental.
	EndpointPublicAccess interface{} `field:"optional" json:"endpointPublicAccess" yaml:"endpointPublicAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#public_access_cidrs TfCluster#public_access_cidrs}.
	// Experimental.
	PublicAccessCidrs *[]*string `field:"optional" json:"publicAccessCidrs" yaml:"publicAccessCidrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#security_group_ids TfCluster#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

