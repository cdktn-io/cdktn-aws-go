package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconfig/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule aws_config_organization_managed_rule}.
// Experimental.
type AwsConfigOrganizationManagedRule interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	ExcludedAccounts() *[]*string
	// Experimental.
	SetExcludedAccounts(val *[]*string)
	// Experimental.
	ExcludedAccountsInput() *[]*string
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
	InputParameters() *string
	// Experimental.
	SetInputParameters(val *string)
	// Experimental.
	InputParametersInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaximumExecutionFrequency() *string
	// Experimental.
	SetMaximumExecutionFrequency(val *string)
	// Experimental.
	MaximumExecutionFrequencyInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
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
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ResourceIdScope() *string
	// Experimental.
	SetResourceIdScope(val *string)
	// Experimental.
	ResourceIdScopeInput() *string
	// Experimental.
	ResourceTypesScope() *[]*string
	// Experimental.
	SetResourceTypesScope(val *[]*string)
	// Experimental.
	ResourceTypesScopeInput() *[]*string
	// Experimental.
	RuleIdentifier() *string
	// Experimental.
	SetRuleIdentifier(val *string)
	// Experimental.
	RuleIdentifierInput() *string
	// Experimental.
	TagKeyScope() *string
	// Experimental.
	SetTagKeyScope(val *string)
	// Experimental.
	TagKeyScopeInput() *string
	// Experimental.
	TagValueScope() *string
	// Experimental.
	SetTagValueScope(val *string)
	// Experimental.
	TagValueScopeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsConfigOrganizationManagedRule_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
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
	PutTimeouts(value *AwsConfigOrganizationManagedRule_TimeoutsProperty)
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
	ResetDescription()
	// Experimental.
	ResetExcludedAccounts()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInputParameters()
	// Experimental.
	ResetMaximumExecutionFrequency()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetResourceIdScope()
	// Experimental.
	ResetResourceTypesScope()
	// Experimental.
	ResetTagKeyScope()
	// Experimental.
	ResetTagValueScope()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for AwsConfigOrganizationManagedRule
type jsiiProxy_AwsConfigOrganizationManagedRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ExcludedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ExcludedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) InputParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) InputParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) MaximumExecutionFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) MaximumExecutionFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ResourceIdScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ResourceIdScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ResourceTypesScope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) ResourceTypesScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) RuleIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) RuleIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TagKeyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TagKeyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TagValueScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TagValueScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) Timeouts() AwsConfigOrganizationManagedRule_TimeoutsPropertyOutputReference {
	var returns AwsConfigOrganizationManagedRule_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule aws_config_organization_managed_rule} Resource.
// Experimental.
func NewAwsConfigOrganizationManagedRule(scope constructs.Construct, id *string, config *AwsConfigOrganizationManagedRuleConfig) AwsConfigOrganizationManagedRule {
	_init_.Initialize()

	if err := validateNewAwsConfigOrganizationManagedRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConfigOrganizationManagedRule{}

	_jsii_.Create(
		"@cdktn/aws-config.AwsConfigOrganizationManagedRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule aws_config_organization_managed_rule} Resource.
// Experimental.
func NewAwsConfigOrganizationManagedRule_Override(a AwsConfigOrganizationManagedRule, scope constructs.Construct, id *string, config *AwsConfigOrganizationManagedRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-config.AwsConfigOrganizationManagedRule",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetExcludedAccounts(val *[]*string) {
	if err := j.validateSetExcludedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetInputParameters(val *string) {
	if err := j.validateSetInputParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetMaximumExecutionFrequency(val *string) {
	if err := j.validateSetMaximumExecutionFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetResourceIdScope(val *string) {
	if err := j.validateSetResourceIdScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdScope",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetResourceTypesScope(val *[]*string) {
	if err := j.validateSetResourceTypesScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypesScope",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetRuleIdentifier(val *string) {
	if err := j.validateSetRuleIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetTagKeyScope(val *string) {
	if err := j.validateSetTagKeyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKeyScope",
		val,
	)
}

func (j *jsiiProxy_AwsConfigOrganizationManagedRule)SetTagValueScope(val *string) {
	if err := j.validateSetTagValueScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValueScope",
		val,
	)
}

// Generates CDKTN code for importing a AwsConfigOrganizationManagedRule resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsConfigOrganizationManagedRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsConfigOrganizationManagedRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-config.AwsConfigOrganizationManagedRule",
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
func AwsConfigOrganizationManagedRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsConfigOrganizationManagedRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-config.AwsConfigOrganizationManagedRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsConfigOrganizationManagedRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsConfigOrganizationManagedRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-config.AwsConfigOrganizationManagedRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsConfigOrganizationManagedRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsConfigOrganizationManagedRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-config.AwsConfigOrganizationManagedRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsConfigOrganizationManagedRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-config.AwsConfigOrganizationManagedRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) PutTimeouts(value *AwsConfigOrganizationManagedRule_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetExcludedAccounts() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedAccounts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetInputParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetInputParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetMaximumExecutionFrequency() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumExecutionFrequency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetResourceIdScope() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceIdScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetResourceTypesScope() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTypesScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetTagKeyScope() {
	_jsii_.InvokeVoid(
		a,
		"resetTagKeyScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetTagValueScope() {
	_jsii_.InvokeVoid(
		a,
		"resetTagValueScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigOrganizationManagedRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

