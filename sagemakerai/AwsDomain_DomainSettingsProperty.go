package sagemakerai


// Experimental.
type AwsDomain_DomainSettingsProperty struct {
	// docker_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#docker_settings AwsDomain#docker_settings}
	// Experimental.
	DockerSettings *AwsDomain_DockerSettingsProperty `field:"optional" json:"dockerSettings" yaml:"dockerSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#execution_role_identity_config AwsDomain#execution_role_identity_config}.
	// Experimental.
	ExecutionRoleIdentityConfig *string `field:"optional" json:"executionRoleIdentityConfig" yaml:"executionRoleIdentityConfig"`
	// r_studio_server_pro_domain_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#r_studio_server_pro_domain_settings AwsDomain#r_studio_server_pro_domain_settings}
	// Experimental.
	RStudioServerProDomainSettings *AwsDomain_RStudioServerProDomainSettingsProperty `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#security_group_ids AwsDomain#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// trusted_identity_propagation_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#trusted_identity_propagation_settings AwsDomain#trusted_identity_propagation_settings}
	// Experimental.
	TrustedIdentityPropagationSettings *AwsDomain_TrustedIdentityPropagationSettingsProperty `field:"optional" json:"trustedIdentityPropagationSettings" yaml:"trustedIdentityPropagationSettings"`
}

