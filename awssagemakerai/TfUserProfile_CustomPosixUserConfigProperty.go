package awssagemakerai


// Experimental.
type TfUserProfile_CustomPosixUserConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#gid TfUserProfile#gid}.
	// Experimental.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#uid TfUserProfile#uid}.
	// Experimental.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
}

