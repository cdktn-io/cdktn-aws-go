package emr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emr/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/emr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster aws_emr_cluster}.
// Experimental.
type AwsCluster interface {
	cdktn.TerraformResource
	// Experimental.
	AdditionalInfo() *string
	// Experimental.
	SetAdditionalInfo(val *string)
	// Experimental.
	AdditionalInfoInput() *string
	// Experimental.
	Applications() *[]*string
	// Experimental.
	SetApplications(val *[]*string)
	// Experimental.
	ApplicationsInput() *[]*string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoscalingRole() *string
	// Experimental.
	SetAutoscalingRole(val *string)
	// Experimental.
	AutoscalingRoleInput() *string
	// Experimental.
	AutoTerminationPolicy() AwsCluster_AutoTerminationPolicyPropertyOutputReference
	// Experimental.
	AutoTerminationPolicyInput() *AwsCluster_AutoTerminationPolicyProperty
	// Experimental.
	BootstrapAction() AwsCluster_BootstrapActionPropertyList
	// Experimental.
	BootstrapActionInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterState() *string
	// Experimental.
	Configurations() *string
	// Experimental.
	SetConfigurations(val *string)
	// Experimental.
	ConfigurationsInput() *string
	// Experimental.
	ConfigurationsJson() *string
	// Experimental.
	SetConfigurationsJson(val *string)
	// Experimental.
	ConfigurationsJsonInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CoreInstanceFleet() AwsCluster_CoreInstanceFleetPropertyOutputReference
	// Experimental.
	CoreInstanceFleetInput() *AwsCluster_CoreInstanceFleetProperty
	// Experimental.
	CoreInstanceGroup() AwsCluster_CoreInstanceGroupPropertyOutputReference
	// Experimental.
	CoreInstanceGroupInput() *AwsCluster_CoreInstanceGroupProperty
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomAmiId() *string
	// Experimental.
	SetCustomAmiId(val *string)
	// Experimental.
	CustomAmiIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EbsRootVolumeSize() *float64
	// Experimental.
	SetEbsRootVolumeSize(val *float64)
	// Experimental.
	EbsRootVolumeSizeInput() *float64
	// Experimental.
	Ec2Attributes() AwsCluster_Ec2AttributesPropertyOutputReference
	// Experimental.
	Ec2AttributesInput() *AwsCluster_Ec2AttributesProperty
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
	KeepJobFlowAliveWhenNoSteps() interface{}
	// Experimental.
	SetKeepJobFlowAliveWhenNoSteps(val interface{})
	// Experimental.
	KeepJobFlowAliveWhenNoStepsInput() interface{}
	// Experimental.
	KerberosAttributes() AwsCluster_KerberosAttributesPropertyOutputReference
	// Experimental.
	KerberosAttributesInput() *AwsCluster_KerberosAttributesProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	ListStepsStates() *[]*string
	// Experimental.
	SetListStepsStates(val *[]*string)
	// Experimental.
	ListStepsStatesInput() *[]*string
	// Experimental.
	LogEncryptionKmsKeyId() *string
	// Experimental.
	SetLogEncryptionKmsKeyId(val *string)
	// Experimental.
	LogEncryptionKmsKeyIdInput() *string
	// Experimental.
	LogUri() *string
	// Experimental.
	SetLogUri(val *string)
	// Experimental.
	LogUriInput() *string
	// Experimental.
	MasterInstanceFleet() AwsCluster_MasterInstanceFleetPropertyOutputReference
	// Experimental.
	MasterInstanceFleetInput() *AwsCluster_MasterInstanceFleetProperty
	// Experimental.
	MasterInstanceGroup() AwsCluster_MasterInstanceGroupPropertyOutputReference
	// Experimental.
	MasterInstanceGroupInput() *AwsCluster_MasterInstanceGroupProperty
	// Experimental.
	MasterPublicDns() *string
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
	OsReleaseLabel() *string
	// Experimental.
	SetOsReleaseLabel(val *string)
	// Experimental.
	OsReleaseLabelInput() *string
	// Experimental.
	PlacementGroupConfig() AwsCluster_PlacementGroupConfigPropertyList
	// Experimental.
	PlacementGroupConfigInput() interface{}
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
	ReleaseLabel() *string
	// Experimental.
	SetReleaseLabel(val *string)
	// Experimental.
	ReleaseLabelInput() *string
	// Experimental.
	ScaleDownBehavior() *string
	// Experimental.
	SetScaleDownBehavior(val *string)
	// Experimental.
	ScaleDownBehaviorInput() *string
	// Experimental.
	SecurityConfiguration() *string
	// Experimental.
	SetSecurityConfiguration(val *string)
	// Experimental.
	SecurityConfigurationInput() *string
	// Experimental.
	ServiceRole() *string
	// Experimental.
	SetServiceRole(val *string)
	// Experimental.
	ServiceRoleInput() *string
	// Experimental.
	Step() AwsCluster_StepPropertyList
	// Experimental.
	StepConcurrencyLevel() *float64
	// Experimental.
	SetStepConcurrencyLevel(val *float64)
	// Experimental.
	StepConcurrencyLevelInput() *float64
	// Experimental.
	StepInput() interface{}
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
	TerminationProtection() interface{}
	// Experimental.
	SetTerminationProtection(val interface{})
	// Experimental.
	TerminationProtectionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	UnhealthyNodeReplacement() interface{}
	// Experimental.
	SetUnhealthyNodeReplacement(val interface{})
	// Experimental.
	UnhealthyNodeReplacementInput() interface{}
	// Experimental.
	VisibleToAllUsers() interface{}
	// Experimental.
	SetVisibleToAllUsers(val interface{})
	// Experimental.
	VisibleToAllUsersInput() interface{}
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
	PutAutoTerminationPolicy(value *AwsCluster_AutoTerminationPolicyProperty)
	// Experimental.
	PutBootstrapAction(value interface{})
	// Experimental.
	PutCoreInstanceFleet(value *AwsCluster_CoreInstanceFleetProperty)
	// Experimental.
	PutCoreInstanceGroup(value *AwsCluster_CoreInstanceGroupProperty)
	// Experimental.
	PutEc2Attributes(value *AwsCluster_Ec2AttributesProperty)
	// Experimental.
	PutKerberosAttributes(value *AwsCluster_KerberosAttributesProperty)
	// Experimental.
	PutMasterInstanceFleet(value *AwsCluster_MasterInstanceFleetProperty)
	// Experimental.
	PutMasterInstanceGroup(value *AwsCluster_MasterInstanceGroupProperty)
	// Experimental.
	PutPlacementGroupConfig(value interface{})
	// Experimental.
	PutStep(value interface{})
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
	ResetAdditionalInfo()
	// Experimental.
	ResetApplications()
	// Experimental.
	ResetAutoscalingRole()
	// Experimental.
	ResetAutoTerminationPolicy()
	// Experimental.
	ResetBootstrapAction()
	// Experimental.
	ResetConfigurations()
	// Experimental.
	ResetConfigurationsJson()
	// Experimental.
	ResetCoreInstanceFleet()
	// Experimental.
	ResetCoreInstanceGroup()
	// Experimental.
	ResetCustomAmiId()
	// Experimental.
	ResetEbsRootVolumeSize()
	// Experimental.
	ResetEc2Attributes()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKeepJobFlowAliveWhenNoSteps()
	// Experimental.
	ResetKerberosAttributes()
	// Experimental.
	ResetListStepsStates()
	// Experimental.
	ResetLogEncryptionKmsKeyId()
	// Experimental.
	ResetLogUri()
	// Experimental.
	ResetMasterInstanceFleet()
	// Experimental.
	ResetMasterInstanceGroup()
	// Experimental.
	ResetOsReleaseLabel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlacementGroupConfig()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetScaleDownBehavior()
	// Experimental.
	ResetSecurityConfiguration()
	// Experimental.
	ResetStep()
	// Experimental.
	ResetStepConcurrencyLevel()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTerminationProtection()
	// Experimental.
	ResetUnhealthyNodeReplacement()
	// Experimental.
	ResetVisibleToAllUsers()
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

func (j *jsiiProxy_AwsCluster) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) AdditionalInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Applications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationsInput",
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

