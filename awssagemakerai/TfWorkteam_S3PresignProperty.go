package awssagemakerai


// Experimental.
type TfWorkteam_S3PresignProperty struct {
	// iam_policy_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#iam_policy_constraints TfWorkteam#iam_policy_constraints}
	// Experimental.
	IamPolicyConstraints *TfWorkteam_IamPolicyConstraintsProperty `field:"optional" json:"iamPolicyConstraints" yaml:"iamPolicyConstraints"`
}

