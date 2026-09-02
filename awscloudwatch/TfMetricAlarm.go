package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatch/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscloudwatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm}.
// Experimental.
type TfMetricAlarm interface {
	cdktn.TerraformResource
	// Experimental.
	ActionsEnabled() interface{}
	// Experimental.
	SetActionsEnabled(val interface{})
	// Experimental.
	ActionsEnabledInput() interface{}
	// Experimental.
	AlarmActions() *[]*string
	// Experimental.
	SetAlarmActions(val *[]*string)
	// Experimental.
	AlarmActionsInput() *[]*string
	// Experimental.
	AlarmDescription() *string
	// Experimental.
	SetAlarmDescription(val *string)
	// Experimental.
	AlarmDescriptionInput() *string
	// Experimental.
	AlarmName() *string
	// Experimental.
	SetAlarmName(val *string)
	// Experimental.
	AlarmNameInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ComparisonOperator() *string
	// Experimental.
	SetComparisonOperator(val *string)
	// Experimental.
	ComparisonOperatorInput() *string
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
	DatapointsToAlarm() *float64
	// Experimental.
	SetDatapointsToAlarm(val *float64)
	// Experimental.
	DatapointsToAlarmInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Dimensions() *map[string]*string
	// Experimental.
	SetDimensions(val *map[string]*string)
	// Experimental.
	DimensionsInput() *map[string]*string
	// Experimental.
	EvaluateLowSampleCountPercentiles() *string
	// Experimental.
	SetEvaluateLowSampleCountPercentiles(val *string)
	// Experimental.
	EvaluateLowSampleCountPercentilesInput() *string
	// Experimental.
	EvaluationCriteria() TfMetricAlarm_EvaluationCriteriaPropertyOutputReference
	// Experimental.
	EvaluationCriteriaInput() *TfMetricAlarm_EvaluationCriteriaProperty
	// Experimental.
	EvaluationInterval() *float64
	// Experimental.
	SetEvaluationInterval(val *float64)
	// Experimental.
	EvaluationIntervalInput() *float64
	// Experimental.
	EvaluationPeriods() *float64
	// Experimental.
	SetEvaluationPeriods(val *float64)
	// Experimental.
	EvaluationPeriodsInput() *float64
	// Experimental.
	ExtendedStatistic() *string
	// Experimental.
	SetExtendedStatistic(val *string)
	// Experimental.
	ExtendedStatisticInput() *string
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
	InsufficientDataActions() *[]*string
	// Experimental.
	SetInsufficientDataActions(val *[]*string)
	// Experimental.
	InsufficientDataActionsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MetricName() *string
	// Experimental.
	SetMetricName(val *string)
	// Experimental.
	MetricNameInput() *string
	// Experimental.
	MetricQuery() TfMetricAlarm_MetricQueryPropertyList
	// Experimental.
	MetricQueryInput() interface{}
	// Experimental.
	Namespace() *string
	// Experimental.
	SetNamespace(val *string)
	// Experimental.
	NamespaceInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OkActions() *[]*string
	// Experimental.
	SetOkActions(val *[]*string)
	// Experimental.
	OkActionsInput() *[]*string
	// Experimental.
	Period() *float64
	// Experimental.
	SetPeriod(val *float64)
	// Experimental.
	PeriodInput() *float64
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
	Statistic() *string
	// Experimental.
	SetStatistic(val *string)
	// Experimental.
	StatisticInput() *string
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
	Threshold() *float64
	// Experimental.
	SetThreshold(val *float64)
	// Experimental.
	ThresholdInput() *float64
	// Experimental.
	ThresholdMetricId() *string
	// Experimental.
	SetThresholdMetricId(val *string)
	// Experimental.
	ThresholdMetricIdInput() *string
	// Experimental.
	TreatMissingData() *string
	// Experimental.
	SetTreatMissingData(val *string)
	// Experimental.
	TreatMissingDataInput() *string
	// Experimental.
	Unit() *string
	// Experimental.
	SetUnit(val *string)
	// Experimental.
	UnitInput() *string
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
	PutEvaluationCriteria(value *TfMetricAlarm_EvaluationCriteriaProperty)
	// Experimental.
	PutMetricQuery(value interface{})
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
	ResetActionsEnabled()
	// Experimental.
	ResetAlarmActions()
	// Experimental.
	ResetAlarmDescription()
	// Experimental.
	ResetComparisonOperator()
	// Experimental.
	ResetDatapointsToAlarm()
	// Experimental.
	ResetDimensions()
	// Experimental.
	ResetEvaluateLowSampleCountPercentiles()
	// Experimental.
	ResetEvaluationCriteria()
	// Experimental.
	ResetEvaluationInterval()
	// Experimental.
	ResetEvaluationPeriods()
	// Experimental.
	ResetExtendedStatistic()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInsufficientDataActions()
	// Experimental.
	ResetMetricName()
	// Experimental.
	ResetMetricQuery()
	// Experimental.
	ResetNamespace()
	// Experimental.
	ResetOkActions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPeriod()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetStatistic()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetThreshold()
	// Experimental.
	ResetThresholdMetricId()
	// Experimental.
	ResetTreatMissingData()
	// Experimental.
	ResetUnit()
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

// The jsii proxy struct for TfMetricAlarm
type jsiiProxy_TfMetricAlarm struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfMetricAlarm) ActionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ActionsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) AlarmActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) AlarmActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) AlarmDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) AlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ComparisonOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) DatapointsToAlarmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Dimensions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) DimensionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluateLowSampleCountPercentiles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluateLowSampleCountPercentilesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluationCriteria() TfMetricAlarm_EvaluationCriteriaPropertyOutputReference {
	var returns TfMetricAlarm_EvaluationCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"evaluationCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluationCriteriaInput() *TfMetricAlarm_EvaluationCriteriaProperty {
	var returns *TfMetricAlarm_EvaluationCriteriaProperty
	_jsii_.Get(
		j,
		"evaluationCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluationInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluationIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) EvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ExtendedStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ExtendedStatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) InsufficientDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) InsufficientDataActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) MetricQuery() TfMetricAlarm_MetricQueryPropertyList {
	var returns TfMetricAlarm_MetricQueryPropertyList
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) MetricQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) OkActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) OkActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ThresholdMetricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) ThresholdMetricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TreatMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) TreatMissingDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMetricAlarm) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm} Resource.
// Experimental.
func NewTfMetricAlarm(scope constructs.Construct, id *string, config *TfMetricAlarmConfig) TfMetricAlarm {
	_init_.Initialize()

	if err := validateNewTfMetricAlarmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMetricAlarm{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch.TfMetricAlarm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm} Resource.
// Experimental.
func NewTfMetricAlarm_Override(t TfMetricAlarm, scope constructs.Construct, id *string, config *TfMetricAlarmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch.TfMetricAlarm",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetActionsEnabled(val interface{}) {
	if err := j.validateSetActionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetAlarmActions(val *[]*string) {
	if err := j.validateSetAlarmActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmActions",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetAlarmDescription(val *string) {
	if err := j.validateSetAlarmDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetAlarmName(val *string) {
	if err := j.validateSetAlarmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetDatapointsToAlarm(val *float64) {
	if err := j.validateSetDatapointsToAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datapointsToAlarm",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetDimensions(val *map[string]*string) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetEvaluateLowSampleCountPercentiles(val *string) {
	if err := j.validateSetEvaluateLowSampleCountPercentilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateLowSampleCountPercentiles",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetEvaluationInterval(val *float64) {
	if err := j.validateSetEvaluationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationInterval",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetExtendedStatistic(val *string) {
	if err := j.validateSetExtendedStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedStatistic",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetInsufficientDataActions(val *[]*string) {
	if err := j.validateSetInsufficientDataActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insufficientDataActions",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetOkActions(val *[]*string) {
	if err := j.validateSetOkActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"okActions",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetThresholdMetricId(val *string) {
	if err := j.validateSetThresholdMetricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdMetricId",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetTreatMissingData(val *string) {
	if err := j.validateSetTreatMissingDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatMissingData",
		val,
	)
}

func (j *jsiiProxy_TfMetricAlarm)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

// Generates CDKTN code for importing a TfMetricAlarm resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfMetricAlarm_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfMetricAlarm_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.TfMetricAlarm",
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
func TfMetricAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMetricAlarm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.TfMetricAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfMetricAlarm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMetricAlarm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.TfMetricAlarm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfMetricAlarm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMetricAlarm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.TfMetricAlarm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfMetricAlarm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudwatch.TfMetricAlarm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfMetricAlarm) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfMetricAlarm) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfMetricAlarm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMetricAlarm) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMetricAlarm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMetricAlarm) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMetricAlarm) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMetricAlarm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMetricAlarm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMetricAlarm) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMetricAlarm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMetricAlarm) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMetricAlarm) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfMetricAlarm) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMetricAlarm) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfMetricAlarm) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfMetricAlarm) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfMetricAlarm) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfMetricAlarm) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfMetricAlarm) PutEvaluationCriteria(value *TfMetricAlarm_EvaluationCriteriaProperty) {
	if err := t.validatePutEvaluationCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEvaluationCriteria",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMetricAlarm) PutMetricQuery(value interface{}) {
	if err := t.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMetricAlarm) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetActionsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetActionsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetAlarmActions() {
	_jsii_.InvokeVoid(
		t,
		"resetAlarmActions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetAlarmDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetAlarmDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetComparisonOperator() {
	_jsii_.InvokeVoid(
		t,
		"resetComparisonOperator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetDatapointsToAlarm() {
	_jsii_.InvokeVoid(
		t,
		"resetDatapointsToAlarm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetDimensions() {
	_jsii_.InvokeVoid(
		t,
		"resetDimensions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetEvaluateLowSampleCountPercentiles() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluateLowSampleCountPercentiles",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetEvaluationCriteria() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluationCriteria",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetEvaluationInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluationInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetEvaluationPeriods() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluationPeriods",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetExtendedStatistic() {
	_jsii_.InvokeVoid(
		t,
		"resetExtendedStatistic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetInsufficientDataActions() {
	_jsii_.InvokeVoid(
		t,
		"resetInsufficientDataActions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetMetricName() {
	_jsii_.InvokeVoid(
		t,
		"resetMetricName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		t,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetOkActions() {
	_jsii_.InvokeVoid(
		t,
		"resetOkActions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetStatistic() {
	_jsii_.InvokeVoid(
		t,
		"resetStatistic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetThreshold() {
	_jsii_.InvokeVoid(
		t,
		"resetThreshold",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetThresholdMetricId() {
	_jsii_.InvokeVoid(
		t,
		"resetThresholdMetricId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetTreatMissingData() {
	_jsii_.InvokeVoid(
		t,
		"resetTreatMissingData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) ResetUnit() {
	_jsii_.InvokeVoid(
		t,
		"resetUnit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMetricAlarm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMetricAlarm) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMetricAlarm) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMetricAlarm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMetricAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMetricAlarm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMetricAlarm) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

