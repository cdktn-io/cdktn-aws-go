package workspaces


// Experimental.
type AwsDirectory_SamlPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#relay_state_parameter_name AwsDirectory#relay_state_parameter_name}.
	// Experimental.
	RelayStateParameterName *string `field:"optional" json:"relayStateParameterName" yaml:"relayStateParameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#status AwsDirectory#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#user_access_url AwsDirectory#user_access_url}.
	// Experimental.
	UserAccessUrl *string `field:"optional" json:"userAccessUrl" yaml:"userAccessUrl"`
}

