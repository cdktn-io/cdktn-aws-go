package awsconnectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProfile_MailingAddressPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Address1() *string
	// Experimental.
	SetAddress1(val *string)
	// Experimental.
	Address1Input() *string
	// Experimental.
	Address2() *string
	// Experimental.
	SetAddress2(val *string)
	// Experimental.
	Address2Input() *string
	// Experimental.
	Address3() *string
	// Experimental.
	SetAddress3(val *string)
	// Experimental.
	Address3Input() *string
	// Experimental.
	Address4() *string
	// Experimental.
	SetAddress4(val *string)
	// Experimental.
	Address4Input() *string
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
	Country() *string
	// Experimental.
	SetCountry(val *string)
	// Experimental.
	CountryInput() *string
	// Experimental.
	County() *string
	// Experimental.
	SetCounty(val *string)
	// Experimental.
	CountyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfProfile_MailingAddressProperty
	// Experimental.
	SetInternalValue(val *TfProfile_MailingAddressProperty)
	// Experimental.
	PostalCode() *string
	// Experimental.
	SetPostalCode(val *string)
	// Experimental.
	PostalCodeInput() *string
	// Experimental.
	Province() *string
	// Experimental.
	SetProvince(val *string)
	// Experimental.
	ProvinceInput() *string
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
	ResetAddress1()
	// Experimental.
	ResetAddress2()
	// Experimental.
	ResetAddress3()
	// Experimental.
	ResetAddress4()
	// Experimental.
	ResetCity()
	// Experimental.
	ResetCountry()
	// Experimental.
	ResetCounty()
	// Experimental.
	ResetPostalCode()
	// Experimental.
	ResetProvince()
	// Experimental.
	ResetState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfProfile_MailingAddressPropertyOutputReference
type jsiiProxy_TfProfile_MailingAddressPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Address4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) County() *string {
	var returns *string
	_jsii_.Get(
		j,
		"county",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) CountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) InternalValue() *TfProfile_MailingAddressProperty {
	var returns *TfProfile_MailingAddressProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Province() *string {
	var returns *string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfProfile_MailingAddressPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfProfile_MailingAddressPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfProfile_MailingAddressPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProfile_MailingAddressPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfProfile.MailingAddressPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfProfile_MailingAddressPropertyOutputReference_Override(t TfProfile_MailingAddressPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfProfile.MailingAddressPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetAddress1(val *string) {
	if err := j.validateSetAddress1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address1",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetAddress2(val *string) {
	if err := j.validateSetAddress2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address2",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetAddress3(val *string) {
	if err := j.validateSetAddress3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address3",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetAddress4(val *string) {
	if err := j.validateSetAddress4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address4",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetCounty(val *string) {
	if err := j.validateSetCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"county",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetInternalValue(val *TfProfile_MailingAddressProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetProvince(val *string) {
	if err := j.validateSetProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"province",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetAddress1() {
	_jsii_.InvokeVoid(
		t,
		"resetAddress1",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetAddress2() {
	_jsii_.InvokeVoid(
		t,
		"resetAddress2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetAddress3() {
	_jsii_.InvokeVoid(
		t,
		"resetAddress3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetAddress4() {
	_jsii_.InvokeVoid(
		t,
		"resetAddress4",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		t,
		"resetCity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetCountry() {
	_jsii_.InvokeVoid(
		t,
		"resetCountry",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetCounty() {
	_jsii_.InvokeVoid(
		t,
		"resetCounty",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		t,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetProvince() {
	_jsii_.InvokeVoid(
		t,
		"resetProvince",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		t,
		"resetState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfProfile_MailingAddressPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

