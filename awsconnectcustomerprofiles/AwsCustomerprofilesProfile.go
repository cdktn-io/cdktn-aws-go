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
type AwsCustomerprofilesProfile interface {
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
	Address() AwsCustomerprofilesProfile_AddressPropertyOutputReference
	// Experimental.
	AddressInput() *AwsCustomerprofilesProfile_AddressProperty
	// Experimental.
	Attributes() *map[string]*string
	// Experimental.
	SetAttributes(val *map[string]*string)
	// Experimental.
	AttributesInput() *map[string]*string
	// Experimental.
	BillingAddress() AwsCustomerprofilesProfile_BillingAddressPropertyOutputReference
	// Experimental.
	BillingAddressInput() *AwsCustomerprofilesProfile_BillingAddressProperty
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
	MailingAddress() AwsCustomerprofilesProfile_MailingAddressPropertyOutputReference
	// Experimental.
	MailingAddressInput() *AwsCustomerprofilesProfile_MailingAddressProperty
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
	ShippingAddress() AwsCustomerprofilesProfile_ShippingAddressPropertyOutputReference
	// Experimental.
	ShippingAddressInput() *AwsCustomerprofilesProfile_ShippingAddressProperty
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
	PutAddress(value *AwsCustomerprofilesProfile_AddressProperty)
	// Experimental.
	PutBillingAddress(value *AwsCustomerprofilesProfile_BillingAddressProperty)
	// Experimental.
	PutMailingAddress(value *AwsCustomerprofilesProfile_MailingAddressProperty)
	// Experimental.
	PutShippingAddress(value *AwsCustomerprofilesProfile_ShippingAddressProperty)
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

// The jsii proxy struct for AwsCustomerprofilesProfile
type jsiiProxy_AwsCustomerprofilesProfile struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) AccountNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) AccountNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) AdditionalInformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) AdditionalInformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Address() AwsCustomerprofilesProfile_AddressPropertyOutputReference {
	var returns AwsCustomerprofilesProfile_AddressPropertyOutputReference
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) AddressInput() *AwsCustomerprofilesProfile_AddressProperty {
	var returns *AwsCustomerprofilesProfile_AddressProperty
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Attributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) AttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BillingAddress() AwsCustomerprofilesProfile_BillingAddressPropertyOutputReference {
	var returns AwsCustomerprofilesProfile_BillingAddressPropertyOutputReference
	_jsii_.Get(
		j,
		"billingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BillingAddressInput() *AwsCustomerprofilesProfile_BillingAddressProperty {
	var returns *AwsCustomerprofilesProfile_BillingAddressProperty
	_jsii_.Get(
		j,
		"billingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BirthDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BirthDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BusinessEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BusinessEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BusinessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BusinessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BusinessPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) BusinessPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) GenderString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genderString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) GenderStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genderStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) HomePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) HomePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) MailingAddress() AwsCustomerprofilesProfile_MailingAddressPropertyOutputReference {
	var returns AwsCustomerprofilesProfile_MailingAddressPropertyOutputReference
	_jsii_.Get(
		j,
		"mailingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) MailingAddressInput() *AwsCustomerprofilesProfile_MailingAddressProperty {
	var returns *AwsCustomerprofilesProfile_MailingAddressProperty
	_jsii_.Get(
		j,
		"mailingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) MiddleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) MobilePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) MobilePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) PartyTypeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyTypeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) PartyTypeStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyTypeStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) PersonalEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) PersonalEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) ShippingAddress() AwsCustomerprofilesProfile_ShippingAddressPropertyOutputReference {
	var returns AwsCustomerprofilesProfile_ShippingAddressPropertyOutputReference
	_jsii_.Get(
		j,
		"shippingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) ShippingAddressInput() *AwsCustomerprofilesProfile_ShippingAddressProperty {
	var returns *AwsCustomerprofilesProfile_ShippingAddressProperty
	_jsii_.Get(
		j,
		"shippingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesProfile) TerraformResourceType() *string {
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
func NewAwsCustomerprofilesProfile(scope constructs.Construct, id *string, config *AwsCustomerprofilesProfileConfig) AwsCustomerprofilesProfile {
	_init_.Initialize()

	if err := validateNewAwsCustomerprofilesProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCustomerprofilesProfile{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile aws_customerprofiles_profile} Resource.
// Experimental.
func NewAwsCustomerprofilesProfile_Override(a AwsCustomerprofilesProfile, scope constructs.Construct, id *string, config *AwsCustomerprofilesProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesProfile",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetAccountNumber(val *string) {
	if err := j.validateSetAccountNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountNumber",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetAdditionalInformation(val *string) {
	if err := j.validateSetAdditionalInformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInformation",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetAttributes(val *map[string]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetBirthDate(val *string) {
	if err := j.validateSetBirthDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"birthDate",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetBusinessEmailAddress(val *string) {
	if err := j.validateSetBusinessEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessEmailAddress",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetBusinessName(val *string) {
	if err := j.validateSetBusinessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessName",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetBusinessPhoneNumber(val *string) {
	if err := j.validateSetBusinessPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetGenderString(val *string) {
	if err := j.validateSetGenderStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"genderString",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetHomePhoneNumber(val *string) {
	if err := j.validateSetHomePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetMiddleName(val *string) {
	if err := j.validateSetMiddleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middleName",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetMobilePhoneNumber(val *string) {
	if err := j.validateSetMobilePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobilePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetPartyTypeString(val *string) {
	if err := j.validateSetPartyTypeStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partyTypeString",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetPersonalEmailAddress(val *string) {
	if err := j.validateSetPersonalEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"personalEmailAddress",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesProfile)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a AwsCustomerprofilesProfile resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCustomerprofilesProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCustomerprofilesProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesProfile",
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
func AwsCustomerprofilesProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomerprofilesProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCustomerprofilesProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomerprofilesProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCustomerprofilesProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomerprofilesProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCustomerprofilesProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCustomerprofilesProfile) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) PutAddress(value *AwsCustomerprofilesProfile_AddressProperty) {
	if err := a.validatePutAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAddress",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) PutBillingAddress(value *AwsCustomerprofilesProfile_BillingAddressProperty) {
	if err := a.validatePutBillingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBillingAddress",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) PutMailingAddress(value *AwsCustomerprofilesProfile_MailingAddressProperty) {
	if err := a.validatePutMailingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMailingAddress",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) PutShippingAddress(value *AwsCustomerprofilesProfile_ShippingAddressProperty) {
	if err := a.validatePutShippingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putShippingAddress",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetAccountNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetAdditionalInformation() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalInformation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetBillingAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetBillingAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetBirthDate() {
	_jsii_.InvokeVoid(
		a,
		"resetBirthDate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetBusinessEmailAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetBusinessEmailAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetBusinessName() {
	_jsii_.InvokeVoid(
		a,
		"resetBusinessName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetBusinessPhoneNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetBusinessPhoneNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetFirstName() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetGenderString() {
	_jsii_.InvokeVoid(
		a,
		"resetGenderString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetHomePhoneNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetHomePhoneNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetLastName() {
	_jsii_.InvokeVoid(
		a,
		"resetLastName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetMailingAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetMailingAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetMiddleName() {
	_jsii_.InvokeVoid(
		a,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetMobilePhoneNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetMobilePhoneNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetPartyTypeString() {
	_jsii_.InvokeVoid(
		a,
		"resetPartyTypeString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetPersonalEmailAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetPersonalEmailAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ResetShippingAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetShippingAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesProfile) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

