package awsroute53domains

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53domains/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53domains/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_RegistrantContactPropertyOutputReference interface {
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
	ExtraParam() TfDomain_RegistrantContactExtraParamPropertyList
	// Experimental.
	ExtraParamInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutExtraParam(value interface{})
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
	ResetExtraParam()
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

// The jsii proxy struct for TfDomain_RegistrantContactPropertyOutputReference
type jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) AddressLine1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) AddressLine1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) AddressLine2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) AddressLine2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ContactType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ContactTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ExtraParam() TfDomain_RegistrantContactExtraParamPropertyList {
	var returns TfDomain_RegistrantContactExtraParamPropertyList
	_jsii_.Get(
		j,
		"extraParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ExtraParamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extraParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) Fax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) FaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) OrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ZipCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCodeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_RegistrantContactPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDomain_RegistrantContactPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_RegistrantContactPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53-domains.TfDomain.RegistrantContactPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_RegistrantContactPropertyOutputReference_Override(t TfDomain_RegistrantContactPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53-domains.TfDomain.RegistrantContactPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetAddressLine1(val *string) {
	if err := j.validateSetAddressLine1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine1",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetAddressLine2(val *string) {
	if err := j.validateSetAddressLine2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine2",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetContactType(val *string) {
	if err := j.validateSetContactTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactType",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetFax(val *string) {
	if err := j.validateSetFaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fax",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetOrganizationName(val *string) {
	if err := j.validateSetOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationName",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference)SetZipCode(val *string) {
	if err := j.validateSetZipCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipCode",
		val,
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) PutExtraParam(value interface{}) {
	if err := t.validatePutExtraParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExtraParam",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetAddressLine1() {
	_jsii_.InvokeVoid(
		t,
		"resetAddressLine1",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetAddressLine2() {
	_jsii_.InvokeVoid(
		t,
		"resetAddressLine2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		t,
		"resetCity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetContactType() {
	_jsii_.InvokeVoid(
		t,
		"resetContactType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetCountryCode() {
	_jsii_.InvokeVoid(
		t,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		t,
		"resetEmail",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetExtraParam() {
	_jsii_.InvokeVoid(
		t,
		"resetExtraParam",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetFax() {
	_jsii_.InvokeVoid(
		t,
		"resetFax",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetFirstName() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetLastName() {
	_jsii_.InvokeVoid(
		t,
		"resetLastName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetOrganizationName() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		t,
		"resetState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ResetZipCode() {
	_jsii_.InvokeVoid(
		t,
		"resetZipCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RegistrantContactPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

