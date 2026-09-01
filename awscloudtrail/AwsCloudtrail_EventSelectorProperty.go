package awscloudtrail


// Experimental.
type AwsCloudtrail_EventSelectorProperty struct {
	// data_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#data_resource AwsCloudtrail#data_resource}
	// Experimental.
	DataResource interface{} `field:"optional" json:"dataResource" yaml:"dataResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#exclude_management_event_sources AwsCloudtrail#exclude_management_event_sources}.
	// Experimental.
	ExcludeManagementEventSources *[]*string `field:"optional" json:"excludeManagementEventSources" yaml:"excludeManagementEventSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#include_management_events AwsCloudtrail#include_management_events}.
	// Experimental.
	IncludeManagementEvents interface{} `field:"optional" json:"includeManagementEvents" yaml:"includeManagementEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#read_write_type AwsCloudtrail#read_write_type}.
	// Experimental.
	ReadWriteType *string `field:"optional" json:"readWriteType" yaml:"readWriteType"`
}

