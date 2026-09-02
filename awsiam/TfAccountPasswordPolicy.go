package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiam/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsiam/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy aws_iam_account_password_policy}.
// Experimental.
type TfAccountPasswordPolicy interface {
	cdktn.TerraformResource
	// Experimental.
	AllowUsersToChangePassword() interface{}
	// Experimental.
	SetAllowUsersToChangePassword(val interface{})
	// Experimental.
	AllowUsersToChangePasswordInput() interface{}
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
	ExpirePasswords() cdktn.IResolvable
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HardExpiry() interface{}
	// Experimental.
	SetHardExpiry(val interface{})
	// Experimental.
	HardExpiryInput() interface{}
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
	MaxPasswordAge() *float64
	// Experimental.
	SetMaxPasswordAge(val *float64)
	// Experimental.
	MaxPasswordAgeInput() *float64
	// Experimental.
	MinimumPasswordLength() *float64
	// Experimental.
	SetMinimumPasswordLength(val *float64)
	// Experimental.
	MinimumPasswordLengthInput() *float64
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PasswordReusePrevention() *float64
	// Experimental.
	SetPasswordReusePrevention(val *float64)
	// Experimental.
	PasswordReusePreventionInput() *float64
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
	RequireLowercaseCharacters() interface{}
	// Experimental.
	SetRequireLowercaseCharacters(val interface{})
	// Experimental.
	RequireLowercaseCharactersInput() interface{}
	// Experimental.
	RequireNumbers() interface{}
	// Experimental.
	SetRequireNumbers(val interface{})
	// Experimental.
	RequireNumbersInput() interface{}
	// Experimental.
	RequireSymbols() interface{}
	// Experimental.
	SetRequireSymbols(val interface{})
	// Experimental.
	RequireSymbolsInput() interface{}
	// Experimental.
	RequireUppercaseCharacters() interface{}
	// Experimental.
	SetRequireUppercaseCharacters(val interface{})
	// Experimental.
	RequireUppercaseCharactersInput() interface{}
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
	ResetAllowUsersToChangePassword()
	// Experimental.
	ResetHardExpiry()
	// Experimental.
	ResetId()
	// Experimental.
	ResetMaxPasswordAge()
	// Experimental.
	ResetMinimumPasswordLength()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPasswordReusePrevention()
	// Experimental.
	ResetRequireLowercaseCharacters()
	// Experimental.
	ResetRequireNumbers()
	// Experimental.
	ResetRequireSymbols()
	// Experimental.
	ResetRequireUppercaseCharacters()
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

// The jsii proxy struct for TfAccountPasswordPolicy
type jsiiProxy_TfAccountPasswordPolicy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfAccountPasswordPolicy) AllowUsersToChangePassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUsersToChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) AllowUsersToChangePasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUsersToChangePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) ExpirePasswords() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"expirePasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) HardExpiry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) HardExpiryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) MaxPasswordAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPasswordAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) MaxPasswordAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPasswordAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) MinimumPasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) MinimumPasswordLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) PasswordReusePrevention() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordReusePrevention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) PasswordReusePreventionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordReusePreventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireLowercaseCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireLowercaseCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireSymbols() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireSymbolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireUppercaseCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) RequireUppercaseCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAccountPasswordPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy aws_iam_account_password_policy} Resource.
// Experimental.
func NewTfAccountPasswordPolicy(scope constructs.Construct, id *string, config *TfAccountPasswordPolicyConfig) TfAccountPasswordPolicy {
	_init_.Initialize()

	if err := validateNewTfAccountPasswordPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAccountPasswordPolicy{}

	_jsii_.Create(
		"@cdktn/aws-iam.TfAccountPasswordPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iam_account_password_policy aws_iam_account_password_policy} Resource.
// Experimental.
func NewTfAccountPasswordPolicy_Override(t TfAccountPasswordPolicy, scope constructs.Construct, id *string, config *TfAccountPasswordPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iam.TfAccountPasswordPolicy",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetAllowUsersToChangePassword(val interface{}) {
	if err := j.validateSetAllowUsersToChangePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUsersToChangePassword",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetHardExpiry(val interface{}) {
	if err := j.validateSetHardExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardExpiry",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetMaxPasswordAge(val *float64) {
	if err := j.validateSetMaxPasswordAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPasswordAge",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetMinimumPasswordLength(val *float64) {
	if err := j.validateSetMinimumPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumPasswordLength",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetPasswordReusePrevention(val *float64) {
	if err := j.validateSetPasswordReusePreventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordReusePrevention",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetRequireLowercaseCharacters(val interface{}) {
	if err := j.validateSetRequireLowercaseCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLowercaseCharacters",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetRequireNumbers(val interface{}) {
	if err := j.validateSetRequireNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireNumbers",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetRequireSymbols(val interface{}) {
	if err := j.validateSetRequireSymbolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSymbols",
		val,
	)
}

func (j *jsiiProxy_TfAccountPasswordPolicy)SetRequireUppercaseCharacters(val interface{}) {
	if err := j.validateSetRequireUppercaseCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireUppercaseCharacters",
		val,
	)
}

// Generates CDKTN code for importing a TfAccountPasswordPolicy resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfAccountPasswordPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfAccountPasswordPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.TfAccountPasswordPolicy",
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
func TfAccountPasswordPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAccountPasswordPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.TfAccountPasswordPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfAccountPasswordPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAccountPasswordPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.TfAccountPasswordPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfAccountPasswordPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAccountPasswordPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.TfAccountPasswordPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfAccountPasswordPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-iam.TfAccountPasswordPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfAccountPasswordPolicy) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetAllowUsersToChangePassword() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowUsersToChangePassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetHardExpiry() {
	_jsii_.InvokeVoid(
		t,
		"resetHardExpiry",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetMaxPasswordAge() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxPasswordAge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetMinimumPasswordLength() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumPasswordLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetPasswordReusePrevention() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordReusePrevention",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetRequireLowercaseCharacters() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireLowercaseCharacters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetRequireNumbers() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireNumbers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetRequireSymbols() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireSymbols",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ResetRequireUppercaseCharacters() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireUppercaseCharacters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAccountPasswordPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAccountPasswordPolicy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

