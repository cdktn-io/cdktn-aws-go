package awscodegurureviewer


// Experimental.
type AwsCodegurureviewerRepositoryAssociation_RepositoryProperty struct {
	// bitbucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#bitbucket AwsCodegurureviewerRepositoryAssociation#bitbucket}
	// Experimental.
	Bitbucket *AwsCodegurureviewerRepositoryAssociation_BitbucketProperty `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// codecommit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#codecommit AwsCodegurureviewerRepositoryAssociation#codecommit}
	// Experimental.
	Codecommit *AwsCodegurureviewerRepositoryAssociation_CodecommitProperty `field:"optional" json:"codecommit" yaml:"codecommit"`
	// github_enterprise_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#github_enterprise_server AwsCodegurureviewerRepositoryAssociation#github_enterprise_server}
	// Experimental.
	GithubEnterpriseServer *AwsCodegurureviewerRepositoryAssociation_GithubEnterpriseServerProperty `field:"optional" json:"githubEnterpriseServer" yaml:"githubEnterpriseServer"`
	// s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#s3_bucket AwsCodegurureviewerRepositoryAssociation#s3_bucket}
	// Experimental.
	S3Bucket *AwsCodegurureviewerRepositoryAssociation_S3BucketProperty `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

