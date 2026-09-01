package awsservicequotas

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsServicequotasAutoManagementConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicequotas_auto_management#opt_in_level AwsServicequotasAutoManagement#opt_in_level}.
	// Experimental.
	OptInLevel *string `field:"required" json:"optInLevel" yaml:"optInLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicequotas_auto_management#opt_in_type AwsServicequotasAutoManagement#opt_in_type}.
	// Experimental.
	OptInType *string `field:"required" json:"optInType" yaml:"optInType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicequotas_auto_management#exclusion_list AwsServicequotasAutoManagement#exclusion_list}.
	// Experimental.
	ExclusionList interface{} `field:"optional" json:"exclusionList" yaml:"exclusionList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicequotas_auto_management#notification_arn AwsServicequotasAutoManagement#notification_arn}.
	// Experimental.
	NotificationArn *string `field:"optional" json:"notificationArn" yaml:"notificationArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicequotas_auto_management#region AwsServicequotasAutoManagement#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

