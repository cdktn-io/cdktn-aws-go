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
type AwsEmrCluster interface {
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
	AutoTerminationPolicy() AwsEmrCluster_AutoTerminationPolicyPropertyOutputReference
	// Experimental.
	AutoTerminationPolicyInput() *AwsEmrCluster_AutoTerminationPolicyProperty
	// Experimental.
	BootstrapAction() AwsEmrCluster_BootstrapActionPropertyList
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
	CoreInstanceFleet() AwsEmrCluster_CoreInstanceFleetPropertyOutputReference
	// Experimental.
	CoreInstanceFleetInput() *AwsEmrCluster_CoreInstanceFleetProperty
	// Experimental.
	CoreInstanceGroup() AwsEmrCluster_CoreInstanceGroupPropertyOutputReference
	// Experimental.
	CoreInstanceGroupInput() *AwsEmrCluster_CoreInstanceGroupProperty
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
	Ec2Attributes() AwsEmrCluster_Ec2AttributesPropertyOutputReference
	// Experimental.
	Ec2AttributesInput() *AwsEmrCluster_Ec2AttributesProperty
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
	KerberosAttributes() AwsEmrCluster_KerberosAttributesPropertyOutputReference
	// Experimental.
	KerberosAttributesInput() *AwsEmrCluster_KerberosAttributesProperty
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
	MasterInstanceFleet() AwsEmrCluster_MasterInstanceFleetPropertyOutputReference
	// Experimental.
	MasterInstanceFleetInput() *AwsEmrCluster_MasterInstanceFleetProperty
	// Experimental.
	MasterInstanceGroup() AwsEmrCluster_MasterInstanceGroupPropertyOutputReference
	// Experimental.
	MasterInstanceGroupInput() *AwsEmrCluster_MasterInstanceGroupProperty
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
	PlacementGroupConfig() AwsEmrCluster_PlacementGroupConfigPropertyList
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
	Step() AwsEmrCluster_StepPropertyList
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
	PutAutoTerminationPolicy(value *AwsEmrCluster_AutoTerminationPolicyProperty)
	// Experimental.
	PutBootstrapAction(value interface{})
	// Experimental.
	PutCoreInstanceFleet(value *AwsEmrCluster_CoreInstanceFleetProperty)
	// Experimental.
	PutCoreInstanceGroup(value *AwsEmrCluster_CoreInstanceGroupProperty)
	// Experimental.
	PutEc2Attributes(value *AwsEmrCluster_Ec2AttributesProperty)
	// Experimental.
	PutKerberosAttributes(value *AwsEmrCluster_KerberosAttributesProperty)
	// Experimental.
	PutMasterInstanceFleet(value *AwsEmrCluster_MasterInstanceFleetProperty)
	// Experimental.
	PutMasterInstanceGroup(value *AwsEmrCluster_MasterInstanceGroupProperty)
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

// The jsii proxy struct for AwsEmrCluster
type jsiiProxy_AwsEmrCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsEmrCluster) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) AdditionalInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Applications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) AutoscalingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) AutoscalingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) AutoTerminationPolicy() AwsEmrCluster_AutoTerminationPolicyPropertyOutputReference {
	var returns AwsEmrCluster_AutoTerminationPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"autoTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) AutoTerminationPolicyInput() *AwsEmrCluster_AutoTerminationPolicyProperty {
	var returns *AwsEmrCluster_AutoTerminationPolicyProperty
	_jsii_.Get(
		j,
		"autoTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) BootstrapAction() AwsEmrCluster_BootstrapActionPropertyList {
	var returns AwsEmrCluster_BootstrapActionPropertyList
	_jsii_.Get(
		j,
		"bootstrapAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) BootstrapActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ClusterState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Configurations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ConfigurationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ConfigurationsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ConfigurationsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) CoreInstanceFleet() AwsEmrCluster_CoreInstanceFleetPropertyOutputReference {
	var returns AwsEmrCluster_CoreInstanceFleetPropertyOutputReference
	_jsii_.Get(
		j,
		"coreInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) CoreInstanceFleetInput() *AwsEmrCluster_CoreInstanceFleetProperty {
	var returns *AwsEmrCluster_CoreInstanceFleetProperty
	_jsii_.Get(
		j,
		"coreInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) CoreInstanceGroup() AwsEmrCluster_CoreInstanceGroupPropertyOutputReference {
	var returns AwsEmrCluster_CoreInstanceGroupPropertyOutputReference
	_jsii_.Get(
		j,
		"coreInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) CoreInstanceGroupInput() *AwsEmrCluster_CoreInstanceGroupProperty {
	var returns *AwsEmrCluster_CoreInstanceGroupProperty
	_jsii_.Get(
		j,
		"coreInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) CustomAmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) EbsRootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Ec2Attributes() AwsEmrCluster_Ec2AttributesPropertyOutputReference {
	var returns AwsEmrCluster_Ec2AttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"ec2Attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Ec2AttributesInput() *AwsEmrCluster_Ec2AttributesProperty {
	var returns *AwsEmrCluster_Ec2AttributesProperty
	_jsii_.Get(
		j,
		"ec2AttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) KeepJobFlowAliveWhenNoSteps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) KeepJobFlowAliveWhenNoStepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) KerberosAttributes() AwsEmrCluster_KerberosAttributesPropertyOutputReference {
	var returns AwsEmrCluster_KerberosAttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"kerberosAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) KerberosAttributesInput() *AwsEmrCluster_KerberosAttributesProperty {
	var returns *AwsEmrCluster_KerberosAttributesProperty
	_jsii_.Get(
		j,
		"kerberosAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ListStepsStates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ListStepsStatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listStepsStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) LogEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) LogEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) LogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) MasterInstanceFleet() AwsEmrCluster_MasterInstanceFleetPropertyOutputReference {
	var returns AwsEmrCluster_MasterInstanceFleetPropertyOutputReference
	_jsii_.Get(
		j,
		"masterInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) MasterInstanceFleetInput() *AwsEmrCluster_MasterInstanceFleetProperty {
	var returns *AwsEmrCluster_MasterInstanceFleetProperty
	_jsii_.Get(
		j,
		"masterInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) MasterInstanceGroup() AwsEmrCluster_MasterInstanceGroupPropertyOutputReference {
	var returns AwsEmrCluster_MasterInstanceGroupPropertyOutputReference
	_jsii_.Get(
		j,
		"masterInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) MasterInstanceGroupInput() *AwsEmrCluster_MasterInstanceGroupProperty {
	var returns *AwsEmrCluster_MasterInstanceGroupProperty
	_jsii_.Get(
		j,
		"masterInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) MasterPublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPublicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) OsReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osReleaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) OsReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osReleaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) PlacementGroupConfig() AwsEmrCluster_PlacementGroupConfigPropertyList {
	var returns AwsEmrCluster_PlacementGroupConfigPropertyList
	_jsii_.Get(
		j,
		"placementGroupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) PlacementGroupConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementGroupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ScaleDownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ScaleDownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Step() AwsEmrCluster_StepPropertyList {
	var returns AwsEmrCluster_StepPropertyList
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) StepConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) StepConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) StepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TerminationProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TerminationProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) UnhealthyNodeReplacement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyNodeReplacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) UnhealthyNodeReplacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyNodeReplacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster) VisibleToAllUsersInput() interface{} {
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
func NewAwsEmrCluster(scope constructs.Construct, id *string, config *AwsEmrClusterConfig) AwsEmrCluster {
	_init_.Initialize()

	if err := validateNewAwsEmrClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrCluster{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster aws_emr_cluster} Resource.
// Experimental.
func NewAwsEmrCluster_Override(a AwsEmrCluster, scope constructs.Construct, id *string, config *AwsEmrClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetAdditionalInfo(val *string) {
	if err := j.validateSetAdditionalInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetApplications(val *[]*string) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetAutoscalingRole(val *string) {
	if err := j.validateSetAutoscalingRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingRole",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetConfigurations(val *string) {
	if err := j.validateSetConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetConfigurationsJson(val *string) {
	if err := j.validateSetConfigurationsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationsJson",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetCustomAmiId(val *string) {
	if err := j.validateSetCustomAmiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetEbsRootVolumeSize(val *float64) {
	if err := j.validateSetEbsRootVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetKeepJobFlowAliveWhenNoSteps(val interface{}) {
	if err := j.validateSetKeepJobFlowAliveWhenNoStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepJobFlowAliveWhenNoSteps",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetListStepsStates(val *[]*string) {
	if err := j.validateSetListStepsStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listStepsStates",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetLogEncryptionKmsKeyId(val *string) {
	if err := j.validateSetLogEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetLogUri(val *string) {
	if err := j.validateSetLogUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetOsReleaseLabel(val *string) {
	if err := j.validateSetOsReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osReleaseLabel",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetScaleDownBehavior(val *string) {
	if err := j.validateSetScaleDownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetStepConcurrencyLevel(val *float64) {
	if err := j.validateSetStepConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetTerminationProtection(val interface{}) {
	if err := j.validateSetTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtection",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetUnhealthyNodeReplacement(val interface{}) {
	if err := j.validateSetUnhealthyNodeReplacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyNodeReplacement",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster)SetVisibleToAllUsers(val interface{}) {
	if err := j.validateSetVisibleToAllUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleToAllUsers",
		val,
	)
}

// Generates CDKTN code for importing a AwsEmrCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsEmrCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsEmrCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr.AwsEmrCluster",
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
func AwsEmrCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEmrCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr.AwsEmrCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEmrCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEmrCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr.AwsEmrCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEmrCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEmrCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr.AwsEmrCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEmrCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-emr.AwsEmrCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEmrCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsEmrCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsEmrCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsEmrCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEmrCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsEmrCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEmrCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutAutoTerminationPolicy(value *AwsEmrCluster_AutoTerminationPolicyProperty) {
	if err := a.validatePutAutoTerminationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoTerminationPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutBootstrapAction(value interface{}) {
	if err := a.validatePutBootstrapActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBootstrapAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutCoreInstanceFleet(value *AwsEmrCluster_CoreInstanceFleetProperty) {
	if err := a.validatePutCoreInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCoreInstanceFleet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutCoreInstanceGroup(value *AwsEmrCluster_CoreInstanceGroupProperty) {
	if err := a.validatePutCoreInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCoreInstanceGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutEc2Attributes(value *AwsEmrCluster_Ec2AttributesProperty) {
	if err := a.validatePutEc2AttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2Attributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutKerberosAttributes(value *AwsEmrCluster_KerberosAttributesProperty) {
	if err := a.validatePutKerberosAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKerberosAttributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutMasterInstanceFleet(value *AwsEmrCluster_MasterInstanceFleetProperty) {
	if err := a.validatePutMasterInstanceFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMasterInstanceFleet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutMasterInstanceGroup(value *AwsEmrCluster_MasterInstanceGroupProperty) {
	if err := a.validatePutMasterInstanceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMasterInstanceGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutPlacementGroupConfig(value interface{}) {
	if err := a.validatePutPlacementGroupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlacementGroupConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) PutStep(value interface{}) {
	if err := a.validatePutStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetAdditionalInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetApplications() {
	_jsii_.InvokeVoid(
		a,
		"resetApplications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetAutoscalingRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoscalingRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetAutoTerminationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoTerminationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetBootstrapAction() {
	_jsii_.InvokeVoid(
		a,
		"resetBootstrapAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetConfigurationsJson() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurationsJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetCoreInstanceFleet() {
	_jsii_.InvokeVoid(
		a,
		"resetCoreInstanceFleet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetCoreInstanceGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetCoreInstanceGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetCustomAmiId() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomAmiId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetEbsRootVolumeSize() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsRootVolumeSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetEc2Attributes() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2Attributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetKeepJobFlowAliveWhenNoSteps() {
	_jsii_.InvokeVoid(
		a,
		"resetKeepJobFlowAliveWhenNoSteps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetKerberosAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetKerberosAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetListStepsStates() {
	_jsii_.InvokeVoid(
		a,
		"resetListStepsStates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetLogEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetLogEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetLogUri() {
	_jsii_.InvokeVoid(
		a,
		"resetLogUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetMasterInstanceFleet() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterInstanceFleet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetMasterInstanceGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterInstanceGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetOsReleaseLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetOsReleaseLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetPlacementGroupConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroupConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetScaleDownBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleDownBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetStep() {
	_jsii_.InvokeVoid(
		a,
		"resetStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetStepConcurrencyLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetStepConcurrencyLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetTerminationProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetUnhealthyNodeReplacement() {
	_jsii_.InvokeVoid(
		a,
		"resetUnhealthyNodeReplacement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) ResetVisibleToAllUsers() {
	_jsii_.InvokeVoid(
		a,
		"resetVisibleToAllUsers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

