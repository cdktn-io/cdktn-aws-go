package awsmanagedgrafana

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkspaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#account_access_type TfWorkspace#account_access_type}.
	// Experimental.
	AccountAccessType *string `field:"required" json:"accountAccessType" yaml:"accountAccessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#authentication_providers TfWorkspace#authentication_providers}.
	// Experimental.
	AuthenticationProviders *[]*string `field:"required" json:"authenticationProviders" yaml:"authenticationProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#permission_type TfWorkspace#permission_type}.
	// Experimental.
	PermissionType *string `field:"required" json:"permissionType" yaml:"permissionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#configuration TfWorkspace#configuration}.
	// Experimental.
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#data_sources TfWorkspace#data_sources}.
	// Experimental.
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#description TfWorkspace#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#grafana_version TfWorkspace#grafana_version}.
	// Experimental.
	GrafanaVersion *string `field:"optional" json:"grafanaVersion" yaml:"grafanaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#id TfWorkspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#kms_key_id TfWorkspace#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#name TfWorkspace#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// network_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#network_access_control TfWorkspace#network_access_control}
	// Experimental.
	NetworkAccessControl *TfWorkspace_NetworkAccessControlProperty `field:"optional" json:"networkAccessControl" yaml:"networkAccessControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#notification_destinations TfWorkspace#notification_destinations}.
	// Experimental.
	NotificationDestinations *[]*string `field:"optional" json:"notificationDestinations" yaml:"notificationDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#organizational_units TfWorkspace#organizational_units}.
	// Experimental.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#organization_role_name TfWorkspace#organization_role_name}.
	// Experimental.
	OrganizationRoleName *string `field:"optional" json:"organizationRoleName" yaml:"organizationRoleName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#region TfWorkspace#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#role_arn TfWorkspace#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#stack_set_name TfWorkspace#stack_set_name}.
	// Experimental.
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#tags TfWorkspace#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#tags_all TfWorkspace#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#timeouts TfWorkspace#timeouts}
	// Experimental.
	Timeouts *TfWorkspace_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace#vpc_configuration TfWorkspace#vpc_configuration}
	// Experimental.
	VpcConfiguration *TfWorkspace_VpcConfigurationProperty `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

