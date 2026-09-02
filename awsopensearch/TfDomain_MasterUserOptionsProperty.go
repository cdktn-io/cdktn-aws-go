package awsopensearch


// Experimental.
type TfDomain_MasterUserOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#master_user_arn TfDomain#master_user_arn}.
	// Experimental.
	MasterUserArn *string `field:"optional" json:"masterUserArn" yaml:"masterUserArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#master_user_name TfDomain#master_user_name}.
	// Experimental.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#master_user_password TfDomain#master_user_password}.
	// Experimental.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
}

