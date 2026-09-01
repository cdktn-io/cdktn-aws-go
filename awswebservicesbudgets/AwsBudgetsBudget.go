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
type AwsBudgetsBudget interface {
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
	AutoAdjustData() AwsBudgetsBudget_AutoAdjustDataPropertyOutputReference
	// Experimental.
	AutoAdjustDataInput() *AwsBudgetsBudget_AutoAdjustDataProperty
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
	CostFilter() AwsBudgetsBudget_CostFilterPropertyList
	// Experimental.
	CostFilterInput() interface{}
	// Experimental.
	CostTypes() AwsBudgetsBudget_CostTypesPropertyOutputReference
	// Experimental.
	CostTypesInput() *AwsBudgetsBudget_CostTypesProperty
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	FilterExpression() AwsBudgetsBudget_FilterExpressionPropertyOutputReference
	// Experimental.
	FilterExpressionInput() *AwsBudgetsBudget_FilterExpressionProperty
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
	Notification() AwsBudgetsBudget_NotificationPropertyList
	// Experimental.
	NotificationInput() interface{}
	// Experimental.
	PlannedLimit() AwsBudgetsBudget_PlannedLimitPropertyList
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
	PutAutoAdjustData(value *AwsBudgetsBudget_AutoAdjustDataProperty)
	// Experimental.
	PutCostFilter(value interface{})
	// Experimental.
	PutCostTypes(value *AwsBudgetsBudget_CostTypesProperty)
	// Experimental.
	PutFilterExpression(value *AwsBudgetsBudget_FilterExpressionProperty)
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

// The jsii proxy struct for AwsBudgetsBudget
type jsiiProxy_AwsBudgetsBudget struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsBudgetsBudget) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) AutoAdjustData() AwsBudgetsBudget_AutoAdjustDataPropertyOutputReference {
	var returns AwsBudgetsBudget_AutoAdjustDataPropertyOutputReference
	_jsii_.Get(
		j,
		"autoAdjustData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) AutoAdjustDataInput() *AwsBudgetsBudget_AutoAdjustDataProperty {
	var returns *AwsBudgetsBudget_AutoAdjustDataProperty
	_jsii_.Get(
		j,
		"autoAdjustDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) BillingViewArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingViewArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) BillingViewArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingViewArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) BudgetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) BudgetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) CostFilter() AwsBudgetsBudget_CostFilterPropertyList {
	var returns AwsBudgetsBudget_CostFilterPropertyList
	_jsii_.Get(
		j,
		"costFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) CostFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"costFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) CostTypes() AwsBudgetsBudget_CostTypesPropertyOutputReference {
	var returns AwsBudgetsBudget_CostTypesPropertyOutputReference
	_jsii_.Get(
		j,
		"costTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) CostTypesInput() *AwsBudgetsBudget_CostTypesProperty {
	var returns *AwsBudgetsBudget_CostTypesProperty
	_jsii_.Get(
		j,
		"costTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) FilterExpression() AwsBudgetsBudget_FilterExpressionPropertyOutputReference {
	var returns AwsBudgetsBudget_FilterExpressionPropertyOutputReference
	_jsii_.Get(
		j,
		"filterExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) FilterExpressionInput() *AwsBudgetsBudget_FilterExpressionProperty {
	var returns *AwsBudgetsBudget_FilterExpressionProperty
	_jsii_.Get(
		j,
		"filterExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) LimitAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) LimitAmountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) LimitUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) LimitUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Metrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) MetricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Notification() AwsBudgetsBudget_NotificationPropertyList {
	var returns AwsBudgetsBudget_NotificationPropertyList
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) NotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) PlannedLimit() AwsBudgetsBudget_PlannedLimitPropertyList {
	var returns AwsBudgetsBudget_PlannedLimitPropertyList
	_jsii_.Get(
		j,
		"plannedLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) PlannedLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plannedLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TimePeriodEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TimePeriodEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TimePeriodStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TimePeriodStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudget) TimeUnitInput() *string {
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
func NewAwsBudgetsBudget(scope constructs.Construct, id *string, config *AwsBudgetsBudgetConfig) AwsBudgetsBudget {
	_init_.Initialize()

	if err := validateNewAwsBudgetsBudgetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBudgetsBudget{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget aws_budgets_budget} Resource.
// Experimental.
func NewAwsBudgetsBudget_Override(a AwsBudgetsBudget, scope constructs.Construct, id *string, config *AwsBudgetsBudgetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudget",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetBillingViewArn(val *string) {
	if err := j.validateSetBillingViewArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingViewArn",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetBudgetType(val *string) {
	if err := j.validateSetBudgetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetType",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetLimitAmount(val *string) {
	if err := j.validateSetLimitAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitAmount",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetLimitUnit(val *string) {
	if err := j.validateSetLimitUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitUnit",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetMetrics(val *[]*string) {
	if err := j.validateSetMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metrics",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetTimePeriodEnd(val *string) {
	if err := j.validateSetTimePeriodEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriodEnd",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetTimePeriodStart(val *string) {
	if err := j.validateSetTimePeriodStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriodStart",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudget)SetTimeUnit(val *string) {
	if err := j.validateSetTimeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeUnit",
		val,
	)
}

// Generates CDKTN code for importing a AwsBudgetsBudget resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsBudgetsBudget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsBudgetsBudget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudget",
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
func AwsBudgetsBudget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBudgetsBudget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBudgetsBudget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBudgetsBudget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBudgetsBudget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBudgetsBudget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsBudgetsBudget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBudgetsBudget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBudgetsBudget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudgetsBudget) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsBudgetsBudget) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) PutAutoAdjustData(value *AwsBudgetsBudget_AutoAdjustDataProperty) {
	if err := a.validatePutAutoAdjustDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoAdjustData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) PutCostFilter(value interface{}) {
	if err := a.validatePutCostFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCostFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) PutCostTypes(value *AwsBudgetsBudget_CostTypesProperty) {
	if err := a.validatePutCostTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCostTypes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) PutFilterExpression(value *AwsBudgetsBudget_FilterExpressionProperty) {
	if err := a.validatePutFilterExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilterExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) PutNotification(value interface{}) {
	if err := a.validatePutNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) PutPlannedLimit(value interface{}) {
	if err := a.validatePutPlannedLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlannedLimit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetAutoAdjustData() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoAdjustData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetBillingViewArn() {
	_jsii_.InvokeVoid(
		a,
		"resetBillingViewArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetCostFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetCostFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetCostTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetCostTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetFilterExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetLimitAmount() {
	_jsii_.InvokeVoid(
		a,
		"resetLimitAmount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetLimitUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetLimitUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetNotification() {
	_jsii_.InvokeVoid(
		a,
		"resetNotification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetPlannedLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetPlannedLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetTimePeriodEnd() {
	_jsii_.InvokeVoid(
		a,
		"resetTimePeriodEnd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) ResetTimePeriodStart() {
	_jsii_.InvokeVoid(
		a,
		"resetTimePeriodStart",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudget) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

