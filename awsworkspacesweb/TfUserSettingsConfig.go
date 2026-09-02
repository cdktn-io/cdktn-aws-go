package awsworkspacesweb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#copy_allowed TfUserSettings#copy_allowed}.
	// Experimental.
	CopyAllowed *string `field:"required" json:"copyAllowed" yaml:"copyAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#download_allowed TfUserSettings#download_allowed}.
	// Experimental.
	DownloadAllowed *string `field:"required" json:"downloadAllowed" yaml:"downloadAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#paste_allowed TfUserSettings#paste_allowed}.
	// Experimental.
	PasteAllowed *string `field:"required" json:"pasteAllowed" yaml:"pasteAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#print_allowed TfUserSettings#print_allowed}.
	// Experimental.
	PrintAllowed *string `field:"required" json:"printAllowed" yaml:"printAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#upload_allowed TfUserSettings#upload_allowed}.
	// Experimental.
	UploadAllowed *string `field:"required" json:"uploadAllowed" yaml:"uploadAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#additional_encryption_context TfUserSettings#additional_encryption_context}.
	// Experimental.
	AdditionalEncryptionContext *map[string]*string `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// cookie_synchronization_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#cookie_synchronization_configuration TfUserSettings#cookie_synchronization_configuration}
	// Experimental.
	CookieSynchronizationConfiguration interface{} `field:"optional" json:"cookieSynchronizationConfiguration" yaml:"cookieSynchronizationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#customer_managed_key TfUserSettings#customer_managed_key}.
	// Experimental.
	CustomerManagedKey *string `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#deep_link_allowed TfUserSettings#deep_link_allowed}.
	// Experimental.
	DeepLinkAllowed *string `field:"optional" json:"deepLinkAllowed" yaml:"deepLinkAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#disconnect_timeout_in_minutes TfUserSettings#disconnect_timeout_in_minutes}.
	// Experimental.
	DisconnectTimeoutInMinutes *float64 `field:"optional" json:"disconnectTimeoutInMinutes" yaml:"disconnectTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#idle_disconnect_timeout_in_minutes TfUserSettings#idle_disconnect_timeout_in_minutes}.
	// Experimental.
	IdleDisconnectTimeoutInMinutes *float64 `field:"optional" json:"idleDisconnectTimeoutInMinutes" yaml:"idleDisconnectTimeoutInMinutes"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#region TfUserSettings#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#tags TfUserSettings#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// toolbar_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#toolbar_configuration TfUserSettings#toolbar_configuration}
	// Experimental.
	ToolbarConfiguration interface{} `field:"optional" json:"toolbarConfiguration" yaml:"toolbarConfiguration"`
}

