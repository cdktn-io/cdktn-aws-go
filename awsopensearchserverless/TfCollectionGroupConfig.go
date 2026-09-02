package awsopensearchserverless

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCollectionGroupConfig struct {
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
	// Name of the collection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#name TfCollectionGroup#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether standby replicas should be used for collections in this group. One of `ENABLED` or `DISABLED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#standby_replicas TfCollectionGroup#standby_replicas}
	// Experimental.
	StandbyReplicas *string `field:"required" json:"standbyReplicas" yaml:"standbyReplicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#capacity_limits TfCollectionGroup#capacity_limits}.
	// Experimental.
	CapacityLimits interface{} `field:"optional" json:"capacityLimits" yaml:"capacityLimits"`
	// Description of the collection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#description TfCollectionGroup#description}
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Generation of the collection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#generation TfCollectionGroup#generation}
	// Experimental.
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#region TfCollectionGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection_group#tags TfCollectionGroup#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

