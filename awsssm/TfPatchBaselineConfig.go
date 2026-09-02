package awsssm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPatchBaselineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#name TfPatchBaseline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// approval_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approval_rule TfPatchBaseline#approval_rule}
	// Experimental.
	ApprovalRule interface{} `field:"optional" json:"approvalRule" yaml:"approvalRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approved_patches TfPatchBaseline#approved_patches}.
	// Experimental.
	ApprovedPatches *[]*string `field:"optional" json:"approvedPatches" yaml:"approvedPatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approved_patches_compliance_level TfPatchBaseline#approved_patches_compliance_level}.
	// Experimental.
	ApprovedPatchesComplianceLevel *string `field:"optional" json:"approvedPatchesComplianceLevel" yaml:"approvedPatchesComplianceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approved_patches_enable_non_security TfPatchBaseline#approved_patches_enable_non_security}.
	// Experimental.
	ApprovedPatchesEnableNonSecurity interface{} `field:"optional" json:"approvedPatchesEnableNonSecurity" yaml:"approvedPatchesEnableNonSecurity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#available_security_updates_compliance_status TfPatchBaseline#available_security_updates_compliance_status}.
	// Experimental.
	AvailableSecurityUpdatesComplianceStatus *string `field:"optional" json:"availableSecurityUpdatesComplianceStatus" yaml:"availableSecurityUpdatesComplianceStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#description TfPatchBaseline#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// global_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#global_filter TfPatchBaseline#global_filter}
	// Experimental.
	GlobalFilter interface{} `field:"optional" json:"globalFilter" yaml:"globalFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#id TfPatchBaseline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#operating_system TfPatchBaseline#operating_system}.
	// Experimental.
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#region TfPatchBaseline#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#rejected_patches TfPatchBaseline#rejected_patches}.
	// Experimental.
	RejectedPatches *[]*string `field:"optional" json:"rejectedPatches" yaml:"rejectedPatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#rejected_patches_action TfPatchBaseline#rejected_patches_action}.
	// Experimental.
	RejectedPatchesAction *string `field:"optional" json:"rejectedPatchesAction" yaml:"rejectedPatchesAction"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#source TfPatchBaseline#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#tags TfPatchBaseline#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#tags_all TfPatchBaseline#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

