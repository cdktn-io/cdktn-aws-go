package awsdatazone

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEnvironmentBlueprintConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#domain_id TfEnvironmentBlueprintConfiguration#domain_id}.
	// Experimental.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#enabled_regions TfEnvironmentBlueprintConfiguration#enabled_regions}.
	// Experimental.
	EnabledRegions *[]*string `field:"required" json:"enabledRegions" yaml:"enabledRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#environment_blueprint_id TfEnvironmentBlueprintConfiguration#environment_blueprint_id}.
	// Experimental.
	EnvironmentBlueprintId *string `field:"required" json:"environmentBlueprintId" yaml:"environmentBlueprintId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#global_parameters TfEnvironmentBlueprintConfiguration#global_parameters}.
	// Experimental.
	GlobalParameters *map[string]*string `field:"optional" json:"globalParameters" yaml:"globalParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#manage_access_role_arn TfEnvironmentBlueprintConfiguration#manage_access_role_arn}.
	// Experimental.
	ManageAccessRoleArn *string `field:"optional" json:"manageAccessRoleArn" yaml:"manageAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#provisioning_role_arn TfEnvironmentBlueprintConfiguration#provisioning_role_arn}.
	// Experimental.
	ProvisioningRoleArn *string `field:"optional" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#region TfEnvironmentBlueprintConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datazone_environment_blueprint_configuration#regional_parameters TfEnvironmentBlueprintConfiguration#regional_parameters}.
	// Experimental.
	RegionalParameters interface{} `field:"optional" json:"regionalParameters" yaml:"regionalParameters"`
}

