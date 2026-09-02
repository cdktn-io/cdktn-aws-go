package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecretsmanager/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awssecretsmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/secretsmanager_random_password aws_secretsmanager_random_password}.
// Experimental.
type DataTfRandomPassword interface {
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ExcludeCharacters() *string
	// Experimental.
	SetExcludeCharacters(val *string)
	// Experimental.
	ExcludeCharactersInput() *string
	// Experimental.
	ExcludeLowercase() interface{}
	// Experimental.
	SetExcludeLowercase(val interface{})
	// Experimental.
	ExcludeLowercaseInput() interface{}
	// Experimental.
	ExcludeNumbers() interface{}
	// Experimental.
	SetExcludeNumbers(val interface{})
	// Experimental.
	ExcludeNumbersInput() interface{}
	// Experimental.
	ExcludePunctuation() interface{}
	// Experimental.
	SetExcludePunctuation(val interface{})
	// Experimental.
	ExcludePunctuationInput() interface{}
	// Experimental.
	ExcludeUppercase() interface{}
	// Experimental.
	SetExcludeUppercase(val interface{})
	// Experimental.
	ExcludeUppercaseInput() interface{}
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
	IncludeSpace() interface{}
	// Experimental.
	SetIncludeSpace(val interface{})
	// Experimental.
	IncludeSpaceInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PasswordLength() *float64
	// Experimental.
	SetPasswordLength(val *float64)
	// Experimental.
	PasswordLengthInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RandomPassword() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RequireEachIncludedType() interface{}
	// Experimental.
	SetRequireEachIncludedType(val interface{})
	// Experimental.
	RequireEachIncludedTypeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetExcludeCharacters()
	// Experimental.
	ResetExcludeLowercase()
	// Experimental.
	ResetExcludeNumbers()
	// Experimental.
	ResetExcludePunctuation()
	// Experimental.
	ResetExcludeUppercase()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIncludeSpace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPasswordLength()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRequireEachIncludedType()
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

// The jsii proxy struct for DataTfRandomPassword
type jsiiProxy_DataTfRandomPassword struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataTfRandomPassword) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludePunctuation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludePunctuationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ExcludeUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) IncludeSpace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) IncludeSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) PasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) PasswordLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) RandomPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"randomPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) RequireEachIncludedType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) RequireEachIncludedTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRandomPassword) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/secretsmanager_random_password aws_secretsmanager_random_password} Data Source.
// Experimental.
func NewDataTfRandomPassword(scope constructs.Construct, id *string, config *DataTfRandomPasswordConfig) DataTfRandomPassword {
	_init_.Initialize()

	if err := validateNewDataTfRandomPasswordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfRandomPassword{}

	_jsii_.Create(
		"@cdktn/aws-secrets-manager.DataTfRandomPassword",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/secretsmanager_random_password aws_secretsmanager_random_password} Data Source.
// Experimental.
func NewDataTfRandomPassword_Override(d DataTfRandomPassword, scope constructs.Construct, id *string, config *DataTfRandomPasswordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-secrets-manager.DataTfRandomPassword",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetExcludeCharacters(val *string) {
	if err := j.validateSetExcludeCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCharacters",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetExcludeLowercase(val interface{}) {
	if err := j.validateSetExcludeLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeLowercase",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetExcludeNumbers(val interface{}) {
	if err := j.validateSetExcludeNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeNumbers",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetExcludePunctuation(val interface{}) {
	if err := j.validateSetExcludePunctuationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePunctuation",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetExcludeUppercase(val interface{}) {
	if err := j.validateSetExcludeUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeUppercase",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetIncludeSpace(val interface{}) {
	if err := j.validateSetIncludeSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSpace",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetPasswordLength(val *float64) {
	if err := j.validateSetPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLength",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataTfRandomPassword)SetRequireEachIncludedType(val interface{}) {
	if err := j.validateSetRequireEachIncludedTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireEachIncludedType",
		val,
	)
}

// Generates CDKTN code for importing a DataTfRandomPassword resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataTfRandomPassword_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfRandomPassword_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.DataTfRandomPassword",
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
func DataTfRandomPassword_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfRandomPassword_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.DataTfRandomPassword",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfRandomPassword_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfRandomPassword_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.DataTfRandomPassword",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfRandomPassword_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfRandomPassword_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-secrets-manager.DataTfRandomPassword",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfRandomPassword_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-secrets-manager.DataTfRandomPassword",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfRandomPassword) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfRandomPassword) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfRandomPassword) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfRandomPassword) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfRandomPassword) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfRandomPassword) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfRandomPassword) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfRandomPassword) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfRandomPassword) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfRandomPassword) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfRandomPassword) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfRandomPassword) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfRandomPassword) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetExcludeCharacters() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeCharacters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetExcludeLowercase() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeLowercase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetExcludeNumbers() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeNumbers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetExcludePunctuation() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludePunctuation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetExcludeUppercase() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeUppercase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetIncludeSpace() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeSpace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetPasswordLength() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) ResetRequireEachIncludedType() {
	_jsii_.InvokeVoid(
		d,
		"resetRequireEachIncludedType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfRandomPassword) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRandomPassword) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRandomPassword) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRandomPassword) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRandomPassword) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRandomPassword) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRandomPassword) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

