package eks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eks/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/eks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster}.
// Experimental.
type AwsCluster interface {
	cdktn.TerraformResource
	// Experimental.
	AccessConfig() AwsCluster_AccessConfigPropertyOutputReference
	// Experimental.
	AccessConfigInput() *AwsCluster_AccessConfigProperty
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
	CertificateAuthority() AwsCluster_CertificateAuthorityPropertyList
	// Experimental.
	ClusterId() *string
	// Experimental.
	ComputeConfig() AwsCluster_ComputeConfigPropertyOutputReference
	// Experimental.
	ComputeConfigInput() *AwsCluster_ComputeConfigProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ControlPlaneScalingConfig() AwsCluster_ControlPlaneScalingConfigPropertyOutputReference
	// Experimental.
	ControlPlaneScalingConfigInput() *AwsCluster_ControlPlaneScalingConfigProperty
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
	EncryptionConfig() AwsCluster_EncryptionConfigPropertyOutputReference
	// Experimental.
	EncryptionConfigInput() *AwsCluster_EncryptionConfigProperty
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
	Identity() AwsCluster_IdentityPropertyList
	// Experimental.
	IdInput() *string
	// Experimental.
	KubeApiServerConfig() AwsCluster_KubeApiServerConfigPropertyOutputReference
	// Experimental.
	KubeApiServerConfigInput() *AwsCluster_KubeApiServerConfigProperty
	// Experimental.
	KubeControllerManagerConfig() AwsCluster_KubeControllerManagerConfigPropertyOutputReference
	// Experimental.
	KubeControllerManagerConfigInput() *AwsCluster_KubeControllerManagerConfigProperty
	// Experimental.
	KubernetesNetworkConfig() AwsCluster_KubernetesNetworkConfigPropertyOutputReference
	// Experimental.
	KubernetesNetworkConfigInput() *AwsCluster_KubernetesNetworkConfigProperty
	// Experimental.
	KubeSchedulerConfig() AwsCluster_KubeSchedulerConfigPropertyOutputReference
	// Experimental.
	KubeSchedulerConfigInput() *AwsCluster_KubeSchedulerConfigProperty
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
	OutpostConfig() AwsCluster_OutpostConfigPropertyOutputReference
	// Experimental.
	OutpostConfigInput() *AwsCluster_OutpostConfigProperty
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
	RemoteNetworkConfig() AwsCluster_RemoteNetworkConfigPropertyOutputReference
	// Experimental.
	RemoteNetworkConfigInput() *AwsCluster_RemoteNetworkConfigProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	StorageConfig() AwsCluster_StorageConfigPropertyOutputReference
	// Experimental.
	StorageConfigInput() *AwsCluster_StorageConfigProperty
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
	Timeouts() AwsCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UpgradePolicy() AwsCluster_UpgradePolicyPropertyOutputReference
	// Experimental.
	UpgradePolicyInput() *AwsCluster_UpgradePolicyProperty
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
	// Experimental.
	VpcConfig() AwsCluster_VpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *AwsCluster_VpcConfigProperty
	// Experimental.
	ZonalShiftConfig() AwsCluster_ZonalShiftConfigPropertyOutputReference
	// Experimental.
	ZonalShiftConfigInput() *AwsCluster_ZonalShiftConfigProperty
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
	PutAccessConfig(value *AwsCluster_AccessConfigProperty)
	// Experimental.
	PutComputeConfig(value *AwsCluster_ComputeConfigProperty)
	// Experimental.
	PutControlPlaneScalingConfig(value *AwsCluster_ControlPlaneScalingConfigProperty)
	// Experimental.
	PutEncryptionConfig(value *AwsCluster_EncryptionConfigProperty)
	// Experimental.
	PutKubeApiServerConfig(value *AwsCluster_KubeApiServerConfigProperty)
	// Experimental.
	PutKubeControllerManagerConfig(value *AwsCluster_KubeControllerManagerConfigProperty)
	// Experimental.
	PutKubernetesNetworkConfig(value *AwsCluster_KubernetesNetworkConfigProperty)
	// Experimental.
	PutKubeSchedulerConfig(value *AwsCluster_KubeSchedulerConfigProperty)
	// Experimental.
	PutOutpostConfig(value *AwsCluster_OutpostConfigProperty)
	// Experimental.
	PutRemoteNetworkConfig(value *AwsCluster_RemoteNetworkConfigProperty)
	// Experimental.
	PutStorageConfig(value *AwsCluster_StorageConfigProperty)
	// Experimental.
	PutTimeouts(value *AwsCluster_TimeoutsProperty)
	// Experimental.
	PutUpgradePolicy(value *AwsCluster_UpgradePolicyProperty)
	// Experimental.
	PutVpcConfig(value *AwsCluster_VpcConfigProperty)
	// Experimental.
	PutZonalShiftConfig(value *AwsCluster_ZonalShiftConfigProperty)
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

// The jsii proxy struct for AwsCluster
type jsiiProxy_AwsCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCluster) AccessConfig() AwsCluster_AccessConfigPropertyOutputReference {
	var returns AwsCluster_AccessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) AccessConfigInput() *AwsCluster_AccessConfigProperty {
	var returns *AwsCluster_AccessConfigProperty
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) BootstrapSelfManagedAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) BootstrapSelfManagedAddonsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) CertificateAuthority() AwsCluster_CertificateAuthorityPropertyList {
	var returns AwsCluster_CertificateAuthorityPropertyList
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ComputeConfig() AwsCluster_ComputeConfigPropertyOutputReference {
	var returns AwsCluster_ComputeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"computeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ComputeConfigInput() *AwsCluster_ComputeConfigProperty {
	var returns *AwsCluster_ComputeConfigProperty
	_jsii_.Get(
		j,
		"computeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ControlPlaneScalingConfig() AwsCluster_ControlPlaneScalingConfigPropertyOutputReference {
	var returns AwsCluster_ControlPlaneScalingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"controlPlaneScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ControlPlaneScalingConfigInput() *AwsCluster_ControlPlaneScalingConfigProperty {
	var returns *AwsCluster_ControlPlaneScalingConfigProperty
	_jsii_.Get(
		j,
		"controlPlaneScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) EnabledClusterLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) EnabledClusterLogTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) EncryptionConfig() AwsCluster_EncryptionConfigPropertyOutputReference {
	var returns AwsCluster_EncryptionConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) EncryptionConfigInput() *AwsCluster_EncryptionConfigProperty {
	var returns *AwsCluster_EncryptionConfigProperty
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ForceUpdateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ForceUpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Identity() AwsCluster_IdentityPropertyList {
	var returns AwsCluster_IdentityPropertyList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubeApiServerConfig() AwsCluster_KubeApiServerConfigPropertyOutputReference {
	var returns AwsCluster_KubeApiServerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeApiServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubeApiServerConfigInput() *AwsCluster_KubeApiServerConfigProperty {
	var returns *AwsCluster_KubeApiServerConfigProperty
	_jsii_.Get(
		j,
		"kubeApiServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubeControllerManagerConfig() AwsCluster_KubeControllerManagerConfigPropertyOutputReference {
	var returns AwsCluster_KubeControllerManagerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeControllerManagerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubeControllerManagerConfigInput() *AwsCluster_KubeControllerManagerConfigProperty {
	var returns *AwsCluster_KubeControllerManagerConfigProperty
	_jsii_.Get(
		j,
		"kubeControllerManagerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubernetesNetworkConfig() AwsCluster_KubernetesNetworkConfigPropertyOutputReference {
	var returns AwsCluster_KubernetesNetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubernetesNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubernetesNetworkConfigInput() *AwsCluster_KubernetesNetworkConfigProperty {
	var returns *AwsCluster_KubernetesNetworkConfigProperty
	_jsii_.Get(
		j,
		"kubernetesNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubeSchedulerConfig() AwsCluster_KubeSchedulerConfigPropertyOutputReference {
	var returns AwsCluster_KubeSchedulerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeSchedulerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KubeSchedulerConfigInput() *AwsCluster_KubeSchedulerConfigProperty {
	var returns *AwsCluster_KubeSchedulerConfigProperty
	_jsii_.Get(
		j,
		"kubeSchedulerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) OutpostConfig() AwsCluster_OutpostConfigPropertyOutputReference {
	var returns AwsCluster_OutpostConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"outpostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) OutpostConfigInput() *AwsCluster_OutpostConfigProperty {
	var returns *AwsCluster_OutpostConfigProperty
	_jsii_.Get(
		j,
		"outpostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) RemoteNetworkConfig() AwsCluster_RemoteNetworkConfigPropertyOutputReference {
	var returns AwsCluster_RemoteNetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"remoteNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) RemoteNetworkConfigInput() *AwsCluster_RemoteNetworkConfigProperty {
	var returns *AwsCluster_RemoteNetworkConfigProperty
	_jsii_.Get(
		j,
		"remoteNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) StorageConfig() AwsCluster_StorageConfigPropertyOutputReference {
	var returns AwsCluster_StorageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) StorageConfigInput() *AwsCluster_StorageConfigProperty {
	var returns *AwsCluster_StorageConfigProperty
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Timeouts() AwsCluster_TimeoutsPropertyOutputReference {
	var returns AwsCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) UpgradePolicy() AwsCluster_UpgradePolicyPropertyOutputReference {
	var returns AwsCluster_UpgradePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) UpgradePolicyInput() *AwsCluster_UpgradePolicyProperty {
	var returns *AwsCluster_UpgradePolicyProperty
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) VpcConfig() AwsCluster_VpcConfigPropertyOutputReference {
	var returns AwsCluster_VpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) VpcConfigInput() *AwsCluster_VpcConfigProperty {
	var returns *AwsCluster_VpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ZonalShiftConfig() AwsCluster_ZonalShiftConfigPropertyOutputReference {
	var returns AwsCluster_ZonalShiftConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"zonalShiftConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ZonalShiftConfigInput() *AwsCluster_ZonalShiftConfigProperty {
	var returns *AwsCluster_ZonalShiftConfigProperty
	_jsii_.Get(
		j,
		"zonalShiftConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
// Experimental.
func NewAwsCluster(scope constructs.Construct, id *string, config *AwsClusterConfig) AwsCluster {
	_init_.Initialize()

	if err := validateNewAwsClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCluster{}

	_jsii_.Create(
		"@cdktn/aws-eks.AwsCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
// Experimental.
func NewAwsCluster_Override(a AwsCluster, scope constructs.Construct, id *string, config *AwsClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.AwsCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCluster)SetBootstrapSelfManagedAddons(val interface{}) {
	if err := j.validateSetBootstrapSelfManagedAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapSelfManagedAddons",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetEnabledClusterLogTypes(val *[]*string) {
	if err := j.validateSetEnabledClusterLogTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledClusterLogTypes",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetForceUpdateVersion(val interface{}) {
	if err := j.validateSetForceUpdateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateVersion",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a AwsCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsCluster",
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
func AwsCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eks.AwsCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCluster) PutAccessConfig(value *AwsCluster_AccessConfigProperty) {
	if err := a.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutComputeConfig(value *AwsCluster_ComputeConfigProperty) {
	if err := a.validatePutComputeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComputeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutControlPlaneScalingConfig(value *AwsCluster_ControlPlaneScalingConfigProperty) {
	if err := a.validatePutControlPlaneScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putControlPlaneScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutEncryptionConfig(value *AwsCluster_EncryptionConfigProperty) {
	if err := a.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutKubeApiServerConfig(value *AwsCluster_KubeApiServerConfigProperty) {
	if err := a.validatePutKubeApiServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubeApiServerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutKubeControllerManagerConfig(value *AwsCluster_KubeControllerManagerConfigProperty) {
	if err := a.validatePutKubeControllerManagerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubeControllerManagerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutKubernetesNetworkConfig(value *AwsCluster_KubernetesNetworkConfigProperty) {
	if err := a.validatePutKubernetesNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubernetesNetworkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutKubeSchedulerConfig(value *AwsCluster_KubeSchedulerConfigProperty) {
	if err := a.validatePutKubeSchedulerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubeSchedulerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutOutpostConfig(value *AwsCluster_OutpostConfigProperty) {
	if err := a.validatePutOutpostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutpostConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutRemoteNetworkConfig(value *AwsCluster_RemoteNetworkConfigProperty) {
	if err := a.validatePutRemoteNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRemoteNetworkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutStorageConfig(value *AwsCluster_StorageConfigProperty) {
	if err := a.validatePutStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutTimeouts(value *AwsCluster_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutUpgradePolicy(value *AwsCluster_UpgradePolicyProperty) {
	if err := a.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutVpcConfig(value *AwsCluster_VpcConfigProperty) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutZonalShiftConfig(value *AwsCluster_ZonalShiftConfigProperty) {
	if err := a.validatePutZonalShiftConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZonalShiftConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCluster) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetBootstrapSelfManagedAddons() {
	_jsii_.InvokeVoid(
		a,
		"resetBootstrapSelfManagedAddons",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetComputeConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetControlPlaneScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetControlPlaneScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetEnabledClusterLogTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledClusterLogTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetForceUpdateVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetForceUpdateVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetKubeApiServerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubeApiServerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetKubeControllerManagerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubeControllerManagerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetKubernetesNetworkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubernetesNetworkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetKubeSchedulerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKubeSchedulerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetOutpostConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOutpostConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetRemoteNetworkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteNetworkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetStorageConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetZonalShiftConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetZonalShiftConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

