package codegurureviewer


// Experimental.
type AwsRepositoryAssociation_RepositoryProperty struct {
	// bitbucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#bitbucket AwsRepositoryAssociation#bitbucket}
	// Experimental.
	Bitbucket *AwsRepositoryAssociation_BitbucketProperty `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// codecommit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#codecommit AwsRepositoryAssociation#codecommit}
	// Experimental.
	Codecommit *AwsRepositoryAssociation_CodecommitProperty `field:"optional" json:"codecommit" yaml:"codecommit"`
	// github_enterprise_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#github_enterprise_server AwsRepositoryAssociation#github_enterprise_server}
	// Experimental.
	GithubEnterpriseServer *AwsRepositoryAssociation_GithubEnterpriseServerProperty `field:"optional" json:"githubEnterpriseServer" yaml:"githubEnterpriseServer"`
	// s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#s3_bucket AwsRepositoryAssociation#s3_bucket}
	// Experimental.
	S3Bucket *AwsRepositoryAssociation_S3BucketProperty `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

