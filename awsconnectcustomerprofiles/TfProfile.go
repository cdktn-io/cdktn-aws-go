package awsconnectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile aws_customerprofiles_profile}.
// Experimental.
type TfProfile interface {
	cdktn.TerraformResource
	// Experimental.
	AccountNumber() *string
	// Experimental.
	SetAccountNumber(val *string)
	// Experimental.
	AccountNumberInput() *string
	// Experimental.
	AdditionalInformation() *string
	// Experimental.
	SetAdditionalInformation(val *string)
	// Experimental.
	AdditionalInformationInput() *string
	// Experimental.
	Address() TfProfile_AddressPropertyOutputReference
	// Experimental.
	AddressInput() *TfProfile_AddressProperty
	// Experimental.
	Attributes() *map[string]*string
	// Experimental.
	SetAttributes(val *map[string]*string)
	// Experimental.
	AttributesInput() *map[string]*string
	// Experimental.
	BillingAddress() TfProfile_BillingAddressPropertyOutputReference
	// Experimental.
	BillingAddressInput() *TfProfile_BillingAddressProperty
	// Experimental.
	BirthDate() *string
	// Experimental.
	SetBirthDate(val *string)
	// Experimental.
	BirthDateInput() *string
	// Experimental.
	BusinessEmailAddress() *string
	// Experimental.
	SetBusinessEmailAddress(val *string)
	// Experimental.
	BusinessEmailAddressInput() *string
	// Experimental.
	BusinessName() *string
	// Experimental.
	SetBusinessName(val *string)
	// Experimental.
	BusinessNameInput() *string
	// Experimental.
	BusinessPhoneNumber() *string
	// Experimental.
	SetBusinessPhoneNumber(val *string)
	// Experimental.
	BusinessPhoneNumberInput() *string
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
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameInput() *string
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
	GenderString() *string
	// Experimental.
	SetGenderString(val *string)
	// Experimental.
	GenderStringInput() *string
	// Experimental.
	HomePhoneNumber() *string
	// Experimental.
	SetHomePhoneNumber(val *string)
	// Experimental.
	HomePhoneNumberInput() *string
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
	// Experimental.
	MailingAddress() TfProfile_MailingAddressPropertyOutputReference
	// Experimental.
	MailingAddressInput() *TfProfile_MailingAddressProperty
	// Experimental.
	MiddleName() *string
	// Experimental.
	SetMiddleName(val *string)
	// Experimental.
	MiddleNameInput() *string
	// Experimental.
	MobilePhoneNumber() *string
	// Experimental.
	SetMobilePhoneNumber(val *string)
	// Experimental.
	MobilePhoneNumberInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PartyTypeString() *string
	// Experimental.
	SetPartyTypeString(val *string)
	// Experimental.
	PartyTypeStringInput() *string
	// Experimental.
	PersonalEmailAddress() *string
	// Experimental.
	SetPersonalEmailAddress(val *string)
	// Experimental.
	PersonalEmailAddressInput() *string
	// Experimental.
	PhoneNumber() *string
	// Experimental.
	SetPhoneNumber(val *string)
	// Experimental.
	PhoneNumberInput() *string
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
	ShippingAddress() TfProfile_ShippingAddressPropertyOutputReference
	// Experimental.
	ShippingAddressInput() *TfProfile_ShippingAddressProperty
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
	PutAddress(value *TfProfile_AddressProperty)
	// Experimental.
	PutBillingAddress(value *TfProfile_BillingAddressProperty)
	// Experimental.
	PutMailingAddress(value *TfProfile_MailingAddressProperty)
	// Experimental.
	PutShippingAddress(value *TfProfile_ShippingAddressProperty)
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
	ResetAccountNumber()
	// Experimental.
	ResetAdditionalInformation()
	// Experimental.
	ResetAddress()
	// Experimental.
	ResetAttributes()
	// Experimental.
	ResetBillingAddress()
	// Experimental.
	ResetBirthDate()
	// Experimental.
	ResetBusinessEmailAddress()
	// Experimental.
	ResetBusinessName()
	// Experimental.
	ResetBusinessPhoneNumber()
	// Experimental.
	ResetEmailAddress()
	// Experimental.
	ResetFirstName()
	// Experimental.
	ResetGenderString()
	// Experimental.
	ResetHomePhoneNumber()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLastName()
	// Experimental.
	ResetMailingAddress()
	// Experimental.
	ResetMiddleName()
	// Experimental.
	ResetMobilePhoneNumber()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPartyTypeString()
	// Experimental.
	ResetPersonalEmailAddress()
	// Experimental.
	ResetPhoneNumber()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetShippingAddress()
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

// The jsii proxy struct for TfProfile
type jsiiProxy_TfProfile struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfProfile) AccountNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) AccountNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) AdditionalInformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) AdditionalInformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Address() TfProfile_AddressPropertyOutputReference {
	var returns TfProfile_AddressPropertyOutputReference
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) AddressInput() *TfProfile_AddressProperty {
	var returns *TfProfile_AddressProperty
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Attributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) AttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BillingAddress() TfProfile_BillingAddressPropertyOutputReference {
	var returns TfProfile_BillingAddressPropertyOutputReference
	_jsii_.Get(
		j,
		"billingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BillingAddressInput() *TfProfile_BillingAddressProperty {
	var returns *TfProfile_BillingAddressProperty
	_jsii_.Get(
		j,
		"billingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BirthDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BirthDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BusinessEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BusinessEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BusinessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BusinessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BusinessPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) BusinessPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) GenderString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genderString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) GenderStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genderStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) HomePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) HomePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) MailingAddress() TfProfile_MailingAddressPropertyOutputReference {
	var returns TfProfile_MailingAddressPropertyOutputReference
	_jsii_.Get(
		j,
		"mailingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) MailingAddressInput() *TfProfile_MailingAddressProperty {
	var returns *TfProfile_MailingAddressProperty
	_jsii_.Get(
		j,
		"mailingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) MiddleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) MobilePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) MobilePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) PartyTypeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyTypeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) PartyTypeStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyTypeStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) PersonalEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) PersonalEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) ShippingAddress() TfProfile_ShippingAddressPropertyOutputReference {
	var returns TfProfile_ShippingAddressPropertyOutputReference
	_jsii_.Get(
		j,
		"shippingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) ShippingAddressInput() *TfProfile_ShippingAddressProperty {
	var returns *TfProfile_ShippingAddressProperty
	_jsii_.Get(
		j,
		"shippingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile aws_customerprofiles_profile} Resource.
// Experimental.
func NewTfProfile(scope constructs.Construct, id *string, config *TfProfileConfig) TfProfile {
	_init_.Initialize()

	if err := validateNewTfProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProfile{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile aws_customerprofiles_profile} Resource.
// Experimental.
func NewTfProfile_Override(t TfProfile, scope constructs.Construct, id *string, config *TfProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfProfile",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfProfile)SetAccountNumber(val *string) {
	if err := j.validateSetAccountNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountNumber",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetAdditionalInformation(val *string) {
	if err := j.validateSetAdditionalInformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInformation",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetAttributes(val *map[string]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetBirthDate(val *string) {
	if err := j.validateSetBirthDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"birthDate",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetBusinessEmailAddress(val *string) {
	if err := j.validateSetBusinessEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessEmailAddress",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetBusinessName(val *string) {
	if err := j.validateSetBusinessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessName",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetBusinessPhoneNumber(val *string) {
	if err := j.validateSetBusinessPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetGenderString(val *string) {
	if err := j.validateSetGenderStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"genderString",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetHomePhoneNumber(val *string) {
	if err := j.validateSetHomePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetMiddleName(val *string) {
	if err := j.validateSetMiddleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middleName",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetMobilePhoneNumber(val *string) {
	if err := j.validateSetMobilePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobilePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetPartyTypeString(val *string) {
	if err := j.validateSetPartyTypeStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partyTypeString",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetPersonalEmailAddress(val *string) {
	if err := j.validateSetPersonalEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"personalEmailAddress",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfProfile)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a TfProfile resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.TfProfile",
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
func TfProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.TfProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.TfProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.TfProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-connect-customer-profiles.TfProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfProfile) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfProfile) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProfile) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfProfile) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProfile) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfProfile) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfProfile) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfProfile) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfProfile) PutAddress(value *TfProfile_AddressProperty) {
	if err := t.validatePutAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAddress",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProfile) PutBillingAddress(value *TfProfile_BillingAddressProperty) {
	if err := t.validatePutBillingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBillingAddress",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProfile) PutMailingAddress(value *TfProfile_MailingAddressProperty) {
	if err := t.validatePutMailingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMailingAddress",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProfile) PutShippingAddress(value *TfProfile_ShippingAddressProperty) {
	if err := t.validatePutShippingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putShippingAddress",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProfile) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfProfile) ResetAccountNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetAccountNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetAdditionalInformation() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalInformation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetAttributes() {
	_jsii_.InvokeVoid(
		t,
		"resetAttributes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetBillingAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetBillingAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetBirthDate() {
	_jsii_.InvokeVoid(
		t,
		"resetBirthDate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetBusinessEmailAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetBusinessEmailAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetBusinessName() {
	_jsii_.InvokeVoid(
		t,
		"resetBusinessName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetBusinessPhoneNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetBusinessPhoneNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetFirstName() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetGenderString() {
	_jsii_.InvokeVoid(
		t,
		"resetGenderString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetHomePhoneNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetHomePhoneNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetLastName() {
	_jsii_.InvokeVoid(
		t,
		"resetLastName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetMailingAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetMailingAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetMiddleName() {
	_jsii_.InvokeVoid(
		t,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetMobilePhoneNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetMobilePhoneNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetPartyTypeString() {
	_jsii_.InvokeVoid(
		t,
		"resetPartyTypeString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetPersonalEmailAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetPersonalEmailAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) ResetShippingAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetShippingAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

