package awscomputeoptimizer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscomputeoptimizer/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscomputeoptimizer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences aws_computeoptimizer_recommendation_preferences}.
// Experimental.
type AwsComputeoptimizerRecommendationPreferences interface {
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
	EnhancedInfrastructureMetrics() *string
	// Experimental.
	SetEnhancedInfrastructureMetrics(val *string)
	// Experimental.
	EnhancedInfrastructureMetricsInput() *string
	// Experimental.
	ExternalMetricsPreference() AwsComputeoptimizerRecommendationPreferences_ExternalMetricsPreferencePropertyList
	// Experimental.
	ExternalMetricsPreferenceInput() interface{}
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
	InferredWorkloadTypes() *string
	// Experimental.
	SetInferredWorkloadTypes(val *string)
	// Experimental.
	InferredWorkloadTypesInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LookBackPeriod() *string
	// Experimental.
	SetLookBackPeriod(val *string)
	// Experimental.
	LookBackPeriodInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PreferredResource() AwsComputeoptimizerRecommendationPreferences_PreferredResourcePropertyList
	// Experimental.
	PreferredResourceInput() interface{}
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
	ResourceType() *string
	// Experimental.
	SetResourceType(val *string)
	// Experimental.
	ResourceTypeInput() *string
	// Experimental.
	SavingsEstimationMode() *string
	// Experimental.
	SetSavingsEstimationMode(val *string)
	// Experimental.
	SavingsEstimationModeInput() *string
	// Experimental.
	Scope() AwsComputeoptimizerRecommendationPreferences_ScopePropertyList
	// Experimental.
	ScopeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	UtilizationPreference() AwsComputeoptimizerRecommendationPreferences_UtilizationPreferencePropertyList
	// Experimental.
	UtilizationPreferenceInput() interface{}
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
	PutExternalMetricsPreference(value interface{})
	// Experimental.
	PutPreferredResource(value interface{})
	// Experimental.
	PutScope(value interface{})
	// Experimental.
	PutUtilizationPreference(value interface{})
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
	ResetEnhancedInfrastructureMetrics()
	// Experimental.
	ResetExternalMetricsPreference()
	// Experimental.
	ResetInferredWorkloadTypes()
	// Experimental.
	ResetLookBackPeriod()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPreferredResource()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSavingsEstimationMode()
	// Experimental.
	ResetScope()
	// Experimental.
	ResetUtilizationPreference()
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

// The jsii proxy struct for AwsComputeoptimizerRecommendationPreferences
type jsiiProxy_AwsComputeoptimizerRecommendationPreferences struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) EnhancedInfrastructureMetrics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedInfrastructureMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) EnhancedInfrastructureMetricsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedInfrastructureMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ExternalMetricsPreference() AwsComputeoptimizerRecommendationPreferences_ExternalMetricsPreferencePropertyList {
	var returns AwsComputeoptimizerRecommendationPreferences_ExternalMetricsPreferencePropertyList
	_jsii_.Get(
		j,
		"externalMetricsPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ExternalMetricsPreferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalMetricsPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) InferredWorkloadTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredWorkloadTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) InferredWorkloadTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredWorkloadTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) LookBackPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookBackPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) LookBackPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookBackPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) PreferredResource() AwsComputeoptimizerRecommendationPreferences_PreferredResourcePropertyList {
	var returns AwsComputeoptimizerRecommendationPreferences_PreferredResourcePropertyList
	_jsii_.Get(
		j,
		"preferredResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) PreferredResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferredResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) SavingsEstimationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsEstimationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) SavingsEstimationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsEstimationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) Scope() AwsComputeoptimizerRecommendationPreferences_ScopePropertyList {
	var returns AwsComputeoptimizerRecommendationPreferences_ScopePropertyList
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) UtilizationPreference() AwsComputeoptimizerRecommendationPreferences_UtilizationPreferencePropertyList {
	var returns AwsComputeoptimizerRecommendationPreferences_UtilizationPreferencePropertyList
	_jsii_.Get(
		j,
		"utilizationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) UtilizationPreferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"utilizationPreferenceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences aws_computeoptimizer_recommendation_preferences} Resource.
// Experimental.
func NewAwsComputeoptimizerRecommendationPreferences(scope constructs.Construct, id *string, config *AwsComputeoptimizerRecommendationPreferencesConfig) AwsComputeoptimizerRecommendationPreferences {
	_init_.Initialize()

	if err := validateNewAwsComputeoptimizerRecommendationPreferencesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsComputeoptimizerRecommendationPreferences{}

	_jsii_.Create(
		"@cdktn/aws-compute-optimizer.AwsComputeoptimizerRecommendationPreferences",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/computeoptimizer_recommendation_preferences aws_computeoptimizer_recommendation_preferences} Resource.
// Experimental.
func NewAwsComputeoptimizerRecommendationPreferences_Override(a AwsComputeoptimizerRecommendationPreferences, scope constructs.Construct, id *string, config *AwsComputeoptimizerRecommendationPreferencesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-compute-optimizer.AwsComputeoptimizerRecommendationPreferences",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetEnhancedInfrastructureMetrics(val *string) {
	if err := j.validateSetEnhancedInfrastructureMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedInfrastructureMetrics",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetInferredWorkloadTypes(val *string) {
	if err := j.validateSetInferredWorkloadTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferredWorkloadTypes",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetLookBackPeriod(val *string) {
	if err := j.validateSetLookBackPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookBackPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_AwsComputeoptimizerRecommendationPreferences)SetSavingsEstimationMode(val *string) {
	if err := j.validateSetSavingsEstimationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"savingsEstimationMode",
		val,
	)
}

// Generates CDKTN code for importing a AwsComputeoptimizerRecommendationPreferences resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsComputeoptimizerRecommendationPreferences_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsComputeoptimizerRecommendationPreferences_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-compute-optimizer.AwsComputeoptimizerRecommendationPreferences",
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
func AwsComputeoptimizerRecommendationPreferences_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsComputeoptimizerRecommendationPreferences_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-compute-optimizer.AwsComputeoptimizerRecommendationPreferences",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsComputeoptimizerRecommendationPreferences_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsComputeoptimizerRecommendationPreferences_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-compute-optimizer.AwsComputeoptimizerRecommendationPreferences",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsComputeoptimizerRecommendationPreferences_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsComputeoptimizerRecommendationPreferences_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-compute-optimizer.AwsComputeoptimizerRecommendationPreferences",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsComputeoptimizerRecommendationPreferences_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-compute-optimizer.AwsComputeoptimizerRecommendationPreferences",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) PutExternalMetricsPreference(value interface{}) {
	if err := a.validatePutExternalMetricsPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExternalMetricsPreference",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) PutPreferredResource(value interface{}) {
	if err := a.validatePutPreferredResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPreferredResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) PutScope(value interface{}) {
	if err := a.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScope",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) PutUtilizationPreference(value interface{}) {
	if err := a.validatePutUtilizationPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUtilizationPreference",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetEnhancedInfrastructureMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetEnhancedInfrastructureMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetExternalMetricsPreference() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalMetricsPreference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetInferredWorkloadTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetInferredWorkloadTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetLookBackPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetLookBackPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetPreferredResource() {
	_jsii_.InvokeVoid(
		a,
		"resetPreferredResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetSavingsEstimationMode() {
	_jsii_.InvokeVoid(
		a,
		"resetSavingsEstimationMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ResetUtilizationPreference() {
	_jsii_.InvokeVoid(
		a,
		"resetUtilizationPreference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComputeoptimizerRecommendationPreferences) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

