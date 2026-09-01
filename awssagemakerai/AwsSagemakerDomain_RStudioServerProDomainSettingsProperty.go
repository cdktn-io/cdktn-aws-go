package awssagemakerai


// Experimental.
type AwsSagemakerDomain_RStudioServerProDomainSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#domain_execution_role_arn AwsSagemakerDomain#domain_execution_role_arn}.
	// Experimental.
	DomainExecutionRoleArn *string `field:"required" json:"domainExecutionRoleArn" yaml:"domainExecutionRoleArn"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_resource_spec AwsSagemakerDomain#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *AwsSagemakerDomain_DomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecProperty `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#r_studio_connect_url AwsSagemakerDomain#r_studio_connect_url}.
	// Experimental.
	RStudioConnectUrl *string `field:"optional" json:"rStudioConnectUrl" yaml:"rStudioConnectUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#r_studio_package_manager_url AwsSagemakerDomain#r_studio_package_manager_url}.
	// Experimental.
	RStudioPackageManagerUrl *string `field:"optional" json:"rStudioPackageManagerUrl" yaml:"rStudioPackageManagerUrl"`
}

