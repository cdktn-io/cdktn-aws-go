package quicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/quicksight/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/quicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription aws_quicksight_account_subscription}.
// Experimental.
type AwsAccountSubscription interface {
	cdktn.TerraformResource
	// Experimental.
	AccountName() *string
	// Experimental.
	SetAccountName(val *string)
	// Experimental.
	AccountNameInput() *string
	// Experimental.
	AccountSubscriptionStatus() *string
	// Experimental.
	ActiveDirectoryName() *string
	// Experimental.
	SetActiveDirectoryName(val *string)
	// Experimental.
	ActiveDirectoryNameInput() *string
	// Experimental.
	AdminGroup() *[]*string
	// Experimental.
	SetAdminGroup(val *[]*string)
	// Experimental.
	AdminGroupInput() *[]*string
	// Experimental.
	AdminProGroup() *[]*string
	// Experimental.
	SetAdminProGroup(val *[]*string)
	// Experimental.
	AdminProGroupInput() *[]*string
	// Experimental.
	AuthenticationMethod() *string
	// Experimental.
	SetAuthenticationMethod(val *string)
	// Experimental.
	AuthenticationMethodInput() *string
	// Experimental.
	AuthorGroup() *[]*string
	// Experimental.
	SetAuthorGroup(val *[]*string)
	// Experimental.
	AuthorGroupInput() *[]*string
	// Experimental.
	AuthorProGroup() *[]*string
	// Experimental.
	SetAuthorProGroup(val *[]*string)
	// Experimental.
	AuthorProGroupInput() *[]*string
	// Experimental.
	AwsAccountId() *string
	// Experimental.
	SetAwsAccountId(val *string)
	// Experimental.
	AwsAccountIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ContactNumber() *string
	// Experimental.
	SetContactNumber(val *string)
	// Experimental.
	ContactNumberInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DirectoryId() *string
	// Experimental.
	SetDirectoryId(val *string)
	// Experimental.
	DirectoryIdInput() *string
	// Experimental.
	Edition() *string
	// Experimental.
	SetEdition(val *string)
	// Experimental.
	EditionInput() *string
	// Experimental.
	EmailAddress() *string
	// Experimental.
	SetEmailAddress(val *string)
	// Experimental.
	EmailAddressInput() *string
	// Experimental.
	FirstName() *string
	// Experimental.
	SetFirstName(val *string)
	// Experimental.
	FirstNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	IamIdentityCenterInstanceArn() *string
	// Experimental.
	SetIamIdentityCenterInstanceArn(val *string)
	// Experimental.
	IamIdentityCenterInstanceArnInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	LastName() *string
	// Experimental.
	SetLastName(val *string)
	// Experimental.
	LastNameInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NotificationEmail() *string
	// Experimental.
	SetNotificationEmail(val *string)
	// Experimental.
	NotificationEmailInput() *string
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
	ReaderGroup() *[]*string
	// Experimental.
	SetReaderGroup(val *[]*string)
	// Experimental.
	ReaderGroupInput() *[]*string
	// Experimental.
	ReaderProGroup() *[]*string
	// Experimental.
	SetReaderProGroup(val *[]*string)
	// Experimental.
	ReaderProGroupInput() *[]*string
	// Experimental.
	Realm() *string
	// Experimental.
	SetRealm(val *string)
	// Experimental.
	RealmInput() *string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsAccountSubscription_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *AwsAccountSubscription_TimeoutsProperty)
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
	ResetActiveDirectoryName()
	// Experimental.
	ResetAdminGroup()
	// Experimental.
	ResetAdminProGroup()
	// Experimental.
	ResetAuthorGroup()
	// Experimental.
	ResetAuthorProGroup()
	// Experimental.
	ResetAwsAccountId()
	// Experimental.
	ResetContactNumber()
	// Experimental.
	ResetDirectoryId()
	// Experimental.
	ResetEmailAddress()
	// Experimental.
	ResetFirstName()
	// Experimental.
	ResetIamIdentityCenterInstanceArn()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLastName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetReaderGroup()
	// Experimental.
	ResetReaderProGroup()
	// Experimental.
	ResetRealm()
	// Experimental.
	ResetRegion()
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

// The jsii proxy struct for AwsAccountSubscription
type jsiiProxy_AwsAccountSubscription struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAccountSubscription) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AccountSubscriptionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountSubscriptionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ActiveDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ActiveDirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AdminGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AdminGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AdminProGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminProGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AdminProGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminProGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AuthorGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AuthorGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AuthorProGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorProGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AuthorProGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorProGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ContactNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ContactNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) IamIdentityCenterInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamIdentityCenterInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) IamIdentityCenterInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamIdentityCenterInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) NotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) NotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ReaderGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readerGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ReaderGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readerGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ReaderProGroup() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readerProGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) ReaderProGroupInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readerProGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) Timeouts() AwsAccountSubscription_TimeoutsPropertyOutputReference {
	var returns AwsAccountSubscription_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccountSubscription) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription aws_quicksight_account_subscription} Resource.
// Experimental.
func NewAwsAccountSubscription(scope constructs.Construct, id *string, config *AwsAccountSubscriptionConfig) AwsAccountSubscription {
	_init_.Initialize()

	if err := validateNewAwsAccountSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAccountSubscription{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsAccountSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_account_subscription aws_quicksight_account_subscription} Resource.
// Experimental.
func NewAwsAccountSubscription_Override(a AwsAccountSubscription, scope constructs.Construct, id *string, config *AwsAccountSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsAccountSubscription",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetActiveDirectoryName(val *string) {
	if err := j.validateSetActiveDirectoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryName",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetAdminGroup(val *[]*string) {
	if err := j.validateSetAdminGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminGroup",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetAdminProGroup(val *[]*string) {
	if err := j.validateSetAdminProGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminProGroup",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetAuthorGroup(val *[]*string) {
	if err := j.validateSetAuthorGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorGroup",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetAuthorProGroup(val *[]*string) {
	if err := j.validateSetAuthorProGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorProGroup",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetContactNumber(val *string) {
	if err := j.validateSetContactNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactNumber",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetIamIdentityCenterInstanceArn(val *string) {
	if err := j.validateSetIamIdentityCenterInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamIdentityCenterInstanceArn",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetNotificationEmail(val *string) {
	if err := j.validateSetNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmail",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetReaderGroup(val *[]*string) {
	if err := j.validateSetReaderGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readerGroup",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetReaderProGroup(val *[]*string) {
	if err := j.validateSetReaderProGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readerProGroup",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_AwsAccountSubscription)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a AwsAccountSubscription resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsAccountSubscription_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAccountSubscription_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-quicksight.AwsAccountSubscription",
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
func AwsAccountSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAccountSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-quicksight.AwsAccountSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAccountSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAccountSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-quicksight.AwsAccountSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAccountSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAccountSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-quicksight.AwsAccountSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAccountSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-quicksight.AwsAccountSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAccountSubscription) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAccountSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAccountSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAccountSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAccountSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAccountSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAccountSubscription) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAccountSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAccountSubscription) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAccountSubscription) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsAccountSubscription) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) PutTimeouts(value *AwsAccountSubscription_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetActiveDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetAdminGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetAdminProGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminProGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetAuthorGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetAuthorProGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorProGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetContactNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetContactNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetFirstName() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetIamIdentityCenterInstanceArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamIdentityCenterInstanceArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetLastName() {
	_jsii_.InvokeVoid(
		a,
		"resetLastName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetReaderGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetReaderGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetReaderProGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetReaderProGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetRealm() {
	_jsii_.InvokeVoid(
		a,
		"resetRealm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccountSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccountSubscription) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

