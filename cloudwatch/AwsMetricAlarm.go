package cloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudwatch/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/cloudwatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm}.
// Experimental.
type AwsMetricAlarm interface {
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
	EvaluationCriteria() AwsMetricAlarm_EvaluationCriteriaPropertyOutputReference
	// Experimental.
	EvaluationCriteriaInput() *AwsMetricAlarm_EvaluationCriteriaProperty
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
	MetricQuery() AwsMetricAlarm_MetricQueryPropertyList
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
	PutEvaluationCriteria(value *AwsMetricAlarm_EvaluationCriteriaProperty)
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

// The jsii proxy struct for AwsMetricAlarm
type jsiiProxy_AwsMetricAlarm struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsMetricAlarm) ActionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ActionsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) AlarmActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) AlarmActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) AlarmDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) AlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ComparisonOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) DatapointsToAlarmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Dimensions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) DimensionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluateLowSampleCountPercentiles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluateLowSampleCountPercentilesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluationCriteria() AwsMetricAlarm_EvaluationCriteriaPropertyOutputReference {
	var returns AwsMetricAlarm_EvaluationCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"evaluationCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluationCriteriaInput() *AwsMetricAlarm_EvaluationCriteriaProperty {
	var returns *AwsMetricAlarm_EvaluationCriteriaProperty
	_jsii_.Get(
		j,
		"evaluationCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluationInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluationIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) EvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ExtendedStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ExtendedStatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) InsufficientDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) InsufficientDataActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) MetricQuery() AwsMetricAlarm_MetricQueryPropertyList {
	var returns AwsMetricAlarm_MetricQueryPropertyList
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) MetricQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) OkActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) OkActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ThresholdMetricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) ThresholdMetricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TreatMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) TreatMissingDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMetricAlarm) UnitInput() *string {
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
func NewAwsMetricAlarm(scope constructs.Construct, id *string, config *AwsMetricAlarmConfig) AwsMetricAlarm {
	_init_.Initialize()

	if err := validateNewAwsMetricAlarmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMetricAlarm{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch.AwsMetricAlarm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm} Resource.
// Experimental.
func NewAwsMetricAlarm_Override(a AwsMetricAlarm, scope constructs.Construct, id *string, config *AwsMetricAlarmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch.AwsMetricAlarm",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetActionsEnabled(val interface{}) {
	if err := j.validateSetActionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetAlarmActions(val *[]*string) {
	if err := j.validateSetAlarmActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmActions",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetAlarmDescription(val *string) {
	if err := j.validateSetAlarmDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetAlarmName(val *string) {
	if err := j.validateSetAlarmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetDatapointsToAlarm(val *float64) {
	if err := j.validateSetDatapointsToAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datapointsToAlarm",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetDimensions(val *map[string]*string) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetEvaluateLowSampleCountPercentiles(val *string) {
	if err := j.validateSetEvaluateLowSampleCountPercentilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateLowSampleCountPercentiles",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetEvaluationInterval(val *float64) {
	if err := j.validateSetEvaluationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationInterval",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetExtendedStatistic(val *string) {
	if err := j.validateSetExtendedStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedStatistic",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetInsufficientDataActions(val *[]*string) {
	if err := j.validateSetInsufficientDataActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insufficientDataActions",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetOkActions(val *[]*string) {
	if err := j.validateSetOkActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"okActions",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetThresholdMetricId(val *string) {
	if err := j.validateSetThresholdMetricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdMetricId",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetTreatMissingData(val *string) {
	if err := j.validateSetTreatMissingDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatMissingData",
		val,
	)
}

func (j *jsiiProxy_AwsMetricAlarm)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

// Generates CDKTN code for importing a AwsMetricAlarm resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsMetricAlarm_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsMetricAlarm_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.AwsMetricAlarm",
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
func AwsMetricAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMetricAlarm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.AwsMetricAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMetricAlarm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMetricAlarm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.AwsMetricAlarm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMetricAlarm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMetricAlarm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch.AwsMetricAlarm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsMetricAlarm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudwatch.AwsMetricAlarm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMetricAlarm) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMetricAlarm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMetricAlarm) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMetricAlarm) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMetricAlarm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMetricAlarm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMetricAlarm) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMetricAlarm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMetricAlarm) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMetricAlarm) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsMetricAlarm) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) PutEvaluationCriteria(value *AwsMetricAlarm_EvaluationCriteriaProperty) {
	if err := a.validatePutEvaluationCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEvaluationCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) PutMetricQuery(value interface{}) {
	if err := a.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetActionsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetActionsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetAlarmActions() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarmActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetAlarmDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarmDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetComparisonOperator() {
	_jsii_.InvokeVoid(
		a,
		"resetComparisonOperator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetDatapointsToAlarm() {
	_jsii_.InvokeVoid(
		a,
		"resetDatapointsToAlarm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetDimensions() {
	_jsii_.InvokeVoid(
		a,
		"resetDimensions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetEvaluateLowSampleCountPercentiles() {
	_jsii_.InvokeVoid(
		a,
		"resetEvaluateLowSampleCountPercentiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetEvaluationCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetEvaluationCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetEvaluationInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetEvaluationInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetEvaluationPeriods() {
	_jsii_.InvokeVoid(
		a,
		"resetEvaluationPeriods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetExtendedStatistic() {
	_jsii_.InvokeVoid(
		a,
		"resetExtendedStatistic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetInsufficientDataActions() {
	_jsii_.InvokeVoid(
		a,
		"resetInsufficientDataActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetMetricName() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetOkActions() {
	_jsii_.InvokeVoid(
		a,
		"resetOkActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetStatistic() {
	_jsii_.InvokeVoid(
		a,
		"resetStatistic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetThresholdMetricId() {
	_jsii_.InvokeVoid(
		a,
		"resetThresholdMetricId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetTreatMissingData() {
	_jsii_.InvokeVoid(
		a,
		"resetTreatMissingData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) ResetUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMetricAlarm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMetricAlarm) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

