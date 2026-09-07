package resiliencehubv2


// Experimental.
type AwsService_CrossAccountRoleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_service#cross_account_role_arn AwsService#cross_account_role_arn}.
	// Experimental.
	CrossAccountRoleArn *string `field:"required" json:"crossAccountRoleArn" yaml:"crossAccountRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_service#external_id AwsService#external_id}.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

