package iam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/iam/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/iam/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_principal_policy_simulation aws_iam_principal_policy_simulation}.
// Experimental.
type DataAwsPrincipalPolicySimulation interface {
	cdktn.TerraformDataSource
	// Experimental.
	ActionNames() *[]*string
	// Experimental.
	SetActionNames(val *[]*string)
	// Experimental.
	ActionNamesInput() *[]*string
	// Experimental.
	AdditionalPoliciesJson() *[]*string
	// Experimental.
	SetAdditionalPoliciesJson(val *[]*string)
	// Experimental.
	AdditionalPoliciesJsonInput() *[]*string
	// Experimental.
	AllAllowed() cdktn.IResolvable
	// Experimental.
	CallerArn() *string
	// Experimental.
	SetCallerArn(val *string)
	// Experimental.
	CallerArnInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Context() DataAwsPrincipalPolicySimulation_ContextPropertyList
	// Experimental.
	ContextInput() interface{}
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
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PermissionsBoundaryPoliciesJson() *[]*string
	// Experimental.
	SetPermissionsBoundaryPoliciesJson(val *[]*string)
	// Experimental.
	PermissionsBoundaryPoliciesJsonInput() *[]*string
	// Experimental.
	PolicySourceArn() *string
	// Experimental.
	SetPolicySourceArn(val *string)
	// Experimental.
	PolicySourceArnInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	ResourceArns() *[]*string
	// Experimental.
	SetResourceArns(val *[]*string)
	// Experimental.
	ResourceArnsInput() *[]*string
	// Experimental.
	ResourceHandlingOption() *string
	// Experimental.
	SetResourceHandlingOption(val *string)
	// Experimental.
	ResourceHandlingOptionInput() *string
	// Experimental.
	ResourceOwnerAccountId() *string
	// Experimental.
	SetResourceOwnerAccountId(val *string)
	// Experimental.
	ResourceOwnerAccountIdInput() *string
	// Experimental.
	ResourcePolicyJson() *string
	// Experimental.
	SetResourcePolicyJson(val *string)
	// Experimental.
	ResourcePolicyJsonInput() *string
	// Experimental.
	Results() DataAwsPrincipalPolicySimulation_ResultsPropertyList
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
	// Experimental.
	PutContext(value interface{})
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
	ResetAdditionalPoliciesJson()
	// Experimental.
	ResetCallerArn()
	// Experimental.
	ResetContext()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPermissionsBoundaryPoliciesJson()
	// Experimental.
	ResetResourceArns()
	// Experimental.
	ResetResourceHandlingOption()
	// Experimental.
	ResetResourceOwnerAccountId()
	// Experimental.
	ResetResourcePolicyJson()
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

// The jsii proxy struct for DataAwsPrincipalPolicySimulation
type jsiiProxy_DataAwsPrincipalPolicySimulation struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ActionNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ActionNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) AdditionalPoliciesJson() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPoliciesJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) AdditionalPoliciesJsonInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPoliciesJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) AllAllowed() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"allAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) CallerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) CallerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Context() DataAwsPrincipalPolicySimulation_ContextPropertyList {
	var returns DataAwsPrincipalPolicySimulation_ContextPropertyList
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) PermissionsBoundaryPoliciesJson() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsBoundaryPoliciesJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) PermissionsBoundaryPoliciesJsonInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsBoundaryPoliciesJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) PolicySourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policySourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) PolicySourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policySourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourceArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourceHandlingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceHandlingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourceHandlingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceHandlingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourceOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourceOwnerAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourcePolicyJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePolicyJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) ResourcePolicyJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePolicyJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) Results() DataAwsPrincipalPolicySimulation_ResultsPropertyList {
	var returns DataAwsPrincipalPolicySimulation_ResultsPropertyList
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_principal_policy_simulation aws_iam_principal_policy_simulation} Data Source.
// Experimental.
func NewDataAwsPrincipalPolicySimulation(scope constructs.Construct, id *string, config *DataAwsPrincipalPolicySimulationConfig) DataAwsPrincipalPolicySimulation {
	_init_.Initialize()

	if err := validateNewDataAwsPrincipalPolicySimulationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsPrincipalPolicySimulation{}

	_jsii_.Create(
		"@cdktn/aws-iam.DataAwsPrincipalPolicySimulation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_principal_policy_simulation aws_iam_principal_policy_simulation} Data Source.
// Experimental.
func NewDataAwsPrincipalPolicySimulation_Override(d DataAwsPrincipalPolicySimulation, scope constructs.Construct, id *string, config *DataAwsPrincipalPolicySimulationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iam.DataAwsPrincipalPolicySimulation",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetActionNames(val *[]*string) {
	if err := j.validateSetActionNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionNames",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetAdditionalPoliciesJson(val *[]*string) {
	if err := j.validateSetAdditionalPoliciesJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalPoliciesJson",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetCallerArn(val *string) {
	if err := j.validateSetCallerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callerArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetPermissionsBoundaryPoliciesJson(val *[]*string) {
	if err := j.validateSetPermissionsBoundaryPoliciesJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionsBoundaryPoliciesJson",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetPolicySourceArn(val *string) {
	if err := j.validateSetPolicySourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policySourceArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetResourceArns(val *[]*string) {
	if err := j.validateSetResourceArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArns",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetResourceHandlingOption(val *string) {
	if err := j.validateSetResourceHandlingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceHandlingOption",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetResourceOwnerAccountId(val *string) {
	if err := j.validateSetResourceOwnerAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceOwnerAccountId",
		val,
	)
}

func (j *jsiiProxy_DataAwsPrincipalPolicySimulation)SetResourcePolicyJson(val *string) {
	if err := j.validateSetResourcePolicyJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicyJson",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsPrincipalPolicySimulation resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataAwsPrincipalPolicySimulation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsPrincipalPolicySimulation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.DataAwsPrincipalPolicySimulation",
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
func DataAwsPrincipalPolicySimulation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsPrincipalPolicySimulation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.DataAwsPrincipalPolicySimulation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsPrincipalPolicySimulation_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsPrincipalPolicySimulation_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.DataAwsPrincipalPolicySimulation",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsPrincipalPolicySimulation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsPrincipalPolicySimulation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iam.DataAwsPrincipalPolicySimulation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsPrincipalPolicySimulation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-iam.DataAwsPrincipalPolicySimulation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) PutContext(value interface{}) {
	if err := d.validatePutContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContext",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetAdditionalPoliciesJson() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalPoliciesJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetCallerArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCallerArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetContext() {
	_jsii_.InvokeVoid(
		d,
		"resetContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetPermissionsBoundaryPoliciesJson() {
	_jsii_.InvokeVoid(
		d,
		"resetPermissionsBoundaryPoliciesJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetResourceArns() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceArns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetResourceHandlingOption() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceHandlingOption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetResourceOwnerAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceOwnerAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ResetResourcePolicyJson() {
	_jsii_.InvokeVoid(
		d,
		"resetResourcePolicyJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsPrincipalPolicySimulation) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

