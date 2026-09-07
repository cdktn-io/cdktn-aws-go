package eventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridge/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/eventbridge/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target}.
// Experimental.
type AwsTarget interface {
	cdktn.TerraformResource
	// Experimental.
	AppsyncTarget() AwsTarget_AppsyncTargetPropertyOutputReference
	// Experimental.
	AppsyncTargetInput() *AwsTarget_AppsyncTargetProperty
	// Experimental.
	Arn() *string
	// Experimental.
	SetArn(val *string)
	// Experimental.
	ArnInput() *string
	// Experimental.
	BatchTarget() AwsTarget_BatchTargetPropertyOutputReference
	// Experimental.
	BatchTargetInput() *AwsTarget_BatchTargetProperty
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
	DeadLetterConfig() AwsTarget_DeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *AwsTarget_DeadLetterConfigProperty
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EcsTarget() AwsTarget_EcsTargetPropertyOutputReference
	// Experimental.
	EcsTargetInput() *AwsTarget_EcsTargetProperty
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
	HttpTarget() AwsTarget_HttpTargetPropertyOutputReference
	// Experimental.
	HttpTargetInput() *AwsTarget_HttpTargetProperty
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
	InputTransformer() AwsTarget_InputTransformerPropertyOutputReference
	// Experimental.
	InputTransformerInput() *AwsTarget_InputTransformerProperty
	// Experimental.
	KinesisTarget() AwsTarget_KinesisTargetPropertyOutputReference
	// Experimental.
	KinesisTargetInput() *AwsTarget_KinesisTargetProperty
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
	RedshiftTarget() AwsTarget_RedshiftTargetPropertyOutputReference
	// Experimental.
	RedshiftTargetInput() *AwsTarget_RedshiftTargetProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RetryPolicy() AwsTarget_RetryPolicyPropertyOutputReference
	// Experimental.
	RetryPolicyInput() *AwsTarget_RetryPolicyProperty
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
	RunCommandTargets() AwsTarget_RunCommandTargetsPropertyList
	// Experimental.
	RunCommandTargetsInput() interface{}
	// Experimental.
	SagemakerPipelineTarget() AwsTarget_SagemakerPipelineTargetPropertyOutputReference
	// Experimental.
	SagemakerPipelineTargetInput() *AwsTarget_SagemakerPipelineTargetProperty
	// Experimental.
	SqsTarget() AwsTarget_SqsTargetPropertyOutputReference
	// Experimental.
	SqsTargetInput() *AwsTarget_SqsTargetProperty
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
	PutAppsyncTarget(value *AwsTarget_AppsyncTargetProperty)
	// Experimental.
	PutBatchTarget(value *AwsTarget_BatchTargetProperty)
	// Experimental.
	PutDeadLetterConfig(value *AwsTarget_DeadLetterConfigProperty)
	// Experimental.
	PutEcsTarget(value *AwsTarget_EcsTargetProperty)
	// Experimental.
	PutHttpTarget(value *AwsTarget_HttpTargetProperty)
	// Experimental.
	PutInputTransformer(value *AwsTarget_InputTransformerProperty)
	// Experimental.
	PutKinesisTarget(value *AwsTarget_KinesisTargetProperty)
	// Experimental.
	PutRedshiftTarget(value *AwsTarget_RedshiftTargetProperty)
	// Experimental.
	PutRetryPolicy(value *AwsTarget_RetryPolicyProperty)
	// Experimental.
	PutRunCommandTargets(value interface{})
	// Experimental.
	PutSagemakerPipelineTarget(value *AwsTarget_SagemakerPipelineTargetProperty)
	// Experimental.
	PutSqsTarget(value *AwsTarget_SqsTargetProperty)
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

// The jsii proxy struct for AwsTarget
type jsiiProxy_AwsTarget struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsTarget) AppsyncTarget() AwsTarget_AppsyncTargetPropertyOutputReference {
	var returns AwsTarget_AppsyncTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"appsyncTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) AppsyncTargetInput() *AwsTarget_AppsyncTargetProperty {
	var returns *AwsTarget_AppsyncTargetProperty
	_jsii_.Get(
		j,
		"appsyncTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) BatchTarget() AwsTarget_BatchTargetPropertyOutputReference {
	var returns AwsTarget_BatchTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"batchTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) BatchTargetInput() *AwsTarget_BatchTargetProperty {
	var returns *AwsTarget_BatchTargetProperty
	_jsii_.Get(
		j,
		"batchTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) DeadLetterConfig() AwsTarget_DeadLetterConfigPropertyOutputReference {
	var returns AwsTarget_DeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) DeadLetterConfigInput() *AwsTarget_DeadLetterConfigProperty {
	var returns *AwsTarget_DeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) EcsTarget() AwsTarget_EcsTargetPropertyOutputReference {
	var returns AwsTarget_EcsTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) EcsTargetInput() *AwsTarget_EcsTargetProperty {
	var returns *AwsTarget_EcsTargetProperty
	_jsii_.Get(
		j,
		"ecsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) EventBusNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) HttpTarget() AwsTarget_HttpTargetPropertyOutputReference {
	var returns AwsTarget_HttpTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"httpTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) HttpTargetInput() *AwsTarget_HttpTargetProperty {
	var returns *AwsTarget_HttpTargetProperty
	_jsii_.Get(
		j,
		"httpTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) InputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) InputTransformer() AwsTarget_InputTransformerPropertyOutputReference {
	var returns AwsTarget_InputTransformerPropertyOutputReference
	_jsii_.Get(
		j,
		"inputTransformer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) InputTransformerInput() *AwsTarget_InputTransformerProperty {
	var returns *AwsTarget_InputTransformerProperty
	_jsii_.Get(
		j,
		"inputTransformerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) KinesisTarget() AwsTarget_KinesisTargetPropertyOutputReference {
	var returns AwsTarget_KinesisTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) KinesisTargetInput() *AwsTarget_KinesisTargetProperty {
	var returns *AwsTarget_KinesisTargetProperty
	_jsii_.Get(
		j,
		"kinesisTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RedshiftTarget() AwsTarget_RedshiftTargetPropertyOutputReference {
	var returns AwsTarget_RedshiftTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RedshiftTargetInput() *AwsTarget_RedshiftTargetProperty {
	var returns *AwsTarget_RedshiftTargetProperty
	_jsii_.Get(
		j,
		"redshiftTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RetryPolicy() AwsTarget_RetryPolicyPropertyOutputReference {
	var returns AwsTarget_RetryPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RetryPolicyInput() *AwsTarget_RetryPolicyProperty {
	var returns *AwsTarget_RetryPolicyProperty
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) Rule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RunCommandTargets() AwsTarget_RunCommandTargetsPropertyList {
	var returns AwsTarget_RunCommandTargetsPropertyList
	_jsii_.Get(
		j,
		"runCommandTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) RunCommandTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runCommandTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) SagemakerPipelineTarget() AwsTarget_SagemakerPipelineTargetPropertyOutputReference {
	var returns AwsTarget_SagemakerPipelineTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) SagemakerPipelineTargetInput() *AwsTarget_SagemakerPipelineTargetProperty {
	var returns *AwsTarget_SagemakerPipelineTargetProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) SqsTarget() AwsTarget_SqsTargetPropertyOutputReference {
	var returns AwsTarget_SqsTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) SqsTargetInput() *AwsTarget_SqsTargetProperty {
	var returns *AwsTarget_SqsTargetProperty
	_jsii_.Get(
		j,
		"sqsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) TargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget) TerraformResourceType() *string {
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
func NewAwsTarget(scope constructs.Construct, id *string, config *AwsTargetConfig) AwsTarget {
	_init_.Initialize()

	if err := validateNewAwsTargetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTarget{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target} Resource.
// Experimental.
func NewAwsTarget_Override(a AwsTarget, scope constructs.Construct, id *string, config *AwsTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsTarget",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsTarget)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetEventBusName(val *string) {
	if err := j.validateSetEventBusNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventBusName",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetInputPath(val *string) {
	if err := j.validateSetInputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputPath",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetRule(val *string) {
	if err := j.validateSetRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rule",
		val,
	)
}

func (j *jsiiProxy_AwsTarget)SetTargetId(val *string) {
	if err := j.validateSetTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

// Generates CDKTN code for importing a AwsTarget resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsTarget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsTarget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsTarget",
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
func AwsTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTarget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTarget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsTarget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTarget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTarget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsTarget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eventbridge.AwsTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsTarget) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsTarget) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTarget) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTarget) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTarget) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTarget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsTarget) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTarget) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsTarget) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTarget) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsTarget) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTarget) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsTarget) PutAppsyncTarget(value *AwsTarget_AppsyncTargetProperty) {
	if err := a.validatePutAppsyncTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppsyncTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutBatchTarget(value *AwsTarget_BatchTargetProperty) {
	if err := a.validatePutBatchTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBatchTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutDeadLetterConfig(value *AwsTarget_DeadLetterConfigProperty) {
	if err := a.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutEcsTarget(value *AwsTarget_EcsTargetProperty) {
	if err := a.validatePutEcsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutHttpTarget(value *AwsTarget_HttpTargetProperty) {
	if err := a.validatePutHttpTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutInputTransformer(value *AwsTarget_InputTransformerProperty) {
	if err := a.validatePutInputTransformerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputTransformer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutKinesisTarget(value *AwsTarget_KinesisTargetProperty) {
	if err := a.validatePutKinesisTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutRedshiftTarget(value *AwsTarget_RedshiftTargetProperty) {
	if err := a.validatePutRedshiftTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutRetryPolicy(value *AwsTarget_RetryPolicyProperty) {
	if err := a.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutRunCommandTargets(value interface{}) {
	if err := a.validatePutRunCommandTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRunCommandTargets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutSagemakerPipelineTarget(value *AwsTarget_SagemakerPipelineTargetProperty) {
	if err := a.validatePutSagemakerPipelineTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSagemakerPipelineTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) PutSqsTarget(value *AwsTarget_SqsTargetProperty) {
	if err := a.validatePutSqsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsTarget) ResetAppsyncTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetAppsyncTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetBatchTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetEcsTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetEventBusName() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBusName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetHttpTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetInput() {
	_jsii_.InvokeVoid(
		a,
		"resetInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetInputPath() {
	_jsii_.InvokeVoid(
		a,
		"resetInputPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetInputTransformer() {
	_jsii_.InvokeVoid(
		a,
		"resetInputTransformer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetKinesisTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetRedshiftTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetRunCommandTargets() {
	_jsii_.InvokeVoid(
		a,
		"resetRunCommandTargets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetSagemakerPipelineTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerPipelineTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetSqsTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) ResetTargetId() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

