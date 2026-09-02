package awswebservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget aws_budgets_budget}.
// Experimental.
type TfBudget interface {
	cdktn.TerraformResource
	// Experimental.
	AccountId() *string
	// Experimental.
	SetAccountId(val *string)
	// Experimental.
	AccountIdInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoAdjustData() TfBudget_AutoAdjustDataPropertyOutputReference
	// Experimental.
	AutoAdjustDataInput() *TfBudget_AutoAdjustDataProperty
	// Experimental.
	BillingViewArn() *string
	// Experimental.
	SetBillingViewArn(val *string)
	// Experimental.
	BillingViewArnInput() *string
	// Experimental.
	BudgetType() *string
	// Experimental.
	SetBudgetType(val *string)
	// Experimental.
	BudgetTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CostFilter() TfBudget_CostFilterPropertyList
	// Experimental.
	CostFilterInput() interface{}
	// Experimental.
	CostTypes() TfBudget_CostTypesPropertyOutputReference
	// Experimental.
	CostTypesInput() *TfBudget_CostTypesProperty
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	FilterExpression() TfBudget_FilterExpressionPropertyOutputReference
	// Experimental.
	FilterExpressionInput() *TfBudget_FilterExpressionProperty
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
	LimitAmount() *string
	// Experimental.
	SetLimitAmount(val *string)
	// Experimental.
	LimitAmountInput() *string
	// Experimental.
	LimitUnit() *string
	// Experimental.
	SetLimitUnit(val *string)
	// Experimental.
	LimitUnitInput() *string
	// Experimental.
	Metrics() *[]*string
	// Experimental.
	SetMetrics(val *[]*string)
	// Experimental.
	MetricsInput() *[]*string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Notification() TfBudget_NotificationPropertyList
	// Experimental.
	NotificationInput() interface{}
	// Experimental.
	PlannedLimit() TfBudget_PlannedLimitPropertyList
	// Experimental.
	PlannedLimitInput() interface{}
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
	TimePeriodEnd() *string
	// Experimental.
	SetTimePeriodEnd(val *string)
	// Experimental.
	TimePeriodEndInput() *string
	// Experimental.
	TimePeriodStart() *string
	// Experimental.
	SetTimePeriodStart(val *string)
	// Experimental.
	TimePeriodStartInput() *string
	// Experimental.
	TimeUnit() *string
	// Experimental.
	SetTimeUnit(val *string)
	// Experimental.
	TimeUnitInput() *string
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
	PutAutoAdjustData(value *TfBudget_AutoAdjustDataProperty)
	// Experimental.
	PutCostFilter(value interface{})
	// Experimental.
	PutCostTypes(value *TfBudget_CostTypesProperty)
	// Experimental.
	PutFilterExpression(value *TfBudget_FilterExpressionProperty)
	// Experimental.
	PutNotification(value interface{})
	// Experimental.
	PutPlannedLimit(value interface{})
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
	ResetAccountId()
	// Experimental.
	ResetAutoAdjustData()
	// Experimental.
	ResetBillingViewArn()
	// Experimental.
	ResetCostFilter()
	// Experimental.
	ResetCostTypes()
	// Experimental.
	ResetFilterExpression()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLimitAmount()
	// Experimental.
	ResetLimitUnit()
	// Experimental.
	ResetMetrics()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Experimental.
	ResetNotification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlannedLimit()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimePeriodEnd()
	// Experimental.
	ResetTimePeriodStart()
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

// The jsii proxy struct for TfBudget
type jsiiProxy_TfBudget struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfBudget) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) AutoAdjustData() TfBudget_AutoAdjustDataPropertyOutputReference {
	var returns TfBudget_AutoAdjustDataPropertyOutputReference
	_jsii_.Get(
		j,
		"autoAdjustData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) AutoAdjustDataInput() *TfBudget_AutoAdjustDataProperty {
	var returns *TfBudget_AutoAdjustDataProperty
	_jsii_.Get(
		j,
		"autoAdjustDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) BillingViewArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingViewArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) BillingViewArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingViewArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) BudgetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) BudgetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) CostFilter() TfBudget_CostFilterPropertyList {
	var returns TfBudget_CostFilterPropertyList
	_jsii_.Get(
		j,
		"costFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) CostFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"costFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) CostTypes() TfBudget_CostTypesPropertyOutputReference {
	var returns TfBudget_CostTypesPropertyOutputReference
	_jsii_.Get(
		j,
		"costTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) CostTypesInput() *TfBudget_CostTypesProperty {
	var returns *TfBudget_CostTypesProperty
	_jsii_.Get(
		j,
		"costTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) FilterExpression() TfBudget_FilterExpressionPropertyOutputReference {
	var returns TfBudget_FilterExpressionPropertyOutputReference
	_jsii_.Get(
		j,
		"filterExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) FilterExpressionInput() *TfBudget_FilterExpressionProperty {
	var returns *TfBudget_FilterExpressionProperty
	_jsii_.Get(
		j,
		"filterExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) LimitAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) LimitAmountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) LimitUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) LimitUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Metrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) MetricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Notification() TfBudget_NotificationPropertyList {
	var returns TfBudget_NotificationPropertyList
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) NotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) PlannedLimit() TfBudget_PlannedLimitPropertyList {
	var returns TfBudget_PlannedLimitPropertyList
	_jsii_.Get(
		j,
		"plannedLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) PlannedLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plannedLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TimePeriodEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TimePeriodEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TimePeriodStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TimePeriodStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget) TimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget aws_budgets_budget} Resource.
// Experimental.
func NewTfBudget(scope constructs.Construct, id *string, config *TfBudgetConfig) TfBudget {
	_init_.Initialize()

	if err := validateNewTfBudgetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBudget{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget aws_budgets_budget} Resource.
// Experimental.
func NewTfBudget_Override(t TfBudget, scope constructs.Construct, id *string, config *TfBudgetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudget",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfBudget)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetBillingViewArn(val *string) {
	if err := j.validateSetBillingViewArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingViewArn",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetBudgetType(val *string) {
	if err := j.validateSetBudgetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetType",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetLimitAmount(val *string) {
	if err := j.validateSetLimitAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitAmount",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetLimitUnit(val *string) {
	if err := j.validateSetLimitUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitUnit",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetMetrics(val *[]*string) {
	if err := j.validateSetMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metrics",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetTimePeriodEnd(val *string) {
	if err := j.validateSetTimePeriodEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriodEnd",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetTimePeriodStart(val *string) {
	if err := j.validateSetTimePeriodStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriodStart",
		val,
	)
}

func (j *jsiiProxy_TfBudget)SetTimeUnit(val *string) {
	if err := j.validateSetTimeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeUnit",
		val,
	)
}

// Generates CDKTN code for importing a TfBudget resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfBudget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfBudget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.TfBudget",
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
func TfBudget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBudget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.TfBudget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfBudget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBudget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.TfBudget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfBudget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBudget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.TfBudget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfBudget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-web-services-budgets.TfBudget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfBudget) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfBudget) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfBudget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBudget) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBudget) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBudget) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBudget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBudget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBudget) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBudget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBudget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfBudget) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudget) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfBudget) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfBudget) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfBudget) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfBudget) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfBudget) PutAutoAdjustData(value *TfBudget_AutoAdjustDataProperty) {
	if err := t.validatePutAutoAdjustDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutoAdjustData",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget) PutCostFilter(value interface{}) {
	if err := t.validatePutCostFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget) PutCostTypes(value *TfBudget_CostTypesProperty) {
	if err := t.validatePutCostTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostTypes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget) PutFilterExpression(value *TfBudget_FilterExpressionProperty) {
	if err := t.validatePutFilterExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilterExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget) PutNotification(value interface{}) {
	if err := t.validatePutNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget) PutPlannedLimit(value interface{}) {
	if err := t.validatePutPlannedLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlannedLimit",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfBudget) ResetAccountId() {
	_jsii_.InvokeVoid(
		t,
		"resetAccountId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetAutoAdjustData() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoAdjustData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetBillingViewArn() {
	_jsii_.InvokeVoid(
		t,
		"resetBillingViewArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetCostFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetCostFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetCostTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetCostTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetFilterExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetFilterExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetLimitAmount() {
	_jsii_.InvokeVoid(
		t,
		"resetLimitAmount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetLimitUnit() {
	_jsii_.InvokeVoid(
		t,
		"resetLimitUnit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetNotification() {
	_jsii_.InvokeVoid(
		t,
		"resetNotification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetPlannedLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetPlannedLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetTimePeriodEnd() {
	_jsii_.InvokeVoid(
		t,
		"resetTimePeriodEnd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) ResetTimePeriodStart() {
	_jsii_.InvokeVoid(
		t,
		"resetTimePeriodStart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

