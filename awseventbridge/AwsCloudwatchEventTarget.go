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
type AwsCloudwatchEventTarget interface {
	cdktn.TerraformResource
	// Experimental.
	AppsyncTarget() AwsCloudwatchEventTarget_AppsyncTargetPropertyOutputReference
	// Experimental.
	AppsyncTargetInput() *AwsCloudwatchEventTarget_AppsyncTargetProperty
	// Experimental.
	Arn() *string
	// Experimental.
	SetArn(val *string)
	// Experimental.
	ArnInput() *string
	// Experimental.
	BatchTarget() AwsCloudwatchEventTarget_BatchTargetPropertyOutputReference
	// Experimental.
	BatchTargetInput() *AwsCloudwatchEventTarget_BatchTargetProperty
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
	DeadLetterConfig() AwsCloudwatchEventTarget_DeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *AwsCloudwatchEventTarget_DeadLetterConfigProperty
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EcsTarget() AwsCloudwatchEventTarget_EcsTargetPropertyOutputReference
	// Experimental.
	EcsTargetInput() *AwsCloudwatchEventTarget_EcsTargetProperty
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
	HttpTarget() AwsCloudwatchEventTarget_HttpTargetPropertyOutputReference
	// Experimental.
	HttpTargetInput() *AwsCloudwatchEventTarget_HttpTargetProperty
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
	InputTransformer() AwsCloudwatchEventTarget_InputTransformerPropertyOutputReference
	// Experimental.
	InputTransformerInput() *AwsCloudwatchEventTarget_InputTransformerProperty
	// Experimental.
	KinesisTarget() AwsCloudwatchEventTarget_KinesisTargetPropertyOutputReference
	// Experimental.
	KinesisTargetInput() *AwsCloudwatchEventTarget_KinesisTargetProperty
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
	RedshiftTarget() AwsCloudwatchEventTarget_RedshiftTargetPropertyOutputReference
	// Experimental.
	RedshiftTargetInput() *AwsCloudwatchEventTarget_RedshiftTargetProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RetryPolicy() AwsCloudwatchEventTarget_RetryPolicyPropertyOutputReference
	// Experimental.
	RetryPolicyInput() *AwsCloudwatchEventTarget_RetryPolicyProperty
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
	RunCommandTargets() AwsCloudwatchEventTarget_RunCommandTargetsPropertyList
	// Experimental.
	RunCommandTargetsInput() interface{}
	// Experimental.
	SagemakerPipelineTarget() AwsCloudwatchEventTarget_SagemakerPipelineTargetPropertyOutputReference
	// Experimental.
	SagemakerPipelineTargetInput() *AwsCloudwatchEventTarget_SagemakerPipelineTargetProperty
	// Experimental.
	SqsTarget() AwsCloudwatchEventTarget_SqsTargetPropertyOutputReference
	// Experimental.
	SqsTargetInput() *AwsCloudwatchEventTarget_SqsTargetProperty
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
	PutAppsyncTarget(value *AwsCloudwatchEventTarget_AppsyncTargetProperty)
	// Experimental.
	PutBatchTarget(value *AwsCloudwatchEventTarget_BatchTargetProperty)
	// Experimental.
	PutDeadLetterConfig(value *AwsCloudwatchEventTarget_DeadLetterConfigProperty)
	// Experimental.
	PutEcsTarget(value *AwsCloudwatchEventTarget_EcsTargetProperty)
	// Experimental.
	PutHttpTarget(value *AwsCloudwatchEventTarget_HttpTargetProperty)
	// Experimental.
	PutInputTransformer(value *AwsCloudwatchEventTarget_InputTransformerProperty)
	// Experimental.
	PutKinesisTarget(value *AwsCloudwatchEventTarget_KinesisTargetProperty)
	// Experimental.
	PutRedshiftTarget(value *AwsCloudwatchEventTarget_RedshiftTargetProperty)
	// Experimental.
	PutRetryPolicy(value *AwsCloudwatchEventTarget_RetryPolicyProperty)
	// Experimental.
	PutRunCommandTargets(value interface{})
	// Experimental.
	PutSagemakerPipelineTarget(value *AwsCloudwatchEventTarget_SagemakerPipelineTargetProperty)
	// Experimental.
	PutSqsTarget(value *AwsCloudwatchEventTarget_SqsTargetProperty)
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

// The jsii proxy struct for AwsCloudwatchEventTarget
type jsiiProxy_AwsCloudwatchEventTarget struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) AppsyncTarget() AwsCloudwatchEventTarget_AppsyncTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_AppsyncTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"appsyncTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) AppsyncTargetInput() *AwsCloudwatchEventTarget_AppsyncTargetProperty {
	var returns *AwsCloudwatchEventTarget_AppsyncTargetProperty
	_jsii_.Get(
		j,
		"appsyncTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) BatchTarget() AwsCloudwatchEventTarget_BatchTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_BatchTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"batchTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) BatchTargetInput() *AwsCloudwatchEventTarget_BatchTargetProperty {
	var returns *AwsCloudwatchEventTarget_BatchTargetProperty
	_jsii_.Get(
		j,
		"batchTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) DeadLetterConfig() AwsCloudwatchEventTarget_DeadLetterConfigPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_DeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) DeadLetterConfigInput() *AwsCloudwatchEventTarget_DeadLetterConfigProperty {
	var returns *AwsCloudwatchEventTarget_DeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) EcsTarget() AwsCloudwatchEventTarget_EcsTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_EcsTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) EcsTargetInput() *AwsCloudwatchEventTarget_EcsTargetProperty {
	var returns *AwsCloudwatchEventTarget_EcsTargetProperty
	_jsii_.Get(
		j,
		"ecsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) EventBusNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) HttpTarget() AwsCloudwatchEventTarget_HttpTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_HttpTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"httpTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) HttpTargetInput() *AwsCloudwatchEventTarget_HttpTargetProperty {
	var returns *AwsCloudwatchEventTarget_HttpTargetProperty
	_jsii_.Get(
		j,
		"httpTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) InputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) InputTransformer() AwsCloudwatchEventTarget_InputTransformerPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_InputTransformerPropertyOutputReference
	_jsii_.Get(
		j,
		"inputTransformer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) InputTransformerInput() *AwsCloudwatchEventTarget_InputTransformerProperty {
	var returns *AwsCloudwatchEventTarget_InputTransformerProperty
	_jsii_.Get(
		j,
		"inputTransformerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) KinesisTarget() AwsCloudwatchEventTarget_KinesisTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_KinesisTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) KinesisTargetInput() *AwsCloudwatchEventTarget_KinesisTargetProperty {
	var returns *AwsCloudwatchEventTarget_KinesisTargetProperty
	_jsii_.Get(
		j,
		"kinesisTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RedshiftTarget() AwsCloudwatchEventTarget_RedshiftTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_RedshiftTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RedshiftTargetInput() *AwsCloudwatchEventTarget_RedshiftTargetProperty {
	var returns *AwsCloudwatchEventTarget_RedshiftTargetProperty
	_jsii_.Get(
		j,
		"redshiftTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RetryPolicy() AwsCloudwatchEventTarget_RetryPolicyPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_RetryPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RetryPolicyInput() *AwsCloudwatchEventTarget_RetryPolicyProperty {
	var returns *AwsCloudwatchEventTarget_RetryPolicyProperty
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) Rule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RunCommandTargets() AwsCloudwatchEventTarget_RunCommandTargetsPropertyList {
	var returns AwsCloudwatchEventTarget_RunCommandTargetsPropertyList
	_jsii_.Get(
		j,
		"runCommandTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) RunCommandTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runCommandTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) SagemakerPipelineTarget() AwsCloudwatchEventTarget_SagemakerPipelineTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_SagemakerPipelineTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) SagemakerPipelineTargetInput() *AwsCloudwatchEventTarget_SagemakerPipelineTargetProperty {
	var returns *AwsCloudwatchEventTarget_SagemakerPipelineTargetProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) SqsTarget() AwsCloudwatchEventTarget_SqsTargetPropertyOutputReference {
	var returns AwsCloudwatchEventTarget_SqsTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) SqsTargetInput() *AwsCloudwatchEventTarget_SqsTargetProperty {
	var returns *AwsCloudwatchEventTarget_SqsTargetProperty
	_jsii_.Get(
		j,
		"sqsTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) TargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchEventTarget) TerraformResourceType() *string {
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
func NewAwsCloudwatchEventTarget(scope constructs.Construct, id *string, config *AwsCloudwatchEventTargetConfig) AwsCloudwatchEventTarget {
	_init_.Initialize()

	if err := validateNewAwsCloudwatchEventTargetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudwatchEventTarget{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsCloudwatchEventTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target aws_cloudwatch_event_target} Resource.
// Experimental.
func NewAwsCloudwatchEventTarget_Override(a AwsCloudwatchEventTarget, scope constructs.Construct, id *string, config *AwsCloudwatchEventTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsCloudwatchEventTarget",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetEventBusName(val *string) {
	if err := j.validateSetEventBusNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventBusName",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetInputPath(val *string) {
	if err := j.validateSetInputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputPath",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetRule(val *string) {
	if err := j.validateSetRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rule",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchEventTarget)SetTargetId(val *string) {
	if err := j.validateSetTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

// Generates CDKTN code for importing a AwsCloudwatchEventTarget resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCloudwatchEventTarget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCloudwatchEventTarget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsCloudwatchEventTarget",
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
func AwsCloudwatchEventTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudwatchEventTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsCloudwatchEventTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCloudwatchEventTarget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudwatchEventTarget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsCloudwatchEventTarget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCloudwatchEventTarget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudwatchEventTarget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eventbridge.AwsCloudwatchEventTarget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCloudwatchEventTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eventbridge.AwsCloudwatchEventTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCloudwatchEventTarget) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutAppsyncTarget(value *AwsCloudwatchEventTarget_AppsyncTargetProperty) {
	if err := a.validatePutAppsyncTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppsyncTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutBatchTarget(value *AwsCloudwatchEventTarget_BatchTargetProperty) {
	if err := a.validatePutBatchTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBatchTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutDeadLetterConfig(value *AwsCloudwatchEventTarget_DeadLetterConfigProperty) {
	if err := a.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutEcsTarget(value *AwsCloudwatchEventTarget_EcsTargetProperty) {
	if err := a.validatePutEcsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutHttpTarget(value *AwsCloudwatchEventTarget_HttpTargetProperty) {
	if err := a.validatePutHttpTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutInputTransformer(value *AwsCloudwatchEventTarget_InputTransformerProperty) {
	if err := a.validatePutInputTransformerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputTransformer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutKinesisTarget(value *AwsCloudwatchEventTarget_KinesisTargetProperty) {
	if err := a.validatePutKinesisTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutRedshiftTarget(value *AwsCloudwatchEventTarget_RedshiftTargetProperty) {
	if err := a.validatePutRedshiftTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutRetryPolicy(value *AwsCloudwatchEventTarget_RetryPolicyProperty) {
	if err := a.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutRunCommandTargets(value interface{}) {
	if err := a.validatePutRunCommandTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRunCommandTargets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutSagemakerPipelineTarget(value *AwsCloudwatchEventTarget_SagemakerPipelineTargetProperty) {
	if err := a.validatePutSagemakerPipelineTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSagemakerPipelineTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) PutSqsTarget(value *AwsCloudwatchEventTarget_SqsTargetProperty) {
	if err := a.validatePutSqsTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetAppsyncTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetAppsyncTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetBatchTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetEcsTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetEventBusName() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBusName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetHttpTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetInput() {
	_jsii_.InvokeVoid(
		a,
		"resetInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetInputPath() {
	_jsii_.InvokeVoid(
		a,
		"resetInputPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetInputTransformer() {
	_jsii_.InvokeVoid(
		a,
		"resetInputTransformer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetKinesisTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetRedshiftTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetRunCommandTargets() {
	_jsii_.InvokeVoid(
		a,
		"resetRunCommandTargets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetSagemakerPipelineTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerPipelineTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetSqsTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ResetTargetId() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchEventTarget) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

