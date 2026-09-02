package awsopensearch


// Experimental.
type TfApplication_IamIdentityCenterOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_application#enabled TfApplication#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_application#iam_identity_center_instance_arn TfApplication#iam_identity_center_instance_arn}.
	// Experimental.
	IamIdentityCenterInstanceArn *string `field:"optional" json:"iamIdentityCenterInstanceArn" yaml:"iamIdentityCenterInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_application#iam_role_for_identity_center_application_arn TfApplication#iam_role_for_identity_center_application_arn}.
	// Experimental.
	IamRoleForIdentityCenterApplicationArn *string `field:"optional" json:"iamRoleForIdentityCenterApplicationArn" yaml:"iamRoleForIdentityCenterApplicationArn"`
}

