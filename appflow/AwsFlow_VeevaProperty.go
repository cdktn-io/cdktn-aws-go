package appflow


// Experimental.
type AwsFlow_VeevaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object AwsFlow#object}.
	// Experimental.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#document_type AwsFlow#document_type}.
	// Experimental.
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#include_all_versions AwsFlow#include_all_versions}.
	// Experimental.
	IncludeAllVersions interface{} `field:"optional" json:"includeAllVersions" yaml:"includeAllVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#include_renditions AwsFlow#include_renditions}.
	// Experimental.
	IncludeRenditions interface{} `field:"optional" json:"includeRenditions" yaml:"includeRenditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#include_source_files AwsFlow#include_source_files}.
	// Experimental.
	IncludeSourceFiles interface{} `field:"optional" json:"includeSourceFiles" yaml:"includeSourceFiles"`
}

