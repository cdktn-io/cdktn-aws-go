package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster aws_emr_cluster}.
// Experimental.
type TfCluster interface {
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
	AutoTerminationPolicy() TfCluster_AutoTerminationPolicyPropertyOutputReference
	// Experimental.
	AutoTerminationPolicyInput() *TfCluster_AutoTerminationPolicyProperty
	// Experimental.
	BootstrapAction() TfCluster_BootstrapActionPropertyList
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
	CoreInstanceFleet() TfCluster_CoreInstanceFleetPropertyOutputReference
	// Experimental.
	CoreInstanceFleetInput() *TfCluster_CoreInstanceFleetProperty
	// Experimental.
	CoreInstanceGroup() TfCluster_CoreInstanceGroupPropertyOutputReference
	// Experimental.
	CoreInstanceGroupInput() *TfCluster_CoreInstanceGroupProperty
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
	Ec2Attributes() TfCluster_Ec2AttributesPropertyOutputReference
	// Experimental.
	Ec2AttributesInput() *TfCluster_Ec2AttributesProperty
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
	KerberosAttributes() TfCluster_KerberosAttributesPropertyOutputReference
	// Experimental.
	KerberosAttributesInput() *TfCluster_KerberosAttributesProperty
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
	MasterInstanceFleet() TfCluster_MasterInstanceFleetPropertyOutputReference
	// Experimental.
	MasterInstanceFleetInput() *TfCluster_MasterInstanceFleetProperty
	// Experimental.
	MasterInstanceGroup() TfCluster_MasterInstanceGroupPropertyOutputReference
	// Experimental.
	MasterInstanceGroupInput() *TfCluster_MasterInstanceGroupProperty
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
	PlacementGroupConfig() TfCluster_PlacementGroupConfigPropertyList
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
	Step() TfCluster_StepPropertyList
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
	PutAutoTerminationPolicy(value *TfCluster_AutoTerminationPolicyProperty)
	// Experimental.
	PutBootstrapAction(value interface{})
	// Experimental.
	PutCoreInstanceFleet(value *TfCluster_CoreInstanceFleetProperty)
	// Experimental.
	PutCoreInstanceGroup(value *TfCluster_CoreInstanceGroupProperty)
	// Experimental.
	PutEc2Attributes(value *TfCluster_Ec2AttributesProperty)
	// Experimental.
	PutKerberosAttributes(value *TfCluster_KerberosAttributesProperty)
	// Experimental.
	PutMasterInstanceFleet(value *TfCluster_MasterInstanceFleetProperty)
	// Experimental.
	PutMasterInstanceGroup(value *TfCluster_MasterInstanceGroupProperty)
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

// The jsii proxy struct for TfCluster
type jsiiProxy_TfCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCluster) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AdditionalInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Applications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationsInput",
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

