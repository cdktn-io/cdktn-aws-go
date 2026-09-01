package awsarczonalshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarczonalshift/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsarczonalshift/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arczonalshift_zonal_autoshift_configuration aws_arczonalshift_zonal_autoshift_configuration}.
// Experimental.
type AwsArczonalshiftZonalAutoshiftConfiguration interface {
	cdktn.TerraformResource
	// Experimental.
	AllowedWindows() *[]*string
	// Experimental.
	SetAllowedWindows(val *[]*string)
	// Experimental.
	AllowedWindowsInput() *[]*string
	// Experimental.
	BlockedDates() *[]*string
	// Experimental.
	SetBlockedDates(val *[]*string)
	// Experimental.
	BlockedDatesInput() *[]*string
	// Experimental.
	BlockedWindows() *[]*string
	// Experimental.
	SetBlockedWindows(val *[]*string)
	// Experimental.
	BlockedWindowsInput() *[]*string
	// Experimental.
	BlockingAlarms() AwsArczonalshiftZonalAutoshiftConfiguration_BlockingAlarmsPropertyList
	// Experimental.
	BlockingAlarmsInput() interface{}
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
	OutcomeAlarms() AwsArczonalshiftZonalAutoshiftConfiguration_OutcomeAlarmsPropertyList
	// Experimental.
	OutcomeAlarmsInput() interface{}
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
	ResourceArn() *string
	// Experimental.
	SetResourceArn(val *string)
	// Experimental.
	ResourceArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	ZonalAutoshiftStatus() *string
	// Experimental.
	SetZonalAutoshiftStatus(val *string)
	// Experimental.
	ZonalAutoshiftStatusInput() *string
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
	PutBlockingAlarms(value interface{})
	// Experimental.
	PutOutcomeAlarms(value interface{})
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
	ResetAllowedWindows()
	// Experimental.
	ResetBlockedDates()
	// Experimental.
	ResetBlockedWindows()
	// Experimental.
	ResetBlockingAlarms()
	// Experimental.
	ResetOutcomeAlarms()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
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

// The jsii proxy struct for AwsArczonalshiftZonalAutoshiftConfiguration
type jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) AllowedWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) AllowedWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) BlockedDates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedDates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) BlockedDatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedDatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) BlockedWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) BlockedWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) BlockingAlarms() AwsArczonalshiftZonalAutoshiftConfiguration_BlockingAlarmsPropertyList {
	var returns AwsArczonalshiftZonalAutoshiftConfiguration_BlockingAlarmsPropertyList
	_jsii_.Get(
		j,
		"blockingAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) BlockingAlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockingAlarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) OutcomeAlarms() AwsArczonalshiftZonalAutoshiftConfiguration_OutcomeAlarmsPropertyList {
	var returns AwsArczonalshiftZonalAutoshiftConfiguration_OutcomeAlarmsPropertyList
	_jsii_.Get(
		j,
		"outcomeAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) OutcomeAlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outcomeAlarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ZonalAutoshiftStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zonalAutoshiftStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ZonalAutoshiftStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zonalAutoshiftStatusInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arczonalshift_zonal_autoshift_configuration aws_arczonalshift_zonal_autoshift_configuration} Resource.
// Experimental.
func NewAwsArczonalshiftZonalAutoshiftConfiguration(scope constructs.Construct, id *string, config *AwsArczonalshiftZonalAutoshiftConfigurationConfig) AwsArczonalshiftZonalAutoshiftConfiguration {
	_init_.Initialize()

	if err := validateNewAwsArczonalshiftZonalAutoshiftConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration{}

	_jsii_.Create(
		"@cdktn/aws-arc-zonal-shift.AwsArczonalshiftZonalAutoshiftConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arczonalshift_zonal_autoshift_configuration aws_arczonalshift_zonal_autoshift_configuration} Resource.
// Experimental.
func NewAwsArczonalshiftZonalAutoshiftConfiguration_Override(a AwsArczonalshiftZonalAutoshiftConfiguration, scope constructs.Construct, id *string, config *AwsArczonalshiftZonalAutoshiftConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-zonal-shift.AwsArczonalshiftZonalAutoshiftConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetAllowedWindows(val *[]*string) {
	if err := j.validateSetAllowedWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedWindows",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetBlockedDates(val *[]*string) {
	if err := j.validateSetBlockedDatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedDates",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetBlockedWindows(val *[]*string) {
	if err := j.validateSetBlockedWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedWindows",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetResourceArn(val *string) {
	if err := j.validateSetResourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration)SetZonalAutoshiftStatus(val *string) {
	if err := j.validateSetZonalAutoshiftStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zonalAutoshiftStatus",
		val,
	)
}

// Generates CDKTN code for importing a AwsArczonalshiftZonalAutoshiftConfiguration resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsArczonalshiftZonalAutoshiftConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsArczonalshiftZonalAutoshiftConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-arc-zonal-shift.AwsArczonalshiftZonalAutoshiftConfiguration",
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
func AwsArczonalshiftZonalAutoshiftConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsArczonalshiftZonalAutoshiftConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-arc-zonal-shift.AwsArczonalshiftZonalAutoshiftConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsArczonalshiftZonalAutoshiftConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsArczonalshiftZonalAutoshiftConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-arc-zonal-shift.AwsArczonalshiftZonalAutoshiftConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsArczonalshiftZonalAutoshiftConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsArczonalshiftZonalAutoshiftConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-arc-zonal-shift.AwsArczonalshiftZonalAutoshiftConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsArczonalshiftZonalAutoshiftConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-arc-zonal-shift.AwsArczonalshiftZonalAutoshiftConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) PutBlockingAlarms(value interface{}) {
	if err := a.validatePutBlockingAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlockingAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) PutOutcomeAlarms(value interface{}) {
	if err := a.validatePutOutcomeAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutcomeAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResetAllowedWindows() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedWindows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResetBlockedDates() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockedDates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResetBlockedWindows() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockedWindows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResetBlockingAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockingAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResetOutcomeAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetOutcomeAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArczonalshiftZonalAutoshiftConfiguration) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

