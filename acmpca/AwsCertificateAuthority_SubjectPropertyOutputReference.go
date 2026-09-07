package acmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/acmpca/jsii"

	"github.com/cdktn-io/cdktn-aws-go/acmpca/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCertificateAuthority_SubjectPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CommonName() *string
	// Experimental.
	SetCommonName(val *string)
	// Experimental.
	CommonNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DistinguishedNameQualifier() *string
	// Experimental.
	SetDistinguishedNameQualifier(val *string)
	// Experimental.
	DistinguishedNameQualifierInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GenerationQualifier() *string
	// Experimental.
	SetGenerationQualifier(val *string)
	// Experimental.
	GenerationQualifierInput() *string
	// Experimental.
	GivenName() *string
	// Experimental.
	SetGivenName(val *string)
	// Experimental.
	GivenNameInput() *string
	// Experimental.
	Initials() *string
	// Experimental.
	SetInitials(val *string)
	// Experimental.
	InitialsInput() *string
	// Experimental.
	InternalValue() *AwsCertificateAuthority_SubjectProperty
	// Experimental.
	SetInternalValue(val *AwsCertificateAuthority_SubjectProperty)
	// Experimental.
	Locality() *string
	// Experimental.
	SetLocality(val *string)
	// Experimental.
	LocalityInput() *string
	// Experimental.
	Organization() *string
	// Experimental.
	SetOrganization(val *string)
	// Experimental.
	OrganizationalUnit() *string
	// Experimental.
	SetOrganizationalUnit(val *string)
	// Experimental.
	OrganizationalUnitInput() *string
	// Experimental.
	OrganizationInput() *string
	// Experimental.
	Pseudonym() *string
	// Experimental.
	SetPseudonym(val *string)
	// Experimental.
	PseudonymInput() *string
	// Experimental.
	State() *string
	// Experimental.
	SetState(val *string)
	// Experimental.
	StateInput() *string
	// Experimental.
	Surname() *string
	// Experimental.
	SetSurname(val *string)
	// Experimental.
	SurnameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Title() *string
	// Experimental.
	SetTitle(val *string)
	// Experimental.
	TitleInput() *string
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
	ResetCommonName()
	// Experimental.
	ResetCountry()
	// Experimental.
	ResetDistinguishedNameQualifier()
	// Experimental.
	ResetGenerationQualifier()
	// Experimental.
	ResetGivenName()
	// Experimental.
	ResetInitials()
	// Experimental.
	ResetLocality()
	// Experimental.
	ResetOrganization()
	// Experimental.
	ResetOrganizationalUnit()
	// Experimental.
	ResetPseudonym()
	// Experimental.
	ResetState()
	// Experimental.
	ResetSurname()
	// Experimental.
	ResetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCertificateAuthority_SubjectPropertyOutputReference
type jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) DistinguishedNameQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distinguishedNameQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) DistinguishedNameQualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distinguishedNameQualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GenerationQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generationQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GenerationQualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generationQualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GivenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GivenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Initials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) InitialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) InternalValue() *AwsCertificateAuthority_SubjectProperty {
	var returns *AwsCertificateAuthority_SubjectProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) OrganizationalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) OrganizationalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Pseudonym() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pseudonym",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) PseudonymInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pseudonymInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Surname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) SurnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCertificateAuthority_SubjectPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCertificateAuthority_SubjectPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCertificateAuthority_SubjectPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-acm-pca.AwsCertificateAuthority.SubjectPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCertificateAuthority_SubjectPropertyOutputReference_Override(a AwsCertificateAuthority_SubjectPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-acm-pca.AwsCertificateAuthority.SubjectPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetDistinguishedNameQualifier(val *string) {
	if err := j.validateSetDistinguishedNameQualifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distinguishedNameQualifier",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetGenerationQualifier(val *string) {
	if err := j.validateSetGenerationQualifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generationQualifier",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetGivenName(val *string) {
	if err := j.validateSetGivenNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"givenName",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetInitials(val *string) {
	if err := j.validateSetInitialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initials",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetInternalValue(val *AwsCertificateAuthority_SubjectProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetOrganizationalUnit(val *string) {
	if err := j.validateSetOrganizationalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnit",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetPseudonym(val *string) {
	if err := j.validateSetPseudonymParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pseudonym",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetSurname(val *string) {
	if err := j.validateSetSurnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surname",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		a,
		"resetCommonName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetCountry() {
	_jsii_.InvokeVoid(
		a,
		"resetCountry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetDistinguishedNameQualifier() {
	_jsii_.InvokeVoid(
		a,
		"resetDistinguishedNameQualifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetGenerationQualifier() {
	_jsii_.InvokeVoid(
		a,
		"resetGenerationQualifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetGivenName() {
	_jsii_.InvokeVoid(
		a,
		"resetGivenName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetInitials() {
	_jsii_.InvokeVoid(
		a,
		"resetInitials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetLocality() {
	_jsii_.InvokeVoid(
		a,
		"resetLocality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetOrganization() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetOrganizationalUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationalUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetPseudonym() {
	_jsii_.InvokeVoid(
		a,
		"resetPseudonym",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		a,
		"resetState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetSurname() {
	_jsii_.InvokeVoid(
		a,
		"resetSurname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCertificateAuthority_SubjectPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

