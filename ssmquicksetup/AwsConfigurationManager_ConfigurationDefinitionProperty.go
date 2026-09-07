package ssmquicksetup


// Experimental.
type AwsConfigurationManager_ConfigurationDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmquicksetup_configuration_manager#parameters AwsConfigurationManager#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmquicksetup_configuration_manager#type AwsConfigurationManager#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmquicksetup_configuration_manager#local_deployment_administration_role_arn AwsConfigurationManager#local_deployment_administration_role_arn}.
	// Experimental.
	LocalDeploymentAdministrationRoleArn *string `field:"optional" json:"localDeploymentAdministrationRoleArn" yaml:"localDeploymentAdministrationRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmquicksetup_configuration_manager#local_deployment_execution_role_name AwsConfigurationManager#local_deployment_execution_role_name}.
	// Experimental.
	LocalDeploymentExecutionRoleName *string `field:"optional" json:"localDeploymentExecutionRoleName" yaml:"localDeploymentExecutionRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmquicksetup_configuration_manager#type_version AwsConfigurationManager#type_version}.
	// Experimental.
	TypeVersion *string `field:"optional" json:"typeVersion" yaml:"typeVersion"`
}

