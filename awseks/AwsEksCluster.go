package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster}.
// Experimental.
type AwsEksCluster interface {
	cdktn.TerraformResource
	// Experimental.
	AccessConfig() AwsEksCluster_AccessConfigPropertyOutputReference
	// Experimental.
	AccessConfigInput() *AwsEksCluster_AccessConfigProperty
	// Experimental.
	Arn() *string
	// Experimental.
	BootstrapSelfManagedAddons() interface{}
	// Experimental.
	SetBootstrapSelfManagedAddons(val interface{})
	// Experimental.
	BootstrapSelfManagedAddonsInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CertificateAuthority() AwsEksCluster_CertificateAuthorityPropertyList
	// Experimental.
	ClusterId() *string
	// Experimental.
	ComputeConfig() AwsEksCluster_ComputeConfigPropertyOutputReference
	// Experimental.
	ComputeConfigInput() *AwsEksCluster_ComputeConfigProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ControlPlaneScalingConfig() AwsEksCluster_ControlPlaneScalingConfigPropertyOutputReference
	// Experimental.
	ControlPlaneScalingConfigInput() *AwsEksCluster_ControlPlaneScalingConfigProperty
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CreatedAt() *string
	// Experimental.
	DeletionProtection() interface{}
	// Experimental.
	SetDeletionProtection(val interface{})
	// Experimental.
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EnabledClusterLogTypes() *[]*string
	// Experimental.
	SetEnabledClusterLogTypes(val *[]*string)
	// Experimental.
	EnabledClusterLogTypesInput() *[]*string
	// Experimental.
	EncryptionConfig() AwsEksCluster_EncryptionConfigPropertyOutputReference
	// Experimental.
	EncryptionConfigInput() *AwsEksCluster_EncryptionConfigProperty
	// Experimental.
	Endpoint() *string
	// Experimental.
	ForceUpdateVersion() interface{}
	// Experimental.
	SetForceUpdateVersion(val interface{})
	// Experimental.
	ForceUpdateVersionInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	Identity() AwsEksCluster_IdentityPropertyList
	// Experimental.
	IdInput() *string
	// Experimental.
	KubeApiServerConfig() AwsEksCluster_KubeApiServerConfigPropertyOutputReference
	// Experimental.
	KubeApiServerConfigInput() *AwsEksCluster_KubeApiServerConfigProperty
	// Experimental.
	KubeControllerManagerConfig() AwsEksCluster_KubeControllerManagerConfigPropertyOutputReference
	// Experimental.
	KubeControllerManagerConfigInput() *AwsEksCluster_KubeControllerManagerConfigProperty
	// Experimental.
	KubernetesNetworkConfig() AwsEksCluster_KubernetesNetworkConfigPropertyOutputReference
	// Experimental.
	KubernetesNetworkConfigInput() *AwsEksCluster_KubernetesNetworkConfigProperty
	// Experimental.
	KubeSchedulerConfig() AwsEksCluster_KubeSchedulerConfigPropertyOutputReference
	// Experimental.
	KubeSchedulerConfigInput() *AwsEksCluster_KubeSchedulerConfigProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutpostConfig() AwsEksCluster_OutpostConfigPropertyOutputReference
	// Experimental.
	OutpostConfigInput() *AwsEksCluster_OutpostConfigProperty
	// Experimental.
	PlatformVersion() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RemoteNetworkConfig() AwsEksCluster_RemoteNetworkConfigPropertyOutputReference
	// Experimental.
	RemoteNetworkConfigInput() *AwsEksCluster_RemoteNetworkConfigProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	StorageConfig() AwsEksCluster_StorageConfigPropertyOutputReference
	// Experimental.
	StorageConfigInput() *AwsEksCluster_StorageConfigProperty
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() *map[string]*string
	// Experimental.
	SetTagsAll(val *map[string]*string)
	// Experimental.
	TagsAllInput() *map[string]*string
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsEksCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UpgradePolicy() AwsEksCluster_UpgradePolicyPropertyOutputReference
	// Experimental.
	UpgradePolicyInput() *AwsEksCluster_UpgradePolicyProperty
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
	// Experimental.
	VpcConfig() AwsEksCluster_VpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *AwsEksCluster_VpcConfigProperty
	// Experimental.
	ZonalShiftConfig() AwsEksCluster_ZonalShiftConfigPropertyOutputReference
	// Experimental.
	ZonalShiftConfigInput() *AwsEksCluster_ZonalShiftConfigProperty
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutAccessConfig(value *AwsEksCluster_AccessConfigProperty)
	// Experimental.
	PutComputeConfig(value *AwsEksCluster_ComputeConfigProperty)
	// Experimental.
	PutControlPlaneScalingConfig(value *AwsEksCluster_ControlPlaneScalingConfigProperty)
	// Experimental.
	PutEncryptionConfig(value *AwsEksCluster_EncryptionConfigProperty)
	// Experimental.
	PutKubeApiServerConfig(value *AwsEksCluster_KubeApiServerConfigProperty)
	// Experimental.
	PutKubeControllerManagerConfig(value *AwsEksCluster_KubeControllerManagerConfigProperty)
	// Experimental.
	PutKubernetesNetworkConfig(value *AwsEksCluster_KubernetesNetworkConfigProperty)
	// Experimental.
	PutKubeSchedulerConfig(value *AwsEksCluster_KubeSchedulerConfigProperty)
	// Experimental.
	PutOutpostConfig(value *AwsEksCluster_OutpostConfigProperty)
	// Experimental.
	PutRemoteNetworkConfig(value *AwsEksCluster_RemoteNetworkConfigProperty)
	// Experimental.
	PutStorageConfig(value *AwsEksCluster_StorageConfigProperty)
	// Experimental.
	PutTimeouts(value *AwsEksCluster_TimeoutsProperty)
	// Experimental.
	PutUpgradePolicy(value *AwsEksCluster_UpgradePolicyProperty)
	// Experimental.
	PutVpcConfig(value *AwsEksCluster_VpcConfigProperty)
	// Experimental.
	PutZonalShiftConfig(value *AwsEksCluster_ZonalShiftConfigProperty)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	// Experimental.
	ResetAccessConfig()
	// Experimental.
	ResetBootstrapSelfManagedAddons()
	// Experimental.
	ResetComputeConfig()
	// Experimental.
	ResetControlPlaneScalingConfig()
	// Experimental.
	ResetDeletionProtection()
	// Experimental.
	ResetEnabledClusterLogTypes()
	// Experimental.
	ResetEncryptionConfig()
	// Experimental.
	ResetForceUpdateVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKubeApiServerConfig()
	// Experimental.
	ResetKubeControllerManagerConfig()
	// Experimental.
	ResetKubernetesNetworkConfig()
	// Experimental.
	ResetKubeSchedulerConfig()
	// Experimental.
	ResetOutpostConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRemoteNetworkConfig()
	// Experimental.
	ResetStorageConfig()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetUpgradePolicy()
	// Experimental.
	ResetVersion()
	// Experimental.
	ResetZonalShiftConfig()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AwsEksCluster
type jsiiProxy_AwsEksCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsEksCluster) AccessConfig() AwsEksCluster_AccessConfigPropertyOutputReference {
	var returns AwsEksCluster_AccessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) AccessConfigInput() *AwsEksCluster_AccessConfigProperty {
	var returns *AwsEksCluster_AccessConfigProperty
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) BootstrapSelfManagedAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) BootstrapSelfManagedAddonsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) CertificateAuthority() AwsEksCluster_CertificateAuthorityPropertyList {
	var returns AwsEksCluster_CertificateAuthorityPropertyList
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ComputeConfig() AwsEksCluster_ComputeConfigPropertyOutputReference {
	var returns AwsEksCluster_ComputeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"computeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ComputeConfigInput() *AwsEksCluster_ComputeConfigProperty {
	var returns *AwsEksCluster_ComputeConfigProperty
	_jsii_.Get(
		j,
		"computeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ControlPlaneScalingConfig() AwsEksCluster_ControlPlaneScalingConfigPropertyOutputReference {
	var returns AwsEksCluster_ControlPlaneScalingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"controlPlaneScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ControlPlaneScalingConfigInput() *AwsEksCluster_ControlPlaneScalingConfigProperty {
	var returns *AwsEksCluster_ControlPlaneScalingConfigProperty
	_jsii_.Get(
		j,
		"controlPlaneScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) EnabledClusterLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) EnabledClusterLogTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) EncryptionConfig() AwsEksCluster_EncryptionConfigPropertyOutputReference {
	var returns AwsEksCluster_EncryptionConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) EncryptionConfigInput() *AwsEksCluster_EncryptionConfigProperty {
	var returns *AwsEksCluster_EncryptionConfigProperty
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ForceUpdateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ForceUpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Identity() AwsEksCluster_IdentityPropertyList {
	var returns AwsEksCluster_IdentityPropertyList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubeApiServerConfig() AwsEksCluster_KubeApiServerConfigPropertyOutputReference {
	var returns AwsEksCluster_KubeApiServerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeApiServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubeApiServerConfigInput() *AwsEksCluster_KubeApiServerConfigProperty {
	var returns *AwsEksCluster_KubeApiServerConfigProperty
	_jsii_.Get(
		j,
		"kubeApiServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubeControllerManagerConfig() AwsEksCluster_KubeControllerManagerConfigPropertyOutputReference {
	var returns AwsEksCluster_KubeControllerManagerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeControllerManagerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubeControllerManagerConfigInput() *AwsEksCluster_KubeControllerManagerConfigProperty {
	var returns *AwsEksCluster_KubeControllerManagerConfigProperty
	_jsii_.Get(
		j,
		"kubeControllerManagerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubernetesNetworkConfig() AwsEksCluster_KubernetesNetworkConfigPropertyOutputReference {
	var returns AwsEksCluster_KubernetesNetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubernetesNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubernetesNetworkConfigInput() *AwsEksCluster_KubernetesNetworkConfigProperty {
	var returns *AwsEksCluster_KubernetesNetworkConfigProperty
	_jsii_.Get(
		j,
		"kubernetesNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubeSchedulerConfig() AwsEksCluster_KubeSchedulerConfigPropertyOutputReference {
	var returns AwsEksCluster_KubeSchedulerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeSchedulerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) KubeSchedulerConfigInput() *AwsEksCluster_KubeSchedulerConfigProperty {
	var returns *AwsEksCluster_KubeSchedulerConfigProperty
	_jsii_.Get(
		j,
		"kubeSchedulerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) OutpostConfig() AwsEksCluster_OutpostConfigPropertyOutputReference {
	var returns AwsEksCluster_OutpostConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"outpostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) OutpostConfigInput() *AwsEksCluster_OutpostConfigProperty {
	var returns *AwsEksCluster_OutpostConfigProperty
	_jsii_.Get(
		j,
		"outpostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) RemoteNetworkConfig() AwsEksCluster_RemoteNetworkConfigPropertyOutputReference {
	var returns AwsEksCluster_RemoteNetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"remoteNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) RemoteNetworkConfigInput() *AwsEksCluster_RemoteNetworkConfigProperty {
	var returns *AwsEksCluster_RemoteNetworkConfigProperty
	_jsii_.Get(
		j,
		"remoteNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) StorageConfig() AwsEksCluster_StorageConfigPropertyOutputReference {
	var returns AwsEksCluster_StorageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) StorageConfigInput() *AwsEksCluster_StorageConfigProperty {
	var returns *AwsEksCluster_StorageConfigProperty
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Timeouts() AwsEksCluster_TimeoutsPropertyOutputReference {
	var returns AwsEksCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) UpgradePolicy() AwsEksCluster_UpgradePolicyPropertyOutputReference {
	var returns AwsEksCluster_UpgradePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) UpgradePolicyInput() *AwsEksCluster_UpgradePolicyProperty {
	var returns *AwsEksCluster_UpgradePolicyProperty
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) VpcConfig() AwsEksCluster_VpcConfigPropertyOutputReference {
	var returns AwsEksCluster_VpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) VpcConfigInput() *AwsEksCluster_VpcConfigProperty {
	var returns *AwsEksCluster_VpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ZonalShiftConfig() AwsEksCluster_ZonalShiftConfigPropertyOutputReference {
	var returns AwsEksCluster_ZonalShiftConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"zonalShiftConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster) ZonalShiftConfigInput() *AwsEksCluster_ZonalShiftConfigProperty {
	var returns *AwsEksCluster_ZonalShiftConfigProperty
	_jsii_.Get(
		j,
		"zonalShiftConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
// Experimental.
func NewAwsEksCluster(scope constructs.Construct, id *string, config *AwsEksClusterConfig) AwsEksCluster {
	_init_.Initialize()

	if err := validateNewAwsEksClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEksCluster{}

	_jsii_.Create(
		"@cdktn/aws-eks.AwsEksCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
// Experimental.
func NewAwsEksCluster_Override(a AwsEksCluster, scope constructs.Construct, id *string, config *AwsEksClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.AwsEksCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetBootstrapSelfManagedAddons(val interface{}) {
	if err := j.validateSetBootstrapSelfManagedAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapSelfManagedAddons",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetEnabledClusterLogTypes(val *[]*string) {
	if err := j.validateSetEnabledClusterLogTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledClusterLogTypes",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetForceUpdateVersion(val interface{}) {
	if err := j.validateSetForceUpdateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateVersion",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a AwsEksCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsEksCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsEksCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AwsEksCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEksCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEksCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEksCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEksCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEksCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEksCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eks.AwsEksCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEksCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsEksCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEksCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsEksCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEksCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsEksCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEksCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutAccessConfig(value *AwsEksCluster_AccessConfigProperty) {
	if err := a.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutComputeConfig(value *AwsEksCluster_ComputeConfigProperty) {
	if err := a.validatePutComputeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComputeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutControlPlaneScalingConfig(value *AwsEksCluster_ControlPlaneScalingConfigProperty) {
	if err := a.validatePutControlPlaneScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putControlPlaneScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutEncryptionConfig(value *AwsEksCluster_EncryptionConfigProperty) {
	if err := a.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutKubeApiServerConfig(value *AwsEksCluster_KubeApiServerConfigProperty) {
	if err := a.validatePutKubeApiServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubeApiServerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutKubeControllerManagerConfig(value *AwsEksCluster_KubeControllerManagerConfigProperty) {
	if err := a.validatePutKubeControllerManagerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubeControllerManagerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutKubernetesNetworkConfig(value *AwsEksCluster_KubernetesNetworkConfigProperty) {
	if err := a.validatePutKubernetesNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubernetesNetworkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutKubeSchedulerConfig(value *AwsEksCluster_KubeSchedulerConfigProperty) {
	if err := a.validatePutKubeSchedulerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubeSchedulerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutOutpostConfig(value *AwsEksCluster_OutpostConfigProperty) {
	if err := a.validatePutOutpostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutpostConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutRemoteNetworkConfig(value *AwsEksCluster_RemoteNetworkConfigProperty) {
	if err := a.validatePutRemoteNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRemoteNetworkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutStorageConfig(value *AwsEksCluster_StorageConfigProperty) {
	if err := a.validatePutStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutTimeouts(value *AwsEksCluster_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutUpgradePolicy(value *AwsEksCluster_UpgradePolicyProperty) {
	if err := a.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutVpcConfig(value *AwsEksCluster_VpcConfigProperty) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) PutZonalShiftConfig(value *AwsEksCluster_ZonalShiftConfigProperty) {
	if err := a.validatePutZonalShiftConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZonalShiftConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetBootstrapSelfManagedAddons() {
	_jsii_.InvokeVoid(
		a,
		"resetBootstrapSelfManagedAddons",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetComputeConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetControlPlaneScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetControlPlaneScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetEnabledClusterLogTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledClusterLogTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetForceUpdateVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetForceUpdateVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetKubeApiServerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubeApiServerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetKubeControllerManagerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubeControllerManagerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetKubernetesNetworkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubernetesNetworkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetKubeSchedulerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubeSchedulerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetOutpostConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOutpostConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetRemoteNetworkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteNetworkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetStorageConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) ResetZonalShiftConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetZonalShiftConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

