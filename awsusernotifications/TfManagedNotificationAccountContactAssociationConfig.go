package awsusernotifications

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfManagedNotificationAccountContactAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/notifications_managed_notification_account_contact_association#contact_identifier TfManagedNotificationAccountContactAssociation#contact_identifier}.
	// Experimental.
	ContactIdentifier *string `field:"required" json:"contactIdentifier" yaml:"contactIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/notifications_managed_notification_account_contact_association#managed_notification_configuration_arn TfManagedNotificationAccountContactAssociation#managed_notification_configuration_arn}.
	// Experimental.
	ManagedNotificationConfigurationArn *string `field:"required" json:"managedNotificationConfigurationArn" yaml:"managedNotificationConfigurationArn"`
}