func (j *jsiiProxy_AwsCluster) AutoscalingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) AutoscalingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) AutoTerminationPolicy() AwsCluster_AutoTerminationPolicyPropertyOutputReference {
	var returns AwsCluster_AutoTerminationPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"autoTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) AutoTerminationPolicyInput() *AwsCluster_AutoTerminationPolicyProperty {
	var returns *AwsCluster_AutoTerminationPolicyProperty
	_jsii_.Get(
		j,
		"autoTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) BootstrapAction() AwsCluster_BootstrapActionPropertyList {
	var returns AwsCluster_BootstrapActionPropertyList
	_jsii_.Get(
		j,
		"bootstrapAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) BootstrapActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapActionInput",
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

func (j *jsiiProxy_AwsCluster) ClusterState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Configurations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ConfigurationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ConfigurationsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ConfigurationsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJsonInput",
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

func (j *jsiiProxy_AwsCluster) CoreInstanceFleet() AwsCluster_CoreInstanceFleetPropertyOutputReference {
	var returns AwsCluster_CoreInstanceFleetPropertyOutputReference
	_jsii_.Get(
		j,
		"coreInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) CoreInstanceFleetInput() *AwsCluster_CoreInstanceFleetProperty {
	var returns *AwsCluster_CoreInstanceFleetProperty
	_jsii_.Get(
		j,
		"coreInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) CoreInstanceGroup() AwsCluster_CoreInstanceGroupPropertyOutputReference {
	var returns AwsCluster_CoreInstanceGroupPropertyOutputReference
	_jsii_.Get(
		j,
		"coreInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) CoreInstanceGroupInput() *AwsCluster_CoreInstanceGroupProperty {
	var returns *AwsCluster_CoreInstanceGroupProperty
	_jsii_.Get(
		j,
		"coreInstanceGroupInput",
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

func (j *jsiiProxy_AwsCluster) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) CustomAmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiIdInput",
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

func (j *jsiiProxy_AwsCluster) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) EbsRootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Ec2Attributes() AwsCluster_Ec2AttributesPropertyOutputReference {
	var returns AwsCluster_Ec2AttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"ec2Attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Ec2AttributesInput() *AwsCluster_Ec2AttributesProperty {
	var returns *AwsCluster_Ec2AttributesProperty
	_jsii_.Get(
		j,
		"ec2AttributesInput",
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

func (j *jsiiProxy_AwsCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KeepJobFlowAliveWhenNoSteps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KeepJobFlowAliveWhenNoStepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KerberosAttributes() AwsCluster_KerberosAttributesPropertyOutputReference {
	var returns AwsCluster_KerberosAttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"kerberosAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) KerberosAttributesInput() *AwsCluster_KerberosAttributesProperty {
	var returns *AwsCluster_KerberosAttributesProperty
	_jsii_.Get(
		j,
		"kerberosAttributesInput",
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

func (j *jsiiProxy_AwsCluster) ListStepsStates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ListStepsStatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) LogEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) LogEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) LogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) MasterInstanceFleet() AwsCluster_MasterInstanceFleetPropertyOutputReference {
	var returns AwsCluster_MasterInstanceFleetPropertyOutputReference
	_jsii_.Get(
		j,
		"masterInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) MasterInstanceFleetInput() *AwsCluster_MasterInstanceFleetProperty {
	var returns *AwsCluster_MasterInstanceFleetProperty
	_jsii_.Get(
		j,
		"masterInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) MasterInstanceGroup() AwsCluster_MasterInstanceGroupPropertyOutputReference {
	var returns AwsCluster_MasterInstanceGroupPropertyOutputReference
	_jsii_.Get(
		j,
		"masterInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) MasterInstanceGroupInput() *AwsCluster_MasterInstanceGroupProperty {
	var returns *AwsCluster_MasterInstanceGroupProperty
	_jsii_.Get(
		j,
		"masterInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) MasterPublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPublicDns",
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

func (j *jsiiProxy_AwsCluster) OsReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osReleaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) OsReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osReleaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) PlacementGroupConfig() AwsCluster_PlacementGroupConfigPropertyList {
	var returns AwsCluster_PlacementGroupConfigPropertyList
	_jsii_.Get(
		j,
		"placementGroupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) PlacementGroupConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementGroupConfigInput",
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

func (j *jsiiProxy_AwsCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ScaleDownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ScaleDownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) Step() AwsCluster_StepPropertyList {
	var returns AwsCluster_StepPropertyList
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) StepConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) StepConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) StepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepInput",
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

func (j *jsiiProxy_AwsCluster) TerminationProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) TerminationProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionInput",
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

func (j *jsiiProxy_AwsCluster) UnhealthyNodeReplacement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyNodeReplacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) UnhealthyNodeReplacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyNodeReplacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster) VisibleToAllUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster aws_emr_cluster} Resource.
// Experimental.
func NewAwsCluster(scope constructs.Construct, id *string, config *AwsClusterConfig) AwsCluster {
	_init_.Initialize()

	if err := validateNewAwsClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCluster{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster aws_emr_cluster} Resource.
// Experimental.
func NewAwsCluster_Override(a AwsCluster, scope constructs.Construct, id *string, config *AwsClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCluster)SetAdditionalInfo(val *string) {
	if err := j.validateSetAdditionalInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetApplications(val *[]*string) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetAutoscalingRole(val *string) {
	if err := j.validateSetAutoscalingRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingRole",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetConfigurations(val *string) {
	if err := j.validateSetConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetConfigurationsJson(val *string) {
	if err := j.validateSetConfigurationsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationsJson",
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

func (j *jsiiProxy_AwsCluster)SetCustomAmiId(val *string) {
	if err := j.validateSetCustomAmiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAmiId",
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

func (j *jsiiProxy_AwsCluster)SetEbsRootVolumeSize(val *float64) {
	if err := j.validateSetEbsRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
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

func (j *jsiiProxy_AwsCluster)SetKeepJobFlowAliveWhenNoSteps(val interface{}) {
	if err := j.validateSetKeepJobFlowAliveWhenNoStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepJobFlowAliveWhenNoSteps",
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

func (j *jsiiProxy_AwsCluster)SetListStepsStates(val *[]*string) {
	if err := j.validateSetListStepsStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listStepsStates",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetLogEncryptionKmsKeyId(val *string) {
	if err := j.validateSetLogEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetLogUri(val *string) {
	if err := j.validateSetLogUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logUri",
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

func (j *jsiiProxy_AwsCluster)SetOsReleaseLabel(val *string) {
	if err := j.validateSetOsReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osReleaseLabel",
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

func (j *jsiiProxy_AwsCluster)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetScaleDownBehavior(val *string) {
	if err := j.validateSetScaleDownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetStepConcurrencyLevel(val *float64) {
	if err := j.validateSetStepConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
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

func (j *jsiiProxy_AwsCluster)SetTerminationProtection(val interface{}) {
	if err := j.validateSetTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtection",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetUnhealthyNodeReplacement(val interface{}) {
	if err := j.validateSetUnhealthyNodeReplacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyNodeReplacement",
		val,
	)
}

func (j *jsiiProxy_AwsCluster)SetVisibleToAllUsers(val interface{}) {
	if err := j.validateSetVisibleToAllUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleToAllUsers",
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
		"@cdktn/aws-emr.AwsCluster",
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
		"@cdktn/aws-emr.AwsCluster",
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
		"@cdktn/aws-emr.AwsCluster",
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
		"@cdktn/aws-emr.AwsCluster",
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
		"@cdktn/aws-emr.AwsCluster",
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

func (a *jsiiProxy_AwsCluster) PutAutoTerminationPolicy(value *AwsCluster_AutoTerminationPolicyProperty) {
	if err := a.validatePutAutoTerminationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoTerminationPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutBootstrapAction(value interface{}) {
	if err := a.validatePutBootstrapActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBootstrapAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutCoreInstanceFleet(value *AwsCluster_CoreInstanceFleetProperty) {
	if err := a.validatePutCoreInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCoreInstanceFleet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutCoreInstanceGroup(value *AwsCluster_CoreInstanceGroupProperty) {
	if err := a.validatePutCoreInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCoreInstanceGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutEc2Attributes(value *AwsCluster_Ec2AttributesProperty) {
	if err := a.validatePutEc2AttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2Attributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutKerberosAttributes(value *AwsCluster_KerberosAttributesProperty) {
	if err := a.validatePutKerberosAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKerberosAttributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutMasterInstanceFleet(value *AwsCluster_MasterInstanceFleetProperty) {
	if err := a.validatePutMasterInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMasterInstanceFleet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutMasterInstanceGroup(value *AwsCluster_MasterInstanceGroupProperty) {
	if err := a.validatePutMasterInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMasterInstanceGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutPlacementGroupConfig(value interface{}) {
	if err := a.validatePutPlacementGroupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlacementGroupConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster) PutStep(value interface{}) {
	if err := a.validatePutStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStep",
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

func (a *jsiiProxy_AwsCluster) ResetAdditionalInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetApplications() {
	_jsii_.InvokeVoid(
		a,
		"resetApplications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetAutoscalingRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoscalingRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetAutoTerminationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoTerminationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetBootstrapAction() {
	_jsii_.InvokeVoid(
		a,
		"resetBootstrapAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetConfigurationsJson() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurationsJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetCoreInstanceFleet() {
	_jsii_.InvokeVoid(
		a,
		"resetCoreInstanceFleet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetCoreInstanceGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetCoreInstanceGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetCustomAmiId() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomAmiId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetEbsRootVolumeSize() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsRootVolumeSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetEc2Attributes() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2Attributes",
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

func (a *jsiiProxy_AwsCluster) ResetKeepJobFlowAliveWhenNoSteps() {
	_jsii_.InvokeVoid(
		a,
		"resetKeepJobFlowAliveWhenNoSteps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetKerberosAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetKerberosAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetListStepsStates() {
	_jsii_.InvokeVoid(
		a,
		"resetListStepsStates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetLogEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetLogEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetLogUri() {
	_jsii_.InvokeVoid(
		a,
		"resetLogUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetMasterInstanceFleet() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterInstanceFleet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetMasterInstanceGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterInstanceGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetOsReleaseLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetOsReleaseLabel",
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

func (a *jsiiProxy_AwsCluster) ResetPlacementGroupConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroupConfig",
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

func (a *jsiiProxy_AwsCluster) ResetScaleDownBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleDownBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetStep() {
	_jsii_.InvokeVoid(
		a,
		"resetStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetStepConcurrencyLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetStepConcurrencyLevel",
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

func (a *jsiiProxy_AwsCluster) ResetTerminationProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetUnhealthyNodeReplacement() {
	_jsii_.InvokeVoid(
		a,
		"resetUnhealthyNodeReplacement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster) ResetVisibleToAllUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetVisibleToAllUsers",
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

