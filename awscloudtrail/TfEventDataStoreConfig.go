package awscloudtrail

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEventDataStoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#name TfEventDataStore#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// advanced_event_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#advanced_event_selector TfEventDataStore#advanced_event_selector}
	// Experimental.
	AdvancedEventSelector interface{} `field:"optional" json:"advancedEventSelector" yaml:"advancedEventSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#billing_mode TfEventDataStore#billing_mode}.
	// Experimental.
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#id TfEventDataStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#kms_key_id TfEventDataStore#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#multi_region_enabled TfEventDataStore#multi_region_enabled}.
	// Experimental.
	MultiRegionEnabled interface{} `field:"optional" json:"multiRegionEnabled" yaml:"multiRegionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#organization_enabled TfEventDataStore#organization_enabled}.
	// Experimental.
	OrganizationEnabled interface{} `field:"optional" json:"organizationEnabled" yaml:"organizationEnabled"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#region TfEventDataStore#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#retention_period TfEventDataStore#retention_period}.
	// Experimental.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#suspend TfEventDataStore#suspend}.
	// Experimental.
	Suspend *string `field:"optional" json:"suspend" yaml:"suspend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#tags TfEventDataStore#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#tags_all TfEventDataStore#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#termination_protection_enabled TfEventDataStore#termination_protection_enabled}.
	// Experimental.
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#timeouts TfEventDataStore#timeouts}
	// Experimental.
	Timeouts *TfEventDataStore_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

