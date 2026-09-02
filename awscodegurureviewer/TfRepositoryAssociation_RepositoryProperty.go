package awscodegurureviewer


// Experimental.
type TfRepositoryAssociation_RepositoryProperty struct {
	// bitbucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#bitbucket TfRepositoryAssociation#bitbucket}
	// Experimental.
	Bitbucket *TfRepositoryAssociation_BitbucketProperty `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// codecommit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#codecommit TfRepositoryAssociation#codecommit}
	// Experimental.
	Codecommit *TfRepositoryAssociation_CodecommitProperty `field:"optional" json:"codecommit" yaml:"codecommit"`
	// github_enterprise_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#github_enterprise_server TfRepositoryAssociation#github_enterprise_server}
	// Experimental.
	GithubEnterpriseServer *TfRepositoryAssociation_GithubEnterpriseServerProperty `field:"optional" json:"githubEnterpriseServer" yaml:"githubEnterpriseServer"`
	// s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#s3_bucket TfRepositoryAssociation#s3_bucket}
	// Experimental.
	S3Bucket *TfRepositoryAssociation_S3BucketProperty `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

