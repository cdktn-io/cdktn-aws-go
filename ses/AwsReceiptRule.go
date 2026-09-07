package ses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ses/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/ses/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule aws_ses_receipt_rule}.
// Experimental.
type AwsReceiptRule interface {
	cdktn.TerraformResource
	// Experimental.
	AddHeaderAction() AwsReceiptRule_AddHeaderActionPropertyList
	// Experimental.
	AddHeaderActionInput() interface{}
	// Experimental.
	After() *string
	// Experimental.
	SetAfter(val *string)
	// Experimental.
	AfterInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	BounceAction() AwsReceiptRule_BounceActionPropertyList
	// Experimental.
	BounceActionInput() interface{}
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
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
	LambdaAction() AwsReceiptRule_LambdaActionPropertyList
	// Experimental.
	LambdaActionInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	Recipients() *[]*string
	// Experimental.
	SetRecipients(val *[]*string)
	// Experimental.
	RecipientsInput() *[]*string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RuleSetName() *string
	// Experimental.
	SetRuleSetName(val *string)
	// Experimental.
	RuleSetNameInput() *string
	// Experimental.
	S3Action() AwsReceiptRule_S3ActionPropertyList
	// Experimental.
	S3ActionInput() interface{}
	// Experimental.
	ScanEnabled() interface{}
	// Experimental.
	SetScanEnabled(val interface{})
	// Experimental.
	ScanEnabledInput() interface{}
	// Experimental.
	SnsAction() AwsReceiptRule_SnsActionPropertyList
	// Experimental.
	SnsActionInput() interface{}
	// Experimental.
	StopAction() AwsReceiptRule_StopActionPropertyList
	// Experimental.
	StopActionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TlsPolicy() *string
	// Experimental.
	SetTlsPolicy(val *string)
	// Experimental.
	TlsPolicyInput() *string
	// Experimental.
	WorkmailAction() AwsReceiptRule_WorkmailActionPropertyList
	// Experimental.
	WorkmailActionInput() interface{}
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
	PutAddHeaderAction(value interface{})
	// Experimental.
	PutBounceAction(value interface{})
	// Experimental.
	PutLambdaAction(value interface{})
	// Experimental.
	PutS3Action(value interface{})
	// Experimental.
	PutSnsAction(value interface{})
	// Experimental.
	PutStopAction(value interface{})
	// Experimental.
	PutWorkmailAction(value interface{})
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
	ResetAddHeaderAction()
	// Experimental.
	ResetAfter()
	// Experimental.
	ResetBounceAction()
	// Experimental.
	ResetEnabled()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLambdaAction()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRecipients()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetS3Action()
	// Experimental.
	ResetScanEnabled()
	// Experimental.
	ResetSnsAction()
	// Experimental.
	ResetStopAction()
	// Experimental.
	ResetTlsPolicy()
	// Experimental.
	ResetWorkmailAction()
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

// The jsii proxy struct for AwsReceiptRule
type jsiiProxy_AwsReceiptRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsReceiptRule) AddHeaderAction() AwsReceiptRule_AddHeaderActionPropertyList {
	var returns AwsReceiptRule_AddHeaderActionPropertyList
	_jsii_.Get(
		j,
		"addHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) AddHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) After() *string {
	var returns *string
	_jsii_.Get(
		j,
		"after",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) AfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) BounceAction() AwsReceiptRule_BounceActionPropertyList {
	var returns AwsReceiptRule_BounceActionPropertyList
	_jsii_.Get(
		j,
		"bounceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) BounceActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bounceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) LambdaAction() AwsReceiptRule_LambdaActionPropertyList {
	var returns AwsReceiptRule_LambdaActionPropertyList
	_jsii_.Get(
		j,
		"lambdaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) LambdaActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Recipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) RecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) RuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) RuleSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) S3Action() AwsReceiptRule_S3ActionPropertyList {
	var returns AwsReceiptRule_S3ActionPropertyList
	_jsii_.Get(
		j,
		"s3Action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) S3ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) ScanEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scanEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) ScanEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scanEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) SnsAction() AwsReceiptRule_SnsActionPropertyList {
	var returns AwsReceiptRule_SnsActionPropertyList
	_jsii_.Get(
		j,
		"snsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) SnsActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) StopAction() AwsReceiptRule_StopActionPropertyList {
	var returns AwsReceiptRule_StopActionPropertyList
	_jsii_.Get(
		j,
		"stopAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) StopActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) TlsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) TlsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) WorkmailAction() AwsReceiptRule_WorkmailActionPropertyList {
	var returns AwsReceiptRule_WorkmailActionPropertyList
	_jsii_.Get(
		j,
		"workmailAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReceiptRule) WorkmailActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workmailActionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule aws_ses_receipt_rule} Resource.
// Experimental.
func NewAwsReceiptRule(scope constructs.Construct, id *string, config *AwsReceiptRuleConfig) AwsReceiptRule {
	_init_.Initialize()

	if err := validateNewAwsReceiptRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsReceiptRule{}

	_jsii_.Create(
		"@cdktn/aws-ses.AwsReceiptRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule aws_ses_receipt_rule} Resource.
// Experimental.
func NewAwsReceiptRule_Override(a AwsReceiptRule, scope constructs.Construct, id *string, config *AwsReceiptRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses.AwsReceiptRule",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetAfter(val *string) {
	if err := j.validateSetAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"after",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetRecipients(val *[]*string) {
	if err := j.validateSetRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipients",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetRuleSetName(val *string) {
	if err := j.validateSetRuleSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleSetName",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetScanEnabled(val interface{}) {
	if err := j.validateSetScanEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsReceiptRule)SetTlsPolicy(val *string) {
	if err := j.validateSetTlsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsPolicy",
		val,
	)
}

// Generates CDKTN code for importing a AwsReceiptRule resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsReceiptRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsReceiptRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ses.AwsReceiptRule",
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
func AwsReceiptRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsReceiptRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ses.AwsReceiptRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsReceiptRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsReceiptRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ses.AwsReceiptRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsReceiptRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsReceiptRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ses.AwsReceiptRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsReceiptRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ses.AwsReceiptRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsReceiptRule) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsReceiptRule) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsReceiptRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReceiptRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsReceiptRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsReceiptRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsReceiptRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsReceiptRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsReceiptRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsReceiptRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsReceiptRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReceiptRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsReceiptRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReceiptRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsReceiptRule) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsReceiptRule) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsReceiptRule) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsReceiptRule) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsReceiptRule) PutAddHeaderAction(value interface{}) {
	if err := a.validatePutAddHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAddHeaderAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) PutBounceAction(value interface{}) {
	if err := a.validatePutBounceActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBounceAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) PutLambdaAction(value interface{}) {
	if err := a.validatePutLambdaActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) PutS3Action(value interface{}) {
	if err := a.validatePutS3ActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Action",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) PutSnsAction(value interface{}) {
	if err := a.validatePutSnsActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnsAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) PutStopAction(value interface{}) {
	if err := a.validatePutStopActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStopAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) PutWorkmailAction(value interface{}) {
	if err := a.validatePutWorkmailActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkmailAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReceiptRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetAddHeaderAction() {
	_jsii_.InvokeVoid(
		a,
		"resetAddHeaderAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetAfter() {
	_jsii_.InvokeVoid(
		a,
		"resetAfter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetBounceAction() {
	_jsii_.InvokeVoid(
		a,
		"resetBounceAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetLambdaAction() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetRecipients() {
	_jsii_.InvokeVoid(
		a,
		"resetRecipients",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetS3Action() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Action",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetScanEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetScanEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetSnsAction() {
	_jsii_.InvokeVoid(
		a,
		"resetSnsAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetStopAction() {
	_jsii_.InvokeVoid(
		a,
		"resetStopAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetTlsPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) ResetWorkmailAction() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkmailAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReceiptRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReceiptRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReceiptRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReceiptRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReceiptRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReceiptRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReceiptRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

