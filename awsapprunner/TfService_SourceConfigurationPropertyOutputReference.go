package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfService_SourceConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationConfiguration() TfService_AuthenticationConfigurationPropertyOutputReference
	// Experimental.
	AuthenticationConfigurationInput() *TfService_AuthenticationConfigurationProperty
	// Experimental.
	AutoDeploymentsEnabled() interface{}
	// Experimental.
	SetAutoDeploymentsEnabled(val interface{})
	// Experimental.
	AutoDeploymentsEnabledInput() interface{}
	// Experimental.
	CodeRepository() TfService_CodeRepositoryPropertyOutputReference
	// Experimental.
	CodeRepositoryInput() *TfService_CodeRepositoryProperty
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
	Fqn() *string
	// Experimental.
	ImageRepository() TfService_ImageRepositoryPropertyOutputReference
	// Experimental.
	ImageRepositoryInput() *TfService_ImageRepositoryProperty
	// Experimental.
	InternalValue() *TfService_SourceConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfService_SourceConfigurationProperty)
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
	PutAuthenticationConfiguration(value *TfService_AuthenticationConfigurationProperty)
	// Experimental.
	PutCodeRepository(value *TfService_CodeRepositoryProperty)
	// Experimental.
	PutImageRepository(value *TfService_ImageRepositoryProperty)
	// Experimental.
	ResetAuthenticationConfiguration()
	// Experimental.
	ResetAutoDeploymentsEnabled()
	// Experimental.
	ResetCodeRepository()
	// Experimental.
	ResetImageRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfService_SourceConfigurationPropertyOutputReference
type jsiiProxy_TfService_SourceConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) AuthenticationConfiguration() TfService_AuthenticationConfigurationPropertyOutputReference {
	var returns TfService_AuthenticationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) AuthenticationConfigurationInput() *TfService_AuthenticationConfigurationProperty {
	var returns *TfService_AuthenticationConfigurationProperty
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) AutoDeploymentsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) AutoDeploymentsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) CodeRepository() TfService_CodeRepositoryPropertyOutputReference {
	var returns TfService_CodeRepositoryPropertyOutputReference
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) CodeRepositoryInput() *TfService_CodeRepositoryProperty {
	var returns *TfService_CodeRepositoryProperty
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ImageRepository() TfService_ImageRepositoryPropertyOutputReference {
	var returns TfService_ImageRepositoryPropertyOutputReference
	_jsii_.Get(
		j,
		"imageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ImageRepositoryInput() *TfService_ImageRepositoryProperty {
	var returns *TfService_ImageRepositoryProperty
	_jsii_.Get(
		j,
		"imageRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) InternalValue() *TfService_SourceConfigurationProperty {
	var returns *TfService_SourceConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfService_SourceConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfService_SourceConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfService_SourceConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfService_SourceConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.TfService.SourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfService_SourceConfigurationPropertyOutputReference_Override(t TfService_SourceConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.TfService.SourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference)SetAutoDeploymentsEnabled(val interface{}) {
	if err := j.validateSetAutoDeploymentsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeploymentsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference)SetInternalValue(val *TfService_SourceConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) PutAuthenticationConfiguration(value *TfService_AuthenticationConfigurationProperty) {
	if err := t.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) PutCodeRepository(value *TfService_CodeRepositoryProperty) {
	if err := t.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) PutImageRepository(value *TfService_ImageRepositoryProperty) {
	if err := t.validatePutImageRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageRepository",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ResetAutoDeploymentsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoDeploymentsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ResetImageRepository() {
	_jsii_.InvokeVoid(
		t,
		"resetImageRepository",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfService_SourceConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

