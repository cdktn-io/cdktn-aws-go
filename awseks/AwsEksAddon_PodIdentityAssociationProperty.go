package awseks


// Experimental.
type AwsEksAddon_PodIdentityAssociationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#role_arn AwsEksAddon#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#service_account AwsEksAddon#service_account}.
	// Experimental.
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

