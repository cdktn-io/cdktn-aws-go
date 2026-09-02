package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssm/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsssm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task aws_ssm_maintenance_window_task}.
// Experimental.
type TfMaintenanceWindowTask interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	CutoffBehavior() *string
	// Experimental.
	SetCutoffBehavior(val *string)
	// Experimental.
	CutoffBehaviorInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaxConcurrency() *string
	// Experimental.
	SetMaxConcurrency(val *string)
	// Experimental.
	MaxConcurrencyInput() *string
	// Experimental.
	MaxErrors() *string
	// Experimental.
	SetMaxErrors(val *string)
	// Experimental.
	MaxErrorsInput() *string
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
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
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
	ServiceRoleArn() *string
	// Experimental.
	SetServiceRoleArn(val *string)
	// Experimental.
	ServiceRoleArnInput() *string
	// Experimental.
	Targets() TfMaintenanceWindowTask_TargetsPropertyList
	// Experimental.
	TargetsInput() interface{}
	// Experimental.
	TaskArn() *string
	// Experimental.
	SetTaskArn(val *string)
	// Experimental.
	TaskArnInput() *string
	// Experimental.
	TaskInvocationParameters() TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference
	// Experimental.
	TaskInvocationParametersInput() *TfMaintenanceWindowTask_TaskInvocationParametersProperty
	// Experimental.
	TaskType() *string
	// Experimental.
	SetTaskType(val *string)
	// Experimental.
	TaskTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	WindowId() *string
	// Experimental.
	SetWindowId(val *string)
	// Experimental.
	WindowIdInput() *string
	// Experimental.
	WindowTaskId() *string
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
	PutTargets(value interface{})
	// Experimental.
	PutTaskInvocationParameters(value *TfMaintenanceWindowTask_TaskInvocationParametersProperty)
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
	ResetCutoffBehavior()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetId()
	// Experimental.
	ResetMaxConcurrency()
	// Experimental.
	ResetMaxErrors()
	// Experimental.
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPriority()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetServiceRoleArn()
	// Experimental.
	ResetTargets()
	// Experimental.
	ResetTaskInvocationParameters()
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

// The jsii proxy struct for TfMaintenanceWindowTask
type jsiiProxy_TfMaintenanceWindowTask struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) CutoffBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cutoffBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) CutoffBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cutoffBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) MaxConcurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) MaxConcurrencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) MaxErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) MaxErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) Targets() TfMaintenanceWindowTask_TargetsPropertyList {
	var returns TfMaintenanceWindowTask_TargetsPropertyList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TaskArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TaskInvocationParameters() TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference {
	var returns TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"taskInvocationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TaskInvocationParametersInput() *TfMaintenanceWindowTask_TaskInvocationParametersProperty {
	var returns *TfMaintenanceWindowTask_TaskInvocationParametersProperty
	_jsii_.Get(
		j,
		"taskInvocationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TaskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) WindowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) WindowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask) WindowTaskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowTaskId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task aws_ssm_maintenance_window_task} Resource.
// Experimental.
func NewTfMaintenanceWindowTask(scope constructs.Construct, id *string, config *TfMaintenanceWindowTaskConfig) TfMaintenanceWindowTask {
	_init_.Initialize()

	if err := validateNewTfMaintenanceWindowTaskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMaintenanceWindowTask{}

	_jsii_.Create(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task aws_ssm_maintenance_window_task} Resource.
// Experimental.
func NewTfMaintenanceWindowTask_Override(t TfMaintenanceWindowTask, scope constructs.Construct, id *string, config *TfMaintenanceWindowTaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetCutoffBehavior(val *string) {
	if err := j.validateSetCutoffBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cutoffBehavior",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetMaxConcurrency(val *string) {
	if err := j.validateSetMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetMaxErrors(val *string) {
	if err := j.validateSetMaxErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxErrors",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetTaskArn(val *string) {
	if err := j.validateSetTaskArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskArn",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetTaskType(val *string) {
	if err := j.validateSetTaskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask)SetWindowId(val *string) {
	if err := j.validateSetWindowIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowId",
		val,
	)
}

// Generates CDKTN code for importing a TfMaintenanceWindowTask resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfMaintenanceWindowTask_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfMaintenanceWindowTask_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask",
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
func TfMaintenanceWindowTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMaintenanceWindowTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfMaintenanceWindowTask_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMaintenanceWindowTask_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfMaintenanceWindowTask_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMaintenanceWindowTask_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfMaintenanceWindowTask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfMaintenanceWindowTask) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) PutTargets(value interface{}) {
	if err := t.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargets",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) PutTaskInvocationParameters(value *TfMaintenanceWindowTask_TaskInvocationParametersProperty) {
	if err := t.validatePutTaskInvocationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTaskInvocationParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetCutoffBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetCutoffBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetMaxConcurrency() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxConcurrency",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetMaxErrors() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxErrors",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetPriority() {
	_jsii_.InvokeVoid(
		t,
		"resetPriority",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetTargets() {
	_jsii_.InvokeVoid(
		t,
		"resetTargets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ResetTaskInvocationParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskInvocationParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

