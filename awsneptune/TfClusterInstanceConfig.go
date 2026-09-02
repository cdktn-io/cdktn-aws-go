package awsneptune

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClusterInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#cluster_identifier TfClusterInstance#cluster_identifier}.
	// Experimental.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#instance_class TfClusterInstance#instance_class}.
	// Experimental.
	InstanceClass *string `field:"required" json:"instanceClass" yaml:"instanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#apply_immediately TfClusterInstance#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#auto_minor_version_upgrade TfClusterInstance#auto_minor_version_upgrade}.
	// Experimental.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#availability_zone TfClusterInstance#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#engine TfClusterInstance#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#engine_version TfClusterInstance#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#id TfClusterInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#identifier TfClusterInstance#identifier}.
	// Experimental.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#identifier_prefix TfClusterInstance#identifier_prefix}.
	// Experimental.
	IdentifierPrefix *string `field:"optional" json:"identifierPrefix" yaml:"identifierPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#neptune_parameter_group_name TfClusterInstance#neptune_parameter_group_name}.
	// Experimental.
	NeptuneParameterGroupName *string `field:"optional" json:"neptuneParameterGroupName" yaml:"neptuneParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#neptune_subnet_group_name TfClusterInstance#neptune_subnet_group_name}.
	// Experimental.
	NeptuneSubnetGroupName *string `field:"optional" json:"neptuneSubnetGroupName" yaml:"neptuneSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#port TfClusterInstance#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#preferred_backup_window TfClusterInstance#preferred_backup_window}.
	// Experimental.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#preferred_maintenance_window TfClusterInstance#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#promotion_tier TfClusterInstance#promotion_tier}.
	// Experimental.
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#publicly_accessible TfClusterInstance#publicly_accessible}.
	// Experimental.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#region TfClusterInstance#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#skip_final_snapshot TfClusterInstance#skip_final_snapshot}.
	// Experimental.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#tags TfClusterInstance#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#tags_all TfClusterInstance#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_instance#timeouts TfClusterInstance#timeouts}
	// Experimental.
	Timeouts *TfClusterInstance_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

