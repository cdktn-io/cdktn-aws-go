package sagemakerai


// Experimental.
type AwsUserProfile_CustomPosixUserConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#gid AwsUserProfile#gid}.
	// Experimental.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#uid AwsUserProfile#uid}.
	// Experimental.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
}

