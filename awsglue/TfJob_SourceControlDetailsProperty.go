package awsglue


// Experimental.
type TfJob_SourceControlDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#auth_strategy TfJob#auth_strategy}.
	// Experimental.
	AuthStrategy *string `field:"optional" json:"authStrategy" yaml:"authStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#auth_token TfJob#auth_token}.
	// Experimental.
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#branch TfJob#branch}.
	// Experimental.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#folder TfJob#folder}.
	// Experimental.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#last_commit_id TfJob#last_commit_id}.
	// Experimental.
	LastCommitId *string `field:"optional" json:"lastCommitId" yaml:"lastCommitId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#owner TfJob#owner}.
	// Experimental.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#provider TfJob#provider}.
	// Experimental.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job#repository TfJob#repository}.
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

