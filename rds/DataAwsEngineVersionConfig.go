package rds

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsEngineVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#engine DataAwsEngineVersion#engine}.
	// Experimental.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#default_only DataAwsEngineVersion#default_only}.
	// Experimental.
	DefaultOnly interface{} `field:"optional" json:"defaultOnly" yaml:"defaultOnly"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#filter DataAwsEngineVersion#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#has_major_target DataAwsEngineVersion#has_major_target}.
	// Experimental.
	HasMajorTarget interface{} `field:"optional" json:"hasMajorTarget" yaml:"hasMajorTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#has_minor_target DataAwsEngineVersion#has_minor_target}.
	// Experimental.
	HasMinorTarget interface{} `field:"optional" json:"hasMinorTarget" yaml:"hasMinorTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#id DataAwsEngineVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#include_all DataAwsEngineVersion#include_all}.
	// Experimental.
	IncludeAll interface{} `field:"optional" json:"includeAll" yaml:"includeAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#latest DataAwsEngineVersion#latest}.
	// Experimental.
	Latest interface{} `field:"optional" json:"latest" yaml:"latest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#parameter_group_family DataAwsEngineVersion#parameter_group_family}.
	// Experimental.
	ParameterGroupFamily *string `field:"optional" json:"parameterGroupFamily" yaml:"parameterGroupFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#preferred_major_targets DataAwsEngineVersion#preferred_major_targets}.
	// Experimental.
	PreferredMajorTargets *[]*string `field:"optional" json:"preferredMajorTargets" yaml:"preferredMajorTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#preferred_upgrade_targets DataAwsEngineVersion#preferred_upgrade_targets}.
	// Experimental.
	PreferredUpgradeTargets *[]*string `field:"optional" json:"preferredUpgradeTargets" yaml:"preferredUpgradeTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#preferred_versions DataAwsEngineVersion#preferred_versions}.
	// Experimental.
	PreferredVersions *[]*string `field:"optional" json:"preferredVersions" yaml:"preferredVersions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#region DataAwsEngineVersion#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#version DataAwsEngineVersion#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

