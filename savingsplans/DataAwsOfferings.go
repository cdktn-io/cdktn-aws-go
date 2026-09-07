package savingsplans

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/savingsplans/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/savingsplans/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings aws_savingsplans_offerings}.
// Experimental.
type DataAwsOfferings interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	Currencies() *[]*string
	// Experimental.
	SetCurrencies(val *[]*string)
	// Experimental.
	CurrenciesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Descriptions() *[]*string
	// Experimental.
	SetDescriptions(val *[]*string)
	// Experimental.
	DescriptionsInput() *[]*string
	// Experimental.
	Durations() *[]*float64
	// Experimental.
	SetDurations(val *[]*float64)
	// Experimental.
	DurationsInput() *[]*float64
	// Experimental.
	Filter() DataAwsOfferings_FilterPropertyList
	// Experimental.
	FilterInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OfferingIds() *[]*string
	// Experimental.
	SetOfferingIds(val *[]*string)
	// Experimental.
	OfferingIdsInput() *[]*string
	// Experimental.
	Offerings() DataAwsOfferings_OfferingsPropertyList
	// Experimental.
	Operations() *[]*string
	// Experimental.
	SetOperations(val *[]*string)
	// Experimental.
	OperationsInput() *[]*string
	// Experimental.
	PaymentOptions() *[]*string
	// Experimental.
	SetPaymentOptions(val *[]*string)
	// Experimental.
	PaymentOptionsInput() *[]*string
	// Experimental.
	PlanTypes() *[]*string
	// Experimental.
	SetPlanTypes(val *[]*string)
	// Experimental.
	PlanTypesInput() *[]*string
	// Experimental.
	ProductType() *string
	// Experimental.
	SetProductType(val *string)
	// Experimental.
	ProductTypeInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	ServiceCodes() *[]*string
	// Experimental.
	SetServiceCodes(val *[]*string)
	// Experimental.
	ServiceCodesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	UsageTypes() *[]*string
	// Experimental.
	SetUsageTypes(val *[]*string)
	// Experimental.
	UsageTypesInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutFilter(value interface{})
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
	ResetCurrencies()
	// Experimental.
	ResetDescriptions()
	// Experimental.
	ResetDurations()
	// Experimental.
	ResetFilter()
	// Experimental.
	ResetOfferingIds()
	// Experimental.
	ResetOperations()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPaymentOptions()
	// Experimental.
	ResetPlanTypes()
	// Experimental.
	ResetProductType()
	// Experimental.
	ResetServiceCodes()
	// Experimental.
	ResetUsageTypes()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsOfferings
type jsiiProxy_DataAwsOfferings struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsOfferings) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Currencies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"currencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) CurrenciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"currenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Descriptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"descriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) DescriptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"descriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Durations() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"durations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) DurationsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"durationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Filter() DataAwsOfferings_FilterPropertyList {
	var returns DataAwsOfferings_FilterPropertyList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) OfferingIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"offeringIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) OfferingIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"offeringIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Offerings() DataAwsOfferings_OfferingsPropertyList {
	var returns DataAwsOfferings_OfferingsPropertyList
	_jsii_.Get(
		j,
		"offerings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Operations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) OperationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) PaymentOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paymentOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) PaymentOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paymentOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) PlanTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"planTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) PlanTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"planTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) ProductType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) ProductTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) ServiceCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) ServiceCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) UsageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOfferings) UsageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usageTypesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings aws_savingsplans_offerings} Data Source.
// Experimental.
func NewDataAwsOfferings(scope constructs.Construct, id *string, config *DataAwsOfferingsConfig) DataAwsOfferings {
	_init_.Initialize()

	if err := validateNewDataAwsOfferingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOfferings{}

	_jsii_.Create(
		"@cdktn/aws-savings-plans.DataAwsOfferings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings aws_savingsplans_offerings} Data Source.
// Experimental.
func NewDataAwsOfferings_Override(d DataAwsOfferings, scope constructs.Construct, id *string, config *DataAwsOfferingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-savings-plans.DataAwsOfferings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetCurrencies(val *[]*string) {
	if err := j.validateSetCurrenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currencies",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetDescriptions(val *[]*string) {
	if err := j.validateSetDescriptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"descriptions",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetDurations(val *[]*float64) {
	if err := j.validateSetDurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durations",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetOfferingIds(val *[]*string) {
	if err := j.validateSetOfferingIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offeringIds",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetOperations(val *[]*string) {
	if err := j.validateSetOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operations",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetPaymentOptions(val *[]*string) {
	if err := j.validateSetPaymentOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paymentOptions",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetPlanTypes(val *[]*string) {
	if err := j.validateSetPlanTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planTypes",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetProductType(val *string) {
	if err := j.validateSetProductTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productType",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetServiceCodes(val *[]*string) {
	if err := j.validateSetServiceCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceCodes",
		val,
	)
}

func (j *jsiiProxy_DataAwsOfferings)SetUsageTypes(val *[]*string) {
	if err := j.validateSetUsageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageTypes",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsOfferings resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataAwsOfferings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsOfferings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-savings-plans.DataAwsOfferings",
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
func DataAwsOfferings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOfferings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-savings-plans.DataAwsOfferings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOfferings_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOfferings_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-savings-plans.DataAwsOfferings",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOfferings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOfferings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-savings-plans.DataAwsOfferings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOfferings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-savings-plans.DataAwsOfferings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsOfferings) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOfferings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsOfferings) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsOfferings) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetCurrencies() {
	_jsii_.InvokeVoid(
		d,
		"resetCurrencies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetDescriptions() {
	_jsii_.InvokeVoid(
		d,
		"resetDescriptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetDurations() {
	_jsii_.InvokeVoid(
		d,
		"resetDurations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetOfferingIds() {
	_jsii_.InvokeVoid(
		d,
		"resetOfferingIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetOperations() {
	_jsii_.InvokeVoid(
		d,
		"resetOperations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetPaymentOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetPaymentOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetPlanTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetPlanTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetProductType() {
	_jsii_.InvokeVoid(
		d,
		"resetProductType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetServiceCodes() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceCodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) ResetUsageTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetUsageTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOfferings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOfferings) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

