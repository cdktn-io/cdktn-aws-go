package servicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/servicecatalog/jsii"

	"github.com/cdktn-io/cdktn-aws-go/servicecatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsProduct_ProvisioningArtifactParametersPropertyOutputReference interface {
	cdktn.ComplexObject
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
	DisableTemplateValidation() interface{}
	// Experimental.
	SetDisableTemplateValidation(val interface{})
	// Experimental.
	DisableTemplateValidationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsProduct_ProvisioningArtifactParametersProperty
	// Experimental.
	SetInternalValue(val *AwsProduct_ProvisioningArtifactParametersProperty)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	TemplatePhysicalId() *string
	// Experimental.
	SetTemplatePhysicalId(val *string)
	// Experimental.
	TemplatePhysicalIdInput() *string
	// Experimental.
	TemplateUrl() *string
	// Experimental.
	SetTemplateUrl(val *string)
	// Experimental.
	TemplateUrlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	ResetDescription()
	// Experimental.
	ResetDisableTemplateValidation()
	// Experimental.
	ResetName()
	// Experimental.
	ResetTemplatePhysicalId()
	// Experimental.
	ResetTemplateUrl()
	// Experimental.
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsProduct_ProvisioningArtifactParametersPropertyOutputReference
type jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) DisableTemplateValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTemplateValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) DisableTemplateValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTemplateValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) InternalValue() *AwsProduct_ProvisioningArtifactParametersProperty {
	var returns *AwsProduct_ProvisioningArtifactParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) TemplatePhysicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templatePhysicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) TemplatePhysicalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templatePhysicalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) TemplateUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsProduct_ProvisioningArtifactParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsProduct_ProvisioningArtifactParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsProduct_ProvisioningArtifactParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-service-catalog.AwsProduct.ProvisioningArtifactParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsProduct_ProvisioningArtifactParametersPropertyOutputReference_Override(a AwsProduct_ProvisioningArtifactParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-service-catalog.AwsProduct.ProvisioningArtifactParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetDisableTemplateValidation(val interface{}) {
	if err := j.validateSetDisableTemplateValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTemplateValidation",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetInternalValue(val *AwsProduct_ProvisioningArtifactParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetTemplatePhysicalId(val *string) {
	if err := j.validateSetTemplatePhysicalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templatePhysicalId",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetTemplateUrl(val *string) {
	if err := j.validateSetTemplateUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUrl",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ResetDisableTemplateValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableTemplateValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ResetTemplatePhysicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplatePhysicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ResetTemplateUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplateUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsProduct_ProvisioningArtifactParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

