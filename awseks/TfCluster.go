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
type TfCluster interface {
	cdktn.TerraformResource
	// Experimental.
	AccessConfig() TfCluster_AccessConfigPropertyOutputReference
	// Experimental.
	AccessConfigInput() *TfCluster_AccessConfigProperty
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
	CertificateAuthority() TfCluster_CertificateAuthorityPropertyList
	// Experimental.
	ClusterId() *string
	// Experimental.
	ComputeConfig() TfCluster_ComputeConfigPropertyOutputReference
	// Experimental.
	ComputeConfigInput() *TfCluster_ComputeConfigProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ControlPlaneScalingConfig() TfCluster_ControlPlaneScalingConfigPropertyOutputReference
	// Experimental.
	ControlPlaneScalingConfigInput() *TfCluster_ControlPlaneScalingConfigProperty
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
	EncryptionConfig() TfCluster_EncryptionConfigPropertyOutputReference
	// Experimental.
	EncryptionConfigInput() *TfCluster_EncryptionConfigProperty
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
	Identity() TfCluster_IdentityPropertyList
	// Experimental.
	IdInput() *string
	// Experimental.
	KubeApiServerConfig() TfCluster_KubeApiServerConfigPropertyOutputReference
	// Experimental.
	KubeApiServerConfigInput() *TfCluster_KubeApiServerConfigProperty
	// Experimental.
	KubeControllerManagerConfig() TfCluster_KubeControllerManagerConfigPropertyOutputReference
	// Experimental.
	KubeControllerManagerConfigInput() *TfCluster_KubeControllerManagerConfigProperty
	// Experimental.
	KubernetesNetworkConfig() TfCluster_KubernetesNetworkConfigPropertyOutputReference
	// Experimental.
	KubernetesNetworkConfigInput() *TfCluster_KubernetesNetworkConfigProperty
	// Experimental.
	KubeSchedulerConfig() TfCluster_KubeSchedulerConfigPropertyOutputReference
	// Experimental.
	KubeSchedulerConfigInput() *TfCluster_KubeSchedulerConfigProperty
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
	OutpostConfig() TfCluster_OutpostConfigPropertyOutputReference
	// Experimental.
	OutpostConfigInput() *TfCluster_OutpostConfigProperty
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
	RemoteNetworkConfig() TfCluster_RemoteNetworkConfigPropertyOutputReference
	// Experimental.
	RemoteNetworkConfigInput() *TfCluster_RemoteNetworkConfigProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	StorageConfig() TfCluster_StorageConfigPropertyOutputReference
	// Experimental.
	StorageConfigInput() *TfCluster_StorageConfigProperty
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
	Timeouts() TfCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UpgradePolicy() TfCluster_UpgradePolicyPropertyOutputReference
	// Experimental.
	UpgradePolicyInput() *TfCluster_UpgradePolicyProperty
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
	// Experimental.
	VpcConfig() TfCluster_VpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *TfCluster_VpcConfigProperty
	// Experimental.
	ZonalShiftConfig() TfCluster_ZonalShiftConfigPropertyOutputReference
	// Experimental.
	ZonalShiftConfigInput() *TfCluster_ZonalShiftConfigProperty
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
	PutAccessConfig(value *TfCluster_AccessConfigProperty)
	// Experimental.
	PutComputeConfig(value *TfCluster_ComputeConfigProperty)
	// Experimental.
	PutControlPlaneScalingConfig(value *TfCluster_ControlPlaneScalingConfigProperty)
	// Experimental.
	PutEncryptionConfig(value *TfCluster_EncryptionConfigProperty)
	// Experimental.
	PutKubeApiServerConfig(value *TfCluster_KubeApiServerConfigProperty)
	// Experimental.
	PutKubeControllerManagerConfig(value *TfCluster_KubeControllerManagerConfigProperty)
	// Experimental.
	PutKubernetesNetworkConfig(value *TfCluster_KubernetesNetworkConfigProperty)
	// Experimental.
	PutKubeSchedulerConfig(value *TfCluster_KubeSchedulerConfigProperty)
	// Experimental.
	PutOutpostConfig(value *TfCluster_OutpostConfigProperty)
	// Experimental.
	PutRemoteNetworkConfig(value *TfCluster_RemoteNetworkConfigProperty)
	// Experimental.
	PutStorageConfig(value *TfCluster_StorageConfigProperty)
	// Experimental.
	PutTimeouts(value *TfCluster_TimeoutsProperty)
	// Experimental.
	PutUpgradePolicy(value *TfCluster_UpgradePolicyProperty)
	// Experimental.
	PutVpcConfig(value *TfCluster_VpcConfigProperty)
	// Experimental.
	PutZonalShiftConfig(value *TfCluster_ZonalShiftConfigProperty)
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

// The jsii proxy struct for TfCluster
type jsiiProxy_TfCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCluster) AccessConfig() TfCluster_AccessConfigPropertyOutputReference {
	var returns TfCluster_AccessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AccessConfigInput() *TfCluster_AccessConfigProperty {
	var returns *TfCluster_AccessConfigProperty
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapSelfManagedAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapSelfManagedAddonsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CertificateAuthority() TfCluster_CertificateAuthorityPropertyList {
	var returns TfCluster_CertificateAuthorityPropertyList
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ComputeConfig() TfCluster_ComputeConfigPropertyOutputReference {
	var returns TfCluster_ComputeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"computeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ComputeConfigInput() *TfCluster_ComputeConfigProperty {
	var returns *TfCluster_ComputeConfigProperty
	_jsii_.Get(
		j,
		"computeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ControlPlaneScalingConfig() TfCluster_ControlPlaneScalingConfigPropertyOutputReference {
	var returns TfCluster_ControlPlaneScalingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"controlPlaneScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ControlPlaneScalingConfigInput() *TfCluster_ControlPlaneScalingConfigProperty {
	var returns *TfCluster_ControlPlaneScalingConfigProperty
	_jsii_.Get(
		j,
		"controlPlaneScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnabledClusterLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnabledClusterLogTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledClusterLogTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EncryptionConfig() TfCluster_EncryptionConfigPropertyOutputReference {
	var returns TfCluster_EncryptionConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EncryptionConfigInput() *TfCluster_EncryptionConfigProperty {
	var returns *TfCluster_EncryptionConfigProperty
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ForceUpdateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ForceUpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Identity() TfCluster_IdentityPropertyList {
	var returns TfCluster_IdentityPropertyList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubeApiServerConfig() TfCluster_KubeApiServerConfigPropertyOutputReference {
	var returns TfCluster_KubeApiServerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeApiServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubeApiServerConfigInput() *TfCluster_KubeApiServerConfigProperty {
	var returns *TfCluster_KubeApiServerConfigProperty
	_jsii_.Get(
		j,
		"kubeApiServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubeControllerManagerConfig() TfCluster_KubeControllerManagerConfigPropertyOutputReference {
	var returns TfCluster_KubeControllerManagerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeControllerManagerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubeControllerManagerConfigInput() *TfCluster_KubeControllerManagerConfigProperty {
	var returns *TfCluster_KubeControllerManagerConfigProperty
	_jsii_.Get(
		j,
		"kubeControllerManagerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubernetesNetworkConfig() TfCluster_KubernetesNetworkConfigPropertyOutputReference {
	var returns TfCluster_KubernetesNetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubernetesNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubernetesNetworkConfigInput() *TfCluster_KubernetesNetworkConfigProperty {
	var returns *TfCluster_KubernetesNetworkConfigProperty
	_jsii_.Get(
		j,
		"kubernetesNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubeSchedulerConfig() TfCluster_KubeSchedulerConfigPropertyOutputReference {
	var returns TfCluster_KubeSchedulerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kubeSchedulerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KubeSchedulerConfigInput() *TfCluster_KubeSchedulerConfigProperty {
	var returns *TfCluster_KubeSchedulerConfigProperty
	_jsii_.Get(
		j,
		"kubeSchedulerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) OutpostConfig() TfCluster_OutpostConfigPropertyOutputReference {
	var returns TfCluster_OutpostConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"outpostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) OutpostConfigInput() *TfCluster_OutpostConfigProperty {
	var returns *TfCluster_OutpostConfigProperty
	_jsii_.Get(
		j,
		"outpostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RemoteNetworkConfig() TfCluster_RemoteNetworkConfigPropertyOutputReference {
	var returns TfCluster_RemoteNetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"remoteNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RemoteNetworkConfigInput() *TfCluster_RemoteNetworkConfigProperty {
	var returns *TfCluster_RemoteNetworkConfigProperty
	_jsii_.Get(
		j,
		"remoteNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StorageConfig() TfCluster_StorageConfigPropertyOutputReference {
	var returns TfCluster_StorageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StorageConfigInput() *TfCluster_StorageConfigProperty {
	var returns *TfCluster_StorageConfigProperty
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Timeouts() TfCluster_TimeoutsPropertyOutputReference {
	var returns TfCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) UpgradePolicy() TfCluster_UpgradePolicyPropertyOutputReference {
	var returns TfCluster_UpgradePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) UpgradePolicyInput() *TfCluster_UpgradePolicyProperty {
	var returns *TfCluster_UpgradePolicyProperty
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) VpcConfig() TfCluster_VpcConfigPropertyOutputReference {
	var returns TfCluster_VpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) VpcConfigInput() *TfCluster_VpcConfigProperty {
	var returns *TfCluster_VpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ZonalShiftConfig() TfCluster_ZonalShiftConfigPropertyOutputReference {
	var returns TfCluster_ZonalShiftConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"zonalShiftConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ZonalShiftConfigInput() *TfCluster_ZonalShiftConfigProperty {
	var returns *TfCluster_ZonalShiftConfigProperty
	_jsii_.Get(
		j,
		"zonalShiftConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
// Experimental.
func NewTfCluster(scope constructs.Construct, id *string, config *TfClusterConfig) TfCluster {
	_init_.Initialize()

	if err := validateNewTfClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster aws_eks_cluster} Resource.
// Experimental.
func NewTfCluster_Override(t TfCluster, scope constructs.Construct, id *string, config *TfClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCluster)SetBootstrapSelfManagedAddons(val interface{}) {
	if err := j.validateSetBootstrapSelfManagedAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapSelfManagedAddons",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEnabledClusterLogTypes(val *[]*string) {
	if err := j.validateSetEnabledClusterLogTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledClusterLogTypes",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetForceUpdateVersion(val interface{}) {
	if err := j.validateSetForceUpdateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateVersion",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a TfCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfCluster",
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
func TfCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eks.TfCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCluster) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCluster) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCluster) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCluster) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCluster) PutAccessConfig(value *TfCluster_AccessConfigProperty) {
	if err := t.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutComputeConfig(value *TfCluster_ComputeConfigProperty) {
	if err := t.validatePutComputeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComputeConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutControlPlaneScalingConfig(value *TfCluster_ControlPlaneScalingConfigProperty) {
	if err := t.validatePutControlPlaneScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putControlPlaneScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutEncryptionConfig(value *TfCluster_EncryptionConfigProperty) {
	if err := t.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutKubeApiServerConfig(value *TfCluster_KubeApiServerConfigProperty) {
	if err := t.validatePutKubeApiServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKubeApiServerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutKubeControllerManagerConfig(value *TfCluster_KubeControllerManagerConfigProperty) {
	if err := t.validatePutKubeControllerManagerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKubeControllerManagerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutKubernetesNetworkConfig(value *TfCluster_KubernetesNetworkConfigProperty) {
	if err := t.validatePutKubernetesNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKubernetesNetworkConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutKubeSchedulerConfig(value *TfCluster_KubeSchedulerConfigProperty) {
	if err := t.validatePutKubeSchedulerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKubeSchedulerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutOutpostConfig(value *TfCluster_OutpostConfigProperty) {
	if err := t.validatePutOutpostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutpostConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutRemoteNetworkConfig(value *TfCluster_RemoteNetworkConfigProperty) {
	if err := t.validatePutRemoteNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRemoteNetworkConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutStorageConfig(value *TfCluster_StorageConfigProperty) {
	if err := t.validatePutStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutTimeouts(value *TfCluster_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutUpgradePolicy(value *TfCluster_UpgradePolicyProperty) {
	if err := t.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutVpcConfig(value *TfCluster_VpcConfigProperty) {
	if err := t.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutZonalShiftConfig(value *TfCluster_ZonalShiftConfigProperty) {
	if err := t.validatePutZonalShiftConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putZonalShiftConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCluster) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetBootstrapSelfManagedAddons() {
	_jsii_.InvokeVoid(
		t,
		"resetBootstrapSelfManagedAddons",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetComputeConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetComputeConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetControlPlaneScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetControlPlaneScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEnabledClusterLogTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabledClusterLogTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetForceUpdateVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetForceUpdateVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetKubeApiServerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetKubeApiServerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetKubeControllerManagerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetKubeControllerManagerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetKubernetesNetworkConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetKubernetesNetworkConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetKubeSchedulerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetKubeSchedulerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetOutpostConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOutpostConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetRemoteNetworkConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRemoteNetworkConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetStorageConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetZonalShiftConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetZonalShiftConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

