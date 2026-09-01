package awseks


// Experimental.
type AwsEksCapability_ArgoCdProperty struct {
	// aws_idc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#aws_idc AwsEksCapability#aws_idc}
	// Experimental.
	AwsIdc interface{} `field:"optional" json:"awsIdc" yaml:"awsIdc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#namespace AwsEksCapability#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// network_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#network_access AwsEksCapability#network_access}
	// Experimental.
	NetworkAccess interface{} `field:"optional" json:"networkAccess" yaml:"networkAccess"`
	// rbac_role_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#rbac_role_mapping AwsEksCapability#rbac_role_mapping}
	// Experimental.
	RbacRoleMapping interface{} `field:"optional" json:"rbacRoleMapping" yaml:"rbacRoleMapping"`
}

