package awsworkspaces


// Experimental.
type TfWorkspace_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#create TfWorkspace#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#delete TfWorkspace#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#update TfWorkspace#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

