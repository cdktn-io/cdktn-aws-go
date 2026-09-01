package awssagemakerai


// Experimental.
type AwsSagemakerUserProfile_CustomPosixUserConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#gid AwsSagemakerUserProfile#gid}.
	// Experimental.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#uid AwsSagemakerUserProfile#uid}.
	// Experimental.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
}

