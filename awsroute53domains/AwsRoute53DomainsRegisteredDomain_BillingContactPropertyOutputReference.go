package awsroute53domains

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53domains/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53domains/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddressLine1() *string
	// Experimental.
	SetAddressLine1(val *string)
	// Experimental.
	AddressLine1Input() *string
	// Experimental.
	AddressLine2() *string
	// Experimental.
	SetAddressLine2(val *string)
	// Experimental.
	AddressLine2Input() *string
	// Experimental.
	City() *string
	// Experimental.
	SetCity(val *string)
	// Experimental.
	CityInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	ContactType() *string
	// Experimental.
	SetContactType(val *string)
	// Experimental.
	ContactTypeInput() *string
	// Experimental.
	CountryCode() *string
	// Experimental.
	SetCountryCode(val *string)
	// Experimental.
	CountryCodeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Email() *string
	// Experimental.
	SetEmail(val *string)
	// Experimental.
	EmailInput() *string
	// Experimental.
	ExtraParams() *map[string]*string
	// Experimental.
	SetExtraParams(val *map[string]*string)
	// Experimental.
	ExtraParamsInput() *map[string]*string
	// Experimental.
	Fax() *string
	// Experimental.
	SetFax(val *string)
	// Experimental.
	FaxInput() *string
	// Experimental.
	FirstName() *string
	// Experimental.
	SetFirstName(val *string)
	// Experimental.
	FirstNameInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsRoute53DomainsRegisteredDomain_BillingContactProperty
	// Experimental.
	SetInternalValue(val *AwsRoute53DomainsRegisteredDomain_BillingContactProperty)
	// Experimental.
	LastName() *string
	// Experimental.
	SetLastName(val *string)
	// Experimental.
	LastNameInput() *string
	// Experimental.
	OrganizationName() *string
	// Experimental.
	SetOrganizationName(val *string)
	// Experimental.
	OrganizationNameInput() *string
	// Experimental.
	PhoneNumber() *string
	// Experimental.
	SetPhoneNumber(val *string)
	// Experimental.
	PhoneNumberInput() *string
	// Experimental.
	State() *string
	// Experimental.
	SetState(val *string)
	// Experimental.
	StateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ZipCode() *string
	// Experimental.
	SetZipCode(val *string)
	// Experimental.
	ZipCodeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	ResetAddressLine1()
	// Experimental.
	ResetAddressLine2()
	// Experimental.
	ResetCity()
	// Experimental.
	ResetContactType()
	// Experimental.
	ResetCountryCode()
	// Experimental.
	ResetEmail()
	// Experimental.
	ResetExtraParams()
	// Experimental.
	ResetFax()
	// Experimental.
	ResetFirstName()
	// Experimental.
	ResetLastName()
	// Experimental.
	ResetOrganizationName()
	// Experimental.
	ResetPhoneNumber()
	// Experimental.
	ResetState()
	// Experimental.
	ResetZipCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference
type jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) AddressLine1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) AddressLine1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) AddressLine2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) AddressLine2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ContactType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ContactTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ExtraParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ExtraParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) Fax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) FaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) InternalValue() *AwsRoute53DomainsRegisteredDomain_BillingContactProperty {
	var returns *AwsRoute53DomainsRegisteredDomain_BillingContactProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) OrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ZipCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCodeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain.BillingContactPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference_Override(a AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain.BillingContactPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetAddressLine1(val *string) {
	if err := j.validateSetAddressLine1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine1",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetAddressLine2(val *string) {
	if err := j.validateSetAddressLine2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine2",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetContactType(val *string) {
	if err := j.validateSetContactTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactType",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetExtraParams(val *map[string]*string) {
	if err := j.validateSetExtraParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraParams",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetFax(val *string) {
	if err := j.validateSetFaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fax",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetInternalValue(val *AwsRoute53DomainsRegisteredDomain_BillingContactProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetOrganizationName(val *string) {
	if err := j.validateSetOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationName",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference)SetZipCode(val *string) {
	if err := j.validateSetZipCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipCode",
		val,
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetAddressLine1() {
	_jsii_.InvokeVoid(
		a,
		"resetAddressLine1",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetAddressLine2() {
	_jsii_.InvokeVoid(
		a,
		"resetAddressLine2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		a,
		"resetCity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetContactType() {
	_jsii_.InvokeVoid(
		a,
		"resetContactType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetCountryCode() {
	_jsii_.InvokeVoid(
		a,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetExtraParams() {
	_jsii_.InvokeVoid(
		a,
		"resetExtraParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetFax() {
	_jsii_.InvokeVoid(
		a,
		"resetFax",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetFirstName() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetLastName() {
	_jsii_.InvokeVoid(
		a,
		"resetLastName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetOrganizationName() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		a,
		"resetState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ResetZipCode() {
	_jsii_.InvokeVoid(
		a,
		"resetZipCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

