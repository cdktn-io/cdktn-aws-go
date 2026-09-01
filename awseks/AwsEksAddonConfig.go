package awseks

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEksAddonConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#addon_name AwsEksAddon#addon_name}.
	// Experimental.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#cluster_name AwsEksAddon#cluster_name}.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#addon_version AwsEksAddon#addon_version}.
	// Experimental.
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#configuration_values AwsEksAddon#configuration_values}.
	// Experimental.
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#id AwsEksAddon#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// namespace_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#namespace_config AwsEksAddon#namespace_config}
	// Experimental.
	NamespaceConfig *AwsEksAddon_NamespaceConfigProperty `field:"optional" json:"namespaceConfig" yaml:"namespaceConfig"`
	// pod_identity_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#pod_identity_association AwsEksAddon#pod_identity_association}
	// Experimental.
	PodIdentityAssociation interface{} `field:"optional" json:"podIdentityAssociation" yaml:"podIdentityAssociation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#preserve AwsEksAddon#preserve}.
	// Experimental.
	Preserve interface{} `field:"optional" json:"preserve" yaml:"preserve"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#region AwsEksAddon#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#resolve_conflicts_on_create AwsEksAddon#resolve_conflicts_on_create}.
	// Experimental.
	ResolveConflictsOnCreate *string `field:"optional" json:"resolveConflictsOnCreate" yaml:"resolveConflictsOnCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#resolve_conflicts_on_update AwsEksAddon#resolve_conflicts_on_update}.
	// Experimental.
	ResolveConflictsOnUpdate *string `field:"optional" json:"resolveConflictsOnUpdate" yaml:"resolveConflictsOnUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#service_account_role_arn AwsEksAddon#service_account_role_arn}.
	// Experimental.
	ServiceAccountRoleArn *string `field:"optional" json:"serviceAccountRoleArn" yaml:"serviceAccountRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#tags AwsEksAddon#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#tags_all AwsEksAddon#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_addon#timeouts AwsEksAddon#timeouts}
	// Experimental.
	Timeouts *AwsEksAddon_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

