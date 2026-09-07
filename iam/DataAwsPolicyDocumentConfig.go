package iam

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsPolicyDocumentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#id DataAwsPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#override_json DataAwsPolicyDocument#override_json}.
	// Experimental.
	OverrideJson *string `field:"optional" json:"overrideJson" yaml:"overrideJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#override_policy_documents DataAwsPolicyDocument#override_policy_documents}.
	// Experimental.
	OverridePolicyDocuments *[]*string `field:"optional" json:"overridePolicyDocuments" yaml:"overridePolicyDocuments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#policy_id DataAwsPolicyDocument#policy_id}.
	// Experimental.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#source_json DataAwsPolicyDocument#source_json}.
	// Experimental.
	SourceJson *string `field:"optional" json:"sourceJson" yaml:"sourceJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#source_policy_documents DataAwsPolicyDocument#source_policy_documents}.
	// Experimental.
	SourcePolicyDocuments *[]*string `field:"optional" json:"sourcePolicyDocuments" yaml:"sourcePolicyDocuments"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#statement DataAwsPolicyDocument#statement}
	// Experimental.
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#version DataAwsPolicyDocument#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