func (j *jsiiProxy_TfCluster) AutoscalingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AutoscalingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AutoTerminationPolicy() TfCluster_AutoTerminationPolicyPropertyOutputReference {
	var returns TfCluster_AutoTerminationPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"autoTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AutoTerminationPolicyInput() *TfCluster_AutoTerminationPolicyProperty {
	var returns *TfCluster_AutoTerminationPolicyProperty
	_jsii_.Get(
		j,
		"autoTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapAction() TfCluster_BootstrapActionPropertyList {
	var returns TfCluster_BootstrapActionPropertyList
	_jsii_.Get(
		j,
		"bootstrapAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapActionInput",
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

func (j *jsiiProxy_TfCluster) ClusterState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Configurations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ConfigurationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ConfigurationsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ConfigurationsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJsonInput",
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

func (j *jsiiProxy_TfCluster) CoreInstanceFleet() TfCluster_CoreInstanceFleetPropertyOutputReference {
	var returns TfCluster_CoreInstanceFleetPropertyOutputReference
	_jsii_.Get(
		j,
		"coreInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CoreInstanceFleetInput() *TfCluster_CoreInstanceFleetProperty {
	var returns *TfCluster_CoreInstanceFleetProperty
	_jsii_.Get(
		j,
		"coreInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CoreInstanceGroup() TfCluster_CoreInstanceGroupPropertyOutputReference {
	var returns TfCluster_CoreInstanceGroupPropertyOutputReference
	_jsii_.Get(
		j,
		"coreInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CoreInstanceGroupInput() *TfCluster_CoreInstanceGroupProperty {
	var returns *TfCluster_CoreInstanceGroupProperty
	_jsii_.Get(
		j,
		"coreInstanceGroupInput",
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

func (j *jsiiProxy_TfCluster) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CustomAmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiIdInput",
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

func (j *jsiiProxy_TfCluster) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EbsRootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Ec2Attributes() TfCluster_Ec2AttributesPropertyOutputReference {
	var returns TfCluster_Ec2AttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"ec2Attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Ec2AttributesInput() *TfCluster_Ec2AttributesProperty {
	var returns *TfCluster_Ec2AttributesProperty
	_jsii_.Get(
		j,
		"ec2AttributesInput",
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

func (j *jsiiProxy_TfCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KeepJobFlowAliveWhenNoSteps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KeepJobFlowAliveWhenNoStepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KerberosAttributes() TfCluster_KerberosAttributesPropertyOutputReference {
	var returns TfCluster_KerberosAttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"kerberosAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KerberosAttributesInput() *TfCluster_KerberosAttributesProperty {
	var returns *TfCluster_KerberosAttributesProperty
	_jsii_.Get(
		j,
		"kerberosAttributesInput",
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

func (j *jsiiProxy_TfCluster) ListStepsStates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ListStepsStatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) LogEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) LogEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) LogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterInstanceFleet() TfCluster_MasterInstanceFleetPropertyOutputReference {
	var returns TfCluster_MasterInstanceFleetPropertyOutputReference
	_jsii_.Get(
		j,
		"masterInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterInstanceFleetInput() *TfCluster_MasterInstanceFleetProperty {
	var returns *TfCluster_MasterInstanceFleetProperty
	_jsii_.Get(
		j,
		"masterInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterInstanceGroup() TfCluster_MasterInstanceGroupPropertyOutputReference {
	var returns TfCluster_MasterInstanceGroupPropertyOutputReference
	_jsii_.Get(
		j,
		"masterInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterInstanceGroupInput() *TfCluster_MasterInstanceGroupProperty {
	var returns *TfCluster_MasterInstanceGroupProperty
	_jsii_.Get(
		j,
		"masterInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterPublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPublicDns",
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

func (j *jsiiProxy_TfCluster) OsReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osReleaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) OsReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osReleaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PlacementGroupConfig() TfCluster_PlacementGroupConfigPropertyList {
	var returns TfCluster_PlacementGroupConfigPropertyList
	_jsii_.Get(
		j,
		"placementGroupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PlacementGroupConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementGroupConfigInput",
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

func (j *jsiiProxy_TfCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ScaleDownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ScaleDownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Step() TfCluster_StepPropertyList {
	var returns TfCluster_StepPropertyList
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StepConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StepConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepInput",
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

func (j *jsiiProxy_TfCluster) TerminationProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TerminationProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionInput",
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

func (j *jsiiProxy_TfCluster) UnhealthyNodeReplacement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyNodeReplacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) UnhealthyNodeReplacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyNodeReplacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) VisibleToAllUsersInput() interface{} {
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
func NewTfCluster(scope constructs.Construct, id *string, config *TfClusterConfig) TfCluster {
	_init_.Initialize()

	if err := validateNewTfClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster{}

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster aws_emr_cluster} Resource.
// Experimental.
func NewTfCluster_Override(t TfCluster, scope constructs.Construct, id *string, config *TfClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCluster)SetAdditionalInfo(val *string) {
	if err := j.validateSetAdditionalInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetApplications(val *[]*string) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetAutoscalingRole(val *string) {
	if err := j.validateSetAutoscalingRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingRole",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetConfigurations(val *string) {
	if err := j.validateSetConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetConfigurationsJson(val *string) {
	if err := j.validateSetConfigurationsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationsJson",
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

func (j *jsiiProxy_TfCluster)SetCustomAmiId(val *string) {
	if err := j.validateSetCustomAmiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAmiId",
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

func (j *jsiiProxy_TfCluster)SetEbsRootVolumeSize(val *float64) {
	if err := j.validateSetEbsRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
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

func (j *jsiiProxy_TfCluster)SetKeepJobFlowAliveWhenNoSteps(val interface{}) {
	if err := j.validateSetKeepJobFlowAliveWhenNoStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepJobFlowAliveWhenNoSteps",
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

func (j *jsiiProxy_TfCluster)SetListStepsStates(val *[]*string) {
	if err := j.validateSetListStepsStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listStepsStates",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetLogEncryptionKmsKeyId(val *string) {
	if err := j.validateSetLogEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetLogUri(val *string) {
	if err := j.validateSetLogUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logUri",
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

func (j *jsiiProxy_TfCluster)SetOsReleaseLabel(val *string) {
	if err := j.validateSetOsReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osReleaseLabel",
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

func (j *jsiiProxy_TfCluster)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetScaleDownBehavior(val *string) {
	if err := j.validateSetScaleDownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetStepConcurrencyLevel(val *float64) {
	if err := j.validateSetStepConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
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

func (j *jsiiProxy_TfCluster)SetTerminationProtection(val interface{}) {
	if err := j.validateSetTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtection",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetUnhealthyNodeReplacement(val interface{}) {
	if err := j.validateSetUnhealthyNodeReplacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyNodeReplacement",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetVisibleToAllUsers(val interface{}) {
	if err := j.validateSetVisibleToAllUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleToAllUsers",
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
		"@cdktn/aws-emr.TfCluster",
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
		"@cdktn/aws-emr.TfCluster",
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
		"@cdktn/aws-emr.TfCluster",
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
		"@cdktn/aws-emr.TfCluster",
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
		"@cdktn/aws-emr.TfCluster",
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

func (t *jsiiProxy_TfCluster) PutAutoTerminationPolicy(value *TfCluster_AutoTerminationPolicyProperty) {
	if err := t.validatePutAutoTerminationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutoTerminationPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutBootstrapAction(value interface{}) {
	if err := t.validatePutBootstrapActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBootstrapAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutCoreInstanceFleet(value *TfCluster_CoreInstanceFleetProperty) {
	if err := t.validatePutCoreInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCoreInstanceFleet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutCoreInstanceGroup(value *TfCluster_CoreInstanceGroupProperty) {
	if err := t.validatePutCoreInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCoreInstanceGroup",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutEc2Attributes(value *TfCluster_Ec2AttributesProperty) {
	if err := t.validatePutEc2AttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2Attributes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutKerberosAttributes(value *TfCluster_KerberosAttributesProperty) {
	if err := t.validatePutKerberosAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKerberosAttributes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutMasterInstanceFleet(value *TfCluster_MasterInstanceFleetProperty) {
	if err := t.validatePutMasterInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMasterInstanceFleet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutMasterInstanceGroup(value *TfCluster_MasterInstanceGroupProperty) {
	if err := t.validatePutMasterInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMasterInstanceGroup",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutPlacementGroupConfig(value interface{}) {
	if err := t.validatePutPlacementGroupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlacementGroupConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutStep(value interface{}) {
	if err := t.validatePutStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStep",
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

func (t *jsiiProxy_TfCluster) ResetAdditionalInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetApplications() {
	_jsii_.InvokeVoid(
		t,
		"resetApplications",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetAutoscalingRole() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoscalingRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetAutoTerminationPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoTerminationPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetBootstrapAction() {
	_jsii_.InvokeVoid(
		t,
		"resetBootstrapAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetConfigurations() {
	_jsii_.InvokeVoid(
		t,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetConfigurationsJson() {
	_jsii_.InvokeVoid(
		t,
		"resetConfigurationsJson",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetCoreInstanceFleet() {
	_jsii_.InvokeVoid(
		t,
		"resetCoreInstanceFleet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetCoreInstanceGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetCoreInstanceGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetCustomAmiId() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomAmiId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEbsRootVolumeSize() {
	_jsii_.InvokeVoid(
		t,
		"resetEbsRootVolumeSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEc2Attributes() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2Attributes",
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

func (t *jsiiProxy_TfCluster) ResetKeepJobFlowAliveWhenNoSteps() {
	_jsii_.InvokeVoid(
		t,
		"resetKeepJobFlowAliveWhenNoSteps",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetKerberosAttributes() {
	_jsii_.InvokeVoid(
		t,
		"resetKerberosAttributes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetListStepsStates() {
	_jsii_.InvokeVoid(
		t,
		"resetListStepsStates",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetLogEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetLogEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetLogUri() {
	_jsii_.InvokeVoid(
		t,
		"resetLogUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMasterInstanceFleet() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterInstanceFleet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMasterInstanceGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterInstanceGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetOsReleaseLabel() {
	_jsii_.InvokeVoid(
		t,
		"resetOsReleaseLabel",
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

func (t *jsiiProxy_TfCluster) ResetPlacementGroupConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementGroupConfig",
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

func (t *jsiiProxy_TfCluster) ResetScaleDownBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetScaleDownBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetStep() {
	_jsii_.InvokeVoid(
		t,
		"resetStep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetStepConcurrencyLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetStepConcurrencyLevel",
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

func (t *jsiiProxy_TfCluster) ResetTerminationProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetTerminationProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetUnhealthyNodeReplacement() {
	_jsii_.InvokeVoid(
		t,
		"resetUnhealthyNodeReplacement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetVisibleToAllUsers() {
	_jsii_.InvokeVoid(
		t,
		"resetVisibleToAllUsers",
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

