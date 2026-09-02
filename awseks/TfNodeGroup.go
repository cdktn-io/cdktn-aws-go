package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group aws_eks_node_group}.
// Experimental.
type TfNodeGroup interface {
	cdktn.TerraformResource
	// Experimental.
	AmiType() *string
	// Experimental.
	SetAmiType(val *string)
	// Experimental.
	AmiTypeInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	CapacityType() *string
	// Experimental.
	SetCapacityType(val *string)
	// Experimental.
	CapacityTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterName() *string
	// Experimental.
	SetClusterName(val *string)
	// Experimental.
	ClusterNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DiskSize() *float64
	// Experimental.
	SetDiskSize(val *float64)
	// Experimental.
	DiskSizeInput() *float64
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
	IdInput() *string
	// Experimental.
	InstanceTypes() *[]*string
	// Experimental.
	SetInstanceTypes(val *[]*string)
	// Experimental.
	InstanceTypesInput() *[]*string
	// Experimental.
	Labels() *map[string]*string
	// Experimental.
	SetLabels(val *map[string]*string)
	// Experimental.
	LabelsInput() *map[string]*string
	// Experimental.
	LaunchTemplate() TfNodeGroup_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *TfNodeGroup_LaunchTemplateProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeGroupName() *string
	// Experimental.
	SetNodeGroupName(val *string)
	// Experimental.
	NodeGroupNameInput() *string
	// Experimental.
	NodeGroupNamePrefix() *string
	// Experimental.
	SetNodeGroupNamePrefix(val *string)
	// Experimental.
	NodeGroupNamePrefixInput() *string
	// Experimental.
	NodeRepairConfig() TfNodeGroup_NodeRepairConfigPropertyOutputReference
	// Experimental.
	NodeRepairConfigInput() *TfNodeGroup_NodeRepairConfigProperty
	// Experimental.
	NodeRoleArn() *string
	// Experimental.
	SetNodeRoleArn(val *string)
	// Experimental.
	NodeRoleArnInput() *string
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
	ReleaseVersion() *string
	// Experimental.
	SetReleaseVersion(val *string)
	// Experimental.
	ReleaseVersionInput() *string
	// Experimental.
	RemoteAccess() TfNodeGroup_RemoteAccessPropertyOutputReference
	// Experimental.
	RemoteAccessInput() *TfNodeGroup_RemoteAccessProperty
	// Experimental.
	Resources() TfNodeGroup_ResourcesPropertyList
	// Experimental.
	ScalingConfig() TfNodeGroup_ScalingConfigPropertyOutputReference
	// Experimental.
	ScalingConfigInput() *TfNodeGroup_ScalingConfigProperty
	// Experimental.
	Status() *string
	// Experimental.
	SubnetIds() *[]*string
	// Experimental.
	SetSubnetIds(val *[]*string)
	// Experimental.
	SubnetIdsInput() *[]*string
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
	Taint() TfNodeGroup_TaintPropertyList
	// Experimental.
	TaintInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfNodeGroup_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UpdateConfig() TfNodeGroup_UpdateConfigPropertyOutputReference
	// Experimental.
	UpdateConfigInput() *TfNodeGroup_UpdateConfigProperty
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
	// Experimental.
	WarmPoolConfig() TfNodeGroup_WarmPoolConfigPropertyOutputReference
	// Experimental.
	WarmPoolConfigInput() *TfNodeGroup_WarmPoolConfigProperty
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
	PutLaunchTemplate(value *TfNodeGroup_LaunchTemplateProperty)
	// Experimental.
	PutNodeRepairConfig(value *TfNodeGroup_NodeRepairConfigProperty)
	// Experimental.
	PutRemoteAccess(value *TfNodeGroup_RemoteAccessProperty)
	// Experimental.
	PutScalingConfig(value *TfNodeGroup_ScalingConfigProperty)
	// Experimental.
	PutTaint(value interface{})
	// Experimental.
	PutTimeouts(value *TfNodeGroup_TimeoutsProperty)
	// Experimental.
	PutUpdateConfig(value *TfNodeGroup_UpdateConfigProperty)
	// Experimental.
	PutWarmPoolConfig(value *TfNodeGroup_WarmPoolConfigProperty)
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
	ResetAmiType()
	// Experimental.
	ResetCapacityType()
	// Experimental.
	ResetDiskSize()
	// Experimental.
	ResetForceUpdateVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInstanceTypes()
	// Experimental.
	ResetLabels()
	// Experimental.
	ResetLaunchTemplate()
	// Experimental.
	ResetNodeGroupName()
	// Experimental.
	ResetNodeGroupNamePrefix()
	// Experimental.
	ResetNodeRepairConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReleaseVersion()
	// Experimental.
	ResetRemoteAccess()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTaint()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetUpdateConfig()
	// Experimental.
	ResetVersion()
	// Experimental.
	ResetWarmPoolConfig()
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

// The jsii proxy struct for TfNodeGroup
type jsiiProxy_TfNodeGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfNodeGroup) AmiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) AmiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) CapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) CapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ForceUpdateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ForceUpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) LaunchTemplate() TfNodeGroup_LaunchTemplatePropertyOutputReference {
	var returns TfNodeGroup_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) LaunchTemplateInput() *TfNodeGroup_LaunchTemplateProperty {
	var returns *TfNodeGroup_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeGroupNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeGroupNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeRepairConfig() TfNodeGroup_NodeRepairConfigPropertyOutputReference {
	var returns TfNodeGroup_NodeRepairConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"nodeRepairConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeRepairConfigInput() *TfNodeGroup_NodeRepairConfigProperty {
	var returns *TfNodeGroup_NodeRepairConfigProperty
	_jsii_.Get(
		j,
		"nodeRepairConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) NodeRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ReleaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) RemoteAccess() TfNodeGroup_RemoteAccessPropertyOutputReference {
	var returns TfNodeGroup_RemoteAccessPropertyOutputReference
	_jsii_.Get(
		j,
		"remoteAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) RemoteAccessInput() *TfNodeGroup_RemoteAccessProperty {
	var returns *TfNodeGroup_RemoteAccessProperty
	_jsii_.Get(
		j,
		"remoteAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Resources() TfNodeGroup_ResourcesPropertyList {
	var returns TfNodeGroup_ResourcesPropertyList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ScalingConfig() TfNodeGroup_ScalingConfigPropertyOutputReference {
	var returns TfNodeGroup_ScalingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) ScalingConfigInput() *TfNodeGroup_ScalingConfigProperty {
	var returns *TfNodeGroup_ScalingConfigProperty
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Taint() TfNodeGroup_TaintPropertyList {
	var returns TfNodeGroup_TaintPropertyList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Timeouts() TfNodeGroup_TimeoutsPropertyOutputReference {
	var returns TfNodeGroup_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) UpdateConfig() TfNodeGroup_UpdateConfigPropertyOutputReference {
	var returns TfNodeGroup_UpdateConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"updateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) UpdateConfigInput() *TfNodeGroup_UpdateConfigProperty {
	var returns *TfNodeGroup_UpdateConfigProperty
	_jsii_.Get(
		j,
		"updateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) WarmPoolConfig() TfNodeGroup_WarmPoolConfigPropertyOutputReference {
	var returns TfNodeGroup_WarmPoolConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"warmPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup) WarmPoolConfigInput() *TfNodeGroup_WarmPoolConfigProperty {
	var returns *TfNodeGroup_WarmPoolConfigProperty
	_jsii_.Get(
		j,
		"warmPoolConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group aws_eks_node_group} Resource.
// Experimental.
func NewTfNodeGroup(scope constructs.Construct, id *string, config *TfNodeGroupConfig) TfNodeGroup {
	_init_.Initialize()

	if err := validateNewTfNodeGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfNodeGroup{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group aws_eks_node_group} Resource.
// Experimental.
func NewTfNodeGroup_Override(t TfNodeGroup, scope constructs.Construct, id *string, config *TfNodeGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetAmiType(val *string) {
	if err := j.validateSetAmiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiType",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetCapacityType(val *string) {
	if err := j.validateSetCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityType",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetDiskSize(val *float64) {
	if err := j.validateSetDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetForceUpdateVersion(val interface{}) {
	if err := j.validateSetForceUpdateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateVersion",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetNodeGroupName(val *string) {
	if err := j.validateSetNodeGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupName",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetNodeGroupNamePrefix(val *string) {
	if err := j.validateSetNodeGroupNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupNamePrefix",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetNodeRoleArn(val *string) {
	if err := j.validateSetNodeRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetReleaseVersion(val *string) {
	if err := j.validateSetReleaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseVersion",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a TfNodeGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfNodeGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfNodeGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfNodeGroup",
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
func TfNodeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfNodeGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfNodeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfNodeGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfNodeGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfNodeGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfNodeGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfNodeGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.TfNodeGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfNodeGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eks.TfNodeGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfNodeGroup) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfNodeGroup) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfNodeGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfNodeGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfNodeGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfNodeGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfNodeGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfNodeGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfNodeGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfNodeGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfNodeGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfNodeGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfNodeGroup) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfNodeGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfNodeGroup) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfNodeGroup) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutLaunchTemplate(value *TfNodeGroup_LaunchTemplateProperty) {
	if err := t.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutNodeRepairConfig(value *TfNodeGroup_NodeRepairConfigProperty) {
	if err := t.validatePutNodeRepairConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNodeRepairConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutRemoteAccess(value *TfNodeGroup_RemoteAccessProperty) {
	if err := t.validatePutRemoteAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRemoteAccess",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutScalingConfig(value *TfNodeGroup_ScalingConfigProperty) {
	if err := t.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutTaint(value interface{}) {
	if err := t.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTaint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutTimeouts(value *TfNodeGroup_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutUpdateConfig(value *TfNodeGroup_UpdateConfigProperty) {
	if err := t.validatePutUpdateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUpdateConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) PutWarmPoolConfig(value *TfNodeGroup_WarmPoolConfigProperty) {
	if err := t.validatePutWarmPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWarmPoolConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetAmiType() {
	_jsii_.InvokeVoid(
		t,
		"resetAmiType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetCapacityType() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetDiskSize() {
	_jsii_.InvokeVoid(
		t,
		"resetDiskSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetForceUpdateVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetForceUpdateVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetLabels() {
	_jsii_.InvokeVoid(
		t,
		"resetLabels",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetNodeGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetNodeGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetNodeGroupNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNodeGroupNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetNodeRepairConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetNodeRepairConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetReleaseVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetReleaseVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetRemoteAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetRemoteAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetTaint() {
	_jsii_.InvokeVoid(
		t,
		"resetTaint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetUpdateConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetUpdateConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) ResetWarmPoolConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetWarmPoolConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

