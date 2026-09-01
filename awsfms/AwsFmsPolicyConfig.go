package awsfms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFmsPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#exclude_resource_tags AwsFmsPolicy#exclude_resource_tags}.
	// Experimental.
	ExcludeResourceTags interface{} `field:"required" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#name AwsFmsPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// security_service_policy_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#security_service_policy_data AwsFmsPolicy#security_service_policy_data}
	// Experimental.
	SecurityServicePolicyData *AwsFmsPolicy_SecurityServicePolicyDataProperty `field:"required" json:"securityServicePolicyData" yaml:"securityServicePolicyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#delete_all_policy_resources AwsFmsPolicy#delete_all_policy_resources}.
	// Experimental.
	DeleteAllPolicyResources interface{} `field:"optional" json:"deleteAllPolicyResources" yaml:"deleteAllPolicyResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#delete_unused_fm_managed_resources AwsFmsPolicy#delete_unused_fm_managed_resources}.
	// Experimental.
	DeleteUnusedFmManagedResources interface{} `field:"optional" json:"deleteUnusedFmManagedResources" yaml:"deleteUnusedFmManagedResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#description AwsFmsPolicy#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// exclude_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#exclude_map AwsFmsPolicy#exclude_map}
	// Experimental.
	ExcludeMap *AwsFmsPolicy_ExcludeMapProperty `field:"optional" json:"excludeMap" yaml:"excludeMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#id AwsFmsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// include_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#include_map AwsFmsPolicy#include_map}
	// Experimental.
	IncludeMap *AwsFmsPolicy_IncludeMapProperty `field:"optional" json:"includeMap" yaml:"includeMap"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#region AwsFmsPolicy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#remediation_enabled AwsFmsPolicy#remediation_enabled}.
	// Experimental.
	RemediationEnabled interface{} `field:"optional" json:"remediationEnabled" yaml:"remediationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#resource_set_ids AwsFmsPolicy#resource_set_ids}.
	// Experimental.
	ResourceSetIds *[]*string `field:"optional" json:"resourceSetIds" yaml:"resourceSetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#resource_tag_logical_operator AwsFmsPolicy#resource_tag_logical_operator}.
	// Experimental.
	ResourceTagLogicalOperator *string `field:"optional" json:"resourceTagLogicalOperator" yaml:"resourceTagLogicalOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#resource_tags AwsFmsPolicy#resource_tags}.
	// Experimental.
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#resource_type AwsFmsPolicy#resource_type}.
	// Experimental.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#resource_type_list AwsFmsPolicy#resource_type_list}.
	// Experimental.
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#tags AwsFmsPolicy#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#tags_all AwsFmsPolicy#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

