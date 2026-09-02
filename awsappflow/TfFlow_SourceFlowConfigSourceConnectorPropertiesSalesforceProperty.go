package awsappflow


// Experimental.
type TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object TfFlow#object}.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#data_transfer_api TfFlow#data_transfer_api}.
	// Experimental.
	DataTransferApi *string `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#enable_dynamic_field_update TfFlow#enable_dynamic_field_update}.
	// Experimental.
	EnableDynamicFieldUpdate interface{} `field:"optional" json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#include_deleted_records TfFlow#include_deleted_records}.
	// Experimental.
	IncludeDeletedRecords interface{} `field:"optional" json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

