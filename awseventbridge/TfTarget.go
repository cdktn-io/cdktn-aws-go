package awseventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridge/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awseventbridge/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target}.
// Experimental.
type TfTarget interface {
	cdktn.TerraformResource
	// Experimental.
	AppsyncTarget() TfTarget_AppsyncTargetPropertyOutputReference
	// Experimental.
	AppsyncTargetInput() *TfTarget_AppsyncTargetProperty
	// Experimental.
	Arn() *string
	// Experimental.
	SetArn(val *string)
	// Experimental.
	ArnInput() *string
	// Experimental.
	BatchTarget() TfTarget_BatchTargetPropertyOutputReference
	// Experimental.
	BatchTargetInput() *TfTarget_BatchTargetProperty
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
	DeadLetterConfig() TfTarget_DeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *TfTarget_DeadLetterConfigProperty
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EcsTarget() TfTarget_EcsTargetPropertyOutputReference
	// Experimental.
	EcsTargetInput() *TfTarget_EcsTargetProperty
	// Experimental.
	EventBusName() *string
	// Experimental.
	SetEventBusName(val *string)
	// Experimental.
	EventBusNameInput() *string
	// Experimental.
	ForceDestroy() interface{}
	// Experimental.
	SetForceDestroy(val interface{})
	// Experimental.
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HttpTarget() TfTarget_HttpTargetPropertyOutputReference
	// Experimental.
	HttpTargetInput() *TfTarget_HttpTargetProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	Input() *string
	// Experimental.
	SetInput(val *string)
	// Experimental.
	InputInput() *string
	// Experimental.
	InputPath() *string
	// Experimental.
	SetInputPath(val *string)
	// Experimental.
	InputPathInput() *string
	// Experimental.
	InputTransformer() TfTarget_InputTransformerPropertyOutputReference
	// Experimental.
	InputTransformerInput() *TfTarget_InputTransformerProperty
	// Experimental.
	KinesisTarget() TfTarget_KinesisTargetPropertyOutputReference
	// Experimental.
	KinesisTargetInput() *TfTarget_KinesisTargetProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	RedshiftTarget() TfTarget_RedshiftTargetPropertyOutputReference
	// Experimental.
	RedshiftTargetInput() *TfTarget_RedshiftTargetProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RetryPolicy() TfTarget_RetryPolicyPropertyOutputReference
	// Experimental.
	RetryPolicyInput() *TfTarget_RetryPolicyProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	Rule() *string
	// Experimental.
	SetRule(val *string)
	// Experimental.
	RuleInput() *string
	// Experimental.
	RunCommandTargets() TfTarget_RunCommandTargetsPropertyList
	// Experimental.
	RunCommandTargetsInput() interface{}
	// Experimental.
	SagemakerPipelineTarget() TfTarget_SagemakerPipelineTargetPropertyOutputReference
	// Experimental.
	SagemakerPipelineTargetInput() *TfTarget_SagemakerPipelineTargetProperty
	// Experimental.
	SqsTarget() TfTarget_SqsTargetPropertyOutputReference
	// Experimental.
	SqsTargetInput() *TfTarget_SqsTargetProperty
	// Experimental.
	TargetId() *string
	// Experimental.
	SetTargetId(val *string)
	// Experimental.
	TargetIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAppsyncTarget(value *TfTarget_AppsyncTargetProperty)
	// Experimental.
	PutBatchTarget(value *TfTarget_BatchTargetProperty)
	// Experimental.
	PutDeadLetterConfig(value *TfTarget_DeadLetterConfigProperty)
	// Experimental.
	PutEcsTarget(value *TfTarget_EcsTargetProperty)
	// Experimental.
	PutHttpTarget(value *TfTarget_HttpTargetProperty)
	// Experimental.
	PutInputTransformer(value *TfTarget_InputTransformerProperty)
	// Experimental.
	PutKinesisTarget(value *TfTarget_KinesisTargetProperty)
	// Experimental.
	PutRedshiftTarget(value *TfTarget_RedshiftTargetProperty)
	// Experimental.
	PutRetryPolicy(value *TfTarget_RetryPolicyProperty)
	// Experimental.
	PutRunCommandTargets(value interface{})
	// Experimental.
	PutSagemakerPipelineTarget(value *TfTarget_SagemakerPipelineTargetProperty)
	// Experimental.
	PutSqsTarget(value *TfTarget_SqsTargetProperty)
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
	ResetAppsyncTarget()
	// Experimental.
	ResetBatchTarget()
	// Experimental.
	ResetDeadLetterConfig()
	// Experimental.
	ResetEcsTarget()
	// Experimental.
	ResetEventBusName()
	// Experimental.
	ResetForceDestroy()
	// Experimental.
	ResetHttpTarget()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInput()
	// Experimental.
	ResetInputPath()
	// Experimental.
	ResetInputTransformer()
	// Experimental.
	ResetKinesisTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRedshiftTarget()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRetryPolicy()
	// Experimental.
	ResetRoleArn()
	// Experimental.
	ResetRunCommandTargets()
	// Experimental.
	ResetSagemakerPipelineTarget()
	// Experimental.
	ResetSqsTarget()
	// Experimental.
	ResetTargetId()
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

// The jsii proxy struct for TfTarget
type jsiiProxy_TfTarget struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTarget) AppsyncTarget() TfTarget_AppsyncTargetPropertyOutputReference {
	var returns TfTarget_AppsyncTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"appsyncTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) AppsyncTargetInput() *TfTarget_AppsyncTargetProperty {
	var returns *TfTarget_AppsyncTargetProperty
	_jsii_.Get(
		j,
		"appsyncTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) BatchTarget() TfTarget_BatchTargetPropertyOutputReference {
	var returns TfTarget_BatchTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"batchTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) BatchTargetInput() *TfTarget_BatchTargetProperty {
	var returns *TfTarget_BatchTargetProperty
	_jsii_.Get(
		j,
		"batchTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) DeadLetterConfig() TfTarget_DeadLetterConfigPropertyOutputReference {
	var returns TfTarget_DeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) DeadLetterConfigInput() *TfTarget_DeadLetterConfigProperty {
	var returns *TfTarget_DeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) EcsTarget() TfTarget_EcsTargetPropertyOutputReference {
	var returns TfTarget_EcsTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) EcsTargetInput() *TfTarget_EcsTargetProperty {
	var returns *TfTarget_EcsTargetProperty
	_jsii_.Get(
		j,
		"ecsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) EventBusNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) HttpTarget() TfTarget_HttpTargetPropertyOutputReference {
	var returns TfTarget_HttpTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"httpTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) HttpTargetInput() *TfTarget_HttpTargetProperty {
	var returns *TfTarget_HttpTargetProperty
	_jsii_.Get(
		j,
		"httpTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) InputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) InputTransformer() TfTarget_InputTransformerPropertyOutputReference {
	var returns TfTarget_InputTransformerPropertyOutputReference
	_jsii_.Get(
		j,
		"inputTransformer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) InputTransformerInput() *TfTarget_InputTransformerProperty {
	var returns *TfTarget_InputTransformerProperty
	_jsii_.Get(
		j,
		"inputTransformerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) KinesisTarget() TfTarget_KinesisTargetPropertyOutputReference {
	var returns TfTarget_KinesisTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) KinesisTargetInput() *TfTarget_KinesisTargetProperty {
	var returns *TfTarget_KinesisTargetProperty
	_jsii_.Get(
		j,
		"kinesisTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RedshiftTarget() TfTarget_RedshiftTargetPropertyOutputReference {
	var returns TfTarget_RedshiftTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RedshiftTargetInput() *TfTarget_RedshiftTargetProperty {
	var returns *TfTarget_RedshiftTargetProperty
	_jsii_.Get(
		j,
		"redshiftTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RetryPolicy() TfTarget_RetryPolicyPropertyOutputReference {
	var returns TfTarget_RetryPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RetryPolicyInput() *TfTarget_RetryPolicyProperty {
	var returns *TfTarget_RetryPolicyProperty
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) Rule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RunCommandTargets() TfTarget_RunCommandTargetsPropertyList {
	var returns TfTarget_RunCommandTargetsPropertyList
	_jsii_.Get(
		j,
		"runCommandTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) RunCommandTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runCommandTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) SagemakerPipelineTarget() TfTarget_SagemakerPipelineTargetPropertyOutputReference {
	var returns TfTarget_SagemakerPipelineTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) SagemakerPipelineTargetInput() *TfTarget_SagemakerPipelineTargetProperty {
	var returns *TfTarget_SagemakerPipelineTargetProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) SqsTarget() TfTarget_SqsTargetPropertyOutputReference {
	var returns TfTarget_SqsTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) SqsTargetInput() *TfTarget_SqsTargetProperty {
	var returns *TfTarget_SqsTargetProperty
	_jsii_.Get(
		j,
		"sqsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) TargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target} Resource.
// Experimental.
func NewTfTarget(scope constructs.Construct, id *string, config *TfTargetConfig) TfTarget {
	_init_.Initialize()

	if err := validateNewTfTargetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTarget{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target} Resource.
// Experimental.
func NewTfTarget_Override(t TfTarget, scope constructs.Construct, id *string, config *TfTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfTarget",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTarget)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetEventBusName(val *string) {
	if err := j.validateSetEventBusNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventBusName",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetInputPath(val *string) {
	if err := j.validateSetInputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputPath",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetRule(val *string) {
	if err := j.validateSetRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rule",
		val,
	)
}

func (j *jsiiProxy_TfTarget)SetTargetId(val *string) {
	if err := j.validateSetTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

// Generates CDKTN code for importing a TfTarget resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTarget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTarget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.TfTarget",
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
func TfTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.TfTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTarget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTarget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.TfTarget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTarget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTarget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.TfTarget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eventbridge.TfTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTarget) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTarget) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTarget) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTarget) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTarget) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTarget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTarget) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTarget) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfTarget) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTarget) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTarget) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTarget) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTarget) PutAppsyncTarget(value *TfTarget_AppsyncTargetProperty) {
	if err := t.validatePutAppsyncTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAppsyncTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutBatchTarget(value *TfTarget_BatchTargetProperty) {
	if err := t.validatePutBatchTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBatchTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutDeadLetterConfig(value *TfTarget_DeadLetterConfigProperty) {
	if err := t.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutEcsTarget(value *TfTarget_EcsTargetProperty) {
	if err := t.validatePutEcsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcsTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutHttpTarget(value *TfTarget_HttpTargetProperty) {
	if err := t.validatePutHttpTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttpTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutInputTransformer(value *TfTarget_InputTransformerProperty) {
	if err := t.validatePutInputTransformerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputTransformer",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutKinesisTarget(value *TfTarget_KinesisTargetProperty) {
	if err := t.validatePutKinesisTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutRedshiftTarget(value *TfTarget_RedshiftTargetProperty) {
	if err := t.validatePutRedshiftTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshiftTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutRetryPolicy(value *TfTarget_RetryPolicyProperty) {
	if err := t.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutRunCommandTargets(value interface{}) {
	if err := t.validatePutRunCommandTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRunCommandTargets",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutSagemakerPipelineTarget(value *TfTarget_SagemakerPipelineTargetProperty) {
	if err := t.validatePutSagemakerPipelineTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSagemakerPipelineTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) PutSqsTarget(value *TfTarget_SqsTargetProperty) {
	if err := t.validatePutSqsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqsTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTarget) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTarget) ResetAppsyncTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetAppsyncTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetBatchTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetBatchTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetEcsTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetEcsTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetEventBusName() {
	_jsii_.InvokeVoid(
		t,
		"resetEventBusName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetHttpTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetInput() {
	_jsii_.InvokeVoid(
		t,
		"resetInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetInputPath() {
	_jsii_.InvokeVoid(
		t,
		"resetInputPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetInputTransformer() {
	_jsii_.InvokeVoid(
		t,
		"resetInputTransformer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetKinesisTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetRedshiftTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshiftTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetRunCommandTargets() {
	_jsii_.InvokeVoid(
		t,
		"resetRunCommandTargets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetSagemakerPipelineTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetSagemakerPipelineTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetSqsTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) ResetTargetId() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

