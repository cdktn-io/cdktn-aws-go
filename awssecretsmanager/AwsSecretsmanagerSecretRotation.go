package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecretsmanager/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awssecretsmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation aws_secretsmanager_secret_rotation}.
// Experimental.
type AwsSecretsmanagerSecretRotation interface {
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
	ExternalSecretRotationMetadata() AwsSecretsmanagerSecretRotation_ExternalSecretRotationMetadataPropertyList
	// Experimental.
	ExternalSecretRotationMetadataInput() interface{}
	// Experimental.
	ExternalSecretRotationRoleArn() *string
	// Experimental.
	SetExternalSecretRotationRoleArn(val *string)
	// Experimental.
	ExternalSecretRotationRoleArnInput() *string
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
	RotateImmediately() interface{}
	// Experimental.
	SetRotateImmediately(val interface{})
	// Experimental.
	RotateImmediatelyInput() interface{}
	// Experimental.
	RotationEnabled() interface{}
	// Experimental.
	SetRotationEnabled(val interface{})
	// Experimental.
	RotationEnabledInput() interface{}
	// Experimental.
	RotationLambdaArn() *string
	// Experimental.
	SetRotationLambdaArn(val *string)
	// Experimental.
	RotationLambdaArnInput() *string
	// Experimental.
	RotationRules() AwsSecretsmanagerSecretRotation_RotationRulesPropertyOutputReference
	// Experimental.
	RotationRulesInput() *AwsSecretsmanagerSecretRotation_RotationRulesProperty
	// Experimental.
	SecretId() *string
	// Experimental.
	SetSecretId(val *string)
	// Experimental.
	SecretIdInput() *string
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
	PutExternalSecretRotationMetadata(value interface{})
	// Experimental.
	PutRotationRules(value *AwsSecretsmanagerSecretRotation_RotationRulesProperty)
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
	ResetExternalSecretRotationMetadata()
	// Experimental.
	ResetExternalSecretRotationRoleArn()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRotateImmediately()
	// Experimental.
	ResetRotationEnabled()
	// Experimental.
	ResetRotationLambdaArn()
	// Experimental.
	ResetRotationRules()
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

// The jsii proxy struct for AwsSecretsmanagerSecretRotation
type jsiiProxy_AwsSecretsmanagerSecretRotation struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) ExternalSecretRotationMetadata() AwsSecretsmanagerSecretRotation_ExternalSecretRotationMetadataPropertyList {
	var returns AwsSecretsmanagerSecretRotation_ExternalSecretRotationMetadataPropertyList
	_jsii_.Get(
		j,
		"externalSecretRotationMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) ExternalSecretRotationMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalSecretRotationMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) ExternalSecretRotationRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalSecretRotationRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) ExternalSecretRotationRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalSecretRotationRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotateImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotateImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotationLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotationLambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationLambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotationRules() AwsSecretsmanagerSecretRotation_RotationRulesPropertyOutputReference {
	var returns AwsSecretsmanagerSecretRotation_RotationRulesPropertyOutputReference
	_jsii_.Get(
		j,
		"rotationRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) RotationRulesInput() *AwsSecretsmanagerSecretRotation_RotationRulesProperty {
	var returns *AwsSecretsmanagerSecretRotation_RotationRulesProperty
	_jsii_.Get(
		j,
		"rotationRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) SecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation aws_secretsmanager_secret_rotation} Resource.
// Experimental.
func NewAwsSecretsmanagerSecretRotation(scope constructs.Construct, id *string, config *AwsSecretsmanagerSecretRotationConfig) AwsSecretsmanagerSecretRotation {
	_init_.Initialize()

	if err := validateNewAwsSecretsmanagerSecretRotationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSecretsmanagerSecretRotation{}

	_jsii_.Create(
		"@cdktn/aws-secrets-manager.AwsSecretsmanagerSecretRotation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation aws_secretsmanager_secret_rotation} Resource.
// Experimental.
func NewAwsSecretsmanagerSecretRotation_Override(a AwsSecretsmanagerSecretRotation, scope constructs.Construct, id *string, config *AwsSecretsmanagerSecretRotationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-secrets-manager.AwsSecretsmanagerSecretRotation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetExternalSecretRotationRoleArn(val *string) {
	if err := j.validateSetExternalSecretRotationRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalSecretRotationRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetRotateImmediately(val interface{}) {
	if err := j.validateSetRotateImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotateImmediately",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetRotationEnabled(val interface{}) {
	if err := j.validateSetRotationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetRotationLambdaArn(val *string) {
	if err := j.validateSetRotationLambdaArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationLambdaArn",
		val,
	)
}

func (j *jsiiProxy_AwsSecretsmanagerSecretRotation)SetSecretId(val *string) {
	if err := j.validateSetSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

// Generates CDKTN code for importing a AwsSecretsmanagerSecretRotation resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsSecretsmanagerSecretRotation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsSecretsmanagerSecretRotation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.AwsSecretsmanagerSecretRotation",
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
func AwsSecretsmanagerSecretRotation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSecretsmanagerSecretRotation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.AwsSecretsmanagerSecretRotation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSecretsmanagerSecretRotation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSecretsmanagerSecretRotation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.AwsSecretsmanagerSecretRotation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSecretsmanagerSecretRotation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSecretsmanagerSecretRotation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.AwsSecretsmanagerSecretRotation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSecretsmanagerSecretRotation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-secrets-manager.AwsSecretsmanagerSecretRotation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) PutExternalSecretRotationMetadata(value interface{}) {
	if err := a.validatePutExternalSecretRotationMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExternalSecretRotationMetadata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) PutRotationRules(value *AwsSecretsmanagerSecretRotation_RotationRulesProperty) {
	if err := a.validatePutRotationRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRotationRules",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetExternalSecretRotationMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalSecretRotationMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetExternalSecretRotationRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalSecretRotationRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetRotateImmediately() {
	_jsii_.InvokeVoid(
		a,
		"resetRotateImmediately",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetRotationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetRotationLambdaArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationLambdaArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ResetRotationRules() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationRules",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretsmanagerSecretRotation) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

