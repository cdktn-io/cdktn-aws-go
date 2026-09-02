package awscodedeploy


// Experimental.
type TfDeploymentGroup_AutoRollbackConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#enabled TfDeploymentGroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#events TfDeploymentGroup#events}.
	// Experimental.
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
}

