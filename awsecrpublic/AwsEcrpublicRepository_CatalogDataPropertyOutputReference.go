package awsecrpublic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecrpublic/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecrpublic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcrpublicRepository_CatalogDataPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AboutText() *string
	// Experimental.
	SetAboutText(val *string)
	// Experimental.
	AboutTextInput() *string
	// Experimental.
	Architectures() *[]*string
	// Experimental.
	SetArchitectures(val *[]*string)
	// Experimental.
	ArchitecturesInput() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEcrpublicRepository_CatalogDataProperty
	// Experimental.
	SetInternalValue(val *AwsEcrpublicRepository_CatalogDataProperty)
	// Experimental.
	LogoImageBlob() *string
	// Experimental.
	SetLogoImageBlob(val *string)
	// Experimental.
	LogoImageBlobInput() *string
	// Experimental.
	OperatingSystems() *[]*string
	// Experimental.
	SetOperatingSystems(val *[]*string)
	// Experimental.
	OperatingSystemsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UsageText() *string
	// Experimental.
	SetUsageText(val *string)
	// Experimental.
	UsageTextInput() *string
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
	ResetAboutText()
	// Experimental.
	ResetArchitectures()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetLogoImageBlob()
	// Experimental.
	ResetOperatingSystems()
	// Experimental.
	ResetUsageText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEcrpublicRepository_CatalogDataPropertyOutputReference
type jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) AboutText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboutText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) AboutTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboutTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) InternalValue() *AwsEcrpublicRepository_CatalogDataProperty {
	var returns *AwsEcrpublicRepository_CatalogDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) LogoImageBlob() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoImageBlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) LogoImageBlobInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoImageBlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) OperatingSystems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatingSystems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) OperatingSystemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatingSystemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) UsageText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) UsageTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageTextInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcrpublicRepository_CatalogDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEcrpublicRepository_CatalogDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcrpublicRepository_CatalogDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecr-public.AwsEcrpublicRepository.CatalogDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcrpublicRepository_CatalogDataPropertyOutputReference_Override(a AwsEcrpublicRepository_CatalogDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecr-public.AwsEcrpublicRepository.CatalogDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetAboutText(val *string) {
	if err := j.validateSetAboutTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aboutText",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetInternalValue(val *AwsEcrpublicRepository_CatalogDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetLogoImageBlob(val *string) {
	if err := j.validateSetLogoImageBlobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoImageBlob",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetOperatingSystems(val *[]*string) {
	if err := j.validateSetOperatingSystemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystems",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference)SetUsageText(val *string) {
	if err := j.validateSetUsageTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageText",
		val,
	)
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ResetAboutText() {
	_jsii_.InvokeVoid(
		a,
		"resetAboutText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ResetArchitectures() {
	_jsii_.InvokeVoid(
		a,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ResetLogoImageBlob() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoImageBlob",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ResetOperatingSystems() {
	_jsii_.InvokeVoid(
		a,
		"resetOperatingSystems",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ResetUsageText() {
	_jsii_.InvokeVoid(
		a,
		"resetUsageText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcrpublicRepository_CatalogDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

