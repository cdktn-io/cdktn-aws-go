package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job aws_sagemaker_labeling_job}.
// Experimental.
type TfLabelingJob interface {
	cdktn.TerraformResource
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	FailureReason() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HumanTaskConfig() TfLabelingJob_HumanTaskConfigPropertyList
	// Experimental.
	HumanTaskConfigInput() interface{}
	// Experimental.
	InputConfig() TfLabelingJob_InputConfigPropertyList
	// Experimental.
	InputConfigInput() interface{}
	// Experimental.
	JobReferenceCode() *string
	// Experimental.
	LabelAttributeName() *string
	// Experimental.
	SetLabelAttributeName(val *string)
	// Experimental.
	LabelAttributeNameInput() *string
	// Experimental.
	LabelCategoryConfigS3Uri() *string
	// Experimental.
	SetLabelCategoryConfigS3Uri(val *string)
	// Experimental.
	LabelCategoryConfigS3UriInput() *string
	// Experimental.
	LabelCounters() TfLabelingJob_LabelCountersPropertyList
	// Experimental.
	LabelingJobAlgorithmsConfig() TfLabelingJob_LabelingJobAlgorithmsConfigPropertyList
	// Experimental.
	LabelingJobAlgorithmsConfigInput() interface{}
	// Experimental.
	LabelingJobArn() *string
	// Experimental.
	LabelingJobName() *string
	// Experimental.
	SetLabelingJobName(val *string)
	// Experimental.
	LabelingJobNameInput() *string
	// Experimental.
	LabelingJobStatus() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutputConfig() TfLabelingJob_OutputConfigPropertyList
	// Experimental.
	OutputConfigInput() interface{}
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
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	StoppingConditions() TfLabelingJob_StoppingConditionsPropertyList
	// Experimental.
	StoppingConditionsInput() interface{}
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
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
	PutHumanTaskConfig(value interface{})
	// Experimental.
	PutInputConfig(value interface{})
	// Experimental.
	PutLabelingJobAlgorithmsConfig(value interface{})
	// Experimental.
	PutOutputConfig(value interface{})
	// Experimental.
	PutStoppingConditions(value interface{})
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
	ResetHumanTaskConfig()
	// Experimental.
	ResetInputConfig()
	// Experimental.
	ResetLabelCategoryConfigS3Uri()
	// Experimental.
	ResetLabelingJobAlgorithmsConfig()
	// Experimental.
	ResetOutputConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetStoppingConditions()
	// Experimental.
	ResetTags()
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

// The jsii proxy struct for TfLabelingJob
type jsiiProxy_TfLabelingJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfLabelingJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) HumanTaskConfig() TfLabelingJob_HumanTaskConfigPropertyList {
	var returns TfLabelingJob_HumanTaskConfigPropertyList
	_jsii_.Get(
		j,
		"humanTaskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) HumanTaskConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"humanTaskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) InputConfig() TfLabelingJob_InputConfigPropertyList {
	var returns TfLabelingJob_InputConfigPropertyList
	_jsii_.Get(
		j,
		"inputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) InputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) JobReferenceCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobReferenceCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelAttributeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelAttributeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelCategoryConfigS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelCategoryConfigS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelCategoryConfigS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelCategoryConfigS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelCounters() TfLabelingJob_LabelCountersPropertyList {
	var returns TfLabelingJob_LabelCountersPropertyList
	_jsii_.Get(
		j,
		"labelCounters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelingJobAlgorithmsConfig() TfLabelingJob_LabelingJobAlgorithmsConfigPropertyList {
	var returns TfLabelingJob_LabelingJobAlgorithmsConfigPropertyList
	_jsii_.Get(
		j,
		"labelingJobAlgorithmsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelingJobAlgorithmsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelingJobAlgorithmsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelingJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelingJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelingJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) LabelingJobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) OutputConfig() TfLabelingJob_OutputConfigPropertyList {
	var returns TfLabelingJob_OutputConfigPropertyList
	_jsii_.Get(
		j,
		"outputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) OutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) StoppingConditions() TfLabelingJob_StoppingConditionsPropertyList {
	var returns TfLabelingJob_StoppingConditionsPropertyList
	_jsii_.Get(
		j,
		"stoppingConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) StoppingConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job aws_sagemaker_labeling_job} Resource.
// Experimental.
func NewTfLabelingJob(scope constructs.Construct, id *string, config *TfLabelingJobConfig) TfLabelingJob {
	_init_.Initialize()

	if err := validateNewTfLabelingJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLabelingJob{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job aws_sagemaker_labeling_job} Resource.
// Experimental.
func NewTfLabelingJob_Override(t TfLabelingJob, scope constructs.Construct, id *string, config *TfLabelingJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetLabelAttributeName(val *string) {
	if err := j.validateSetLabelAttributeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelAttributeName",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetLabelCategoryConfigS3Uri(val *string) {
	if err := j.validateSetLabelCategoryConfigS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelCategoryConfigS3Uri",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetLabelingJobName(val *string) {
	if err := j.validateSetLabelingJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelingJobName",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a TfLabelingJob resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfLabelingJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfLabelingJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob",
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
func TfLabelingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLabelingJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLabelingJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLabelingJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLabelingJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLabelingJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfLabelingJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfLabelingJob) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfLabelingJob) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfLabelingJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLabelingJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLabelingJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLabelingJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLabelingJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLabelingJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLabelingJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLabelingJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLabelingJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLabelingJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfLabelingJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLabelingJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfLabelingJob) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLabelingJob) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfLabelingJob) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLabelingJob) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfLabelingJob) PutHumanTaskConfig(value interface{}) {
	if err := t.validatePutHumanTaskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHumanTaskConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob) PutInputConfig(value interface{}) {
	if err := t.validatePutInputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob) PutLabelingJobAlgorithmsConfig(value interface{}) {
	if err := t.validatePutLabelingJobAlgorithmsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLabelingJobAlgorithmsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob) PutOutputConfig(value interface{}) {
	if err := t.validatePutOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob) PutStoppingConditions(value interface{}) {
	if err := t.validatePutStoppingConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStoppingConditions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetHumanTaskConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetHumanTaskConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetInputConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetInputConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetLabelCategoryConfigS3Uri() {
	_jsii_.InvokeVoid(
		t,
		"resetLabelCategoryConfigS3Uri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetLabelingJobAlgorithmsConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLabelingJobAlgorithmsConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetOutputConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetStoppingConditions() {
	_jsii_.InvokeVoid(
		t,
		"resetStoppingConditions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

