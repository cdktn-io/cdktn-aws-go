package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfService_CodeRepositoryPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CodeConfiguration() TfService_CodeConfigurationPropertyOutputReference
	// Experimental.
	CodeConfigurationInput() *TfService_CodeConfigurationProperty
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
	InternalValue() *TfService_CodeRepositoryProperty
	// Experimental.
	SetInternalValue(val *TfService_CodeRepositoryProperty)
	// Experimental.
	RepositoryUrl() *string
	// Experimental.
	SetRepositoryUrl(val *string)
	// Experimental.
	RepositoryUrlInput() *string
	// Experimental.
	SourceCodeVersion() TfService_SourceCodeVersionPropertyOutputReference
	// Experimental.
	SourceCodeVersionInput() *TfService_SourceCodeVersionProperty
	// Experimental.
	SourceDirectory() *string
	// Experimental.
	SetSourceDirectory(val *string)
	// Experimental.
	SourceDirectoryInput() *string
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
	PutCodeConfiguration(value *TfService_CodeConfigurationProperty)
	// Experimental.
	PutSourceCodeVersion(value *TfService_SourceCodeVersionProperty)
	// Experimental.
	ResetCodeConfiguration()
	// Experimental.
	ResetSourceDirectory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfService_CodeRepositoryPropertyOutputReference
type jsiiProxy_TfService_CodeRepositoryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) CodeConfiguration() TfService_CodeConfigurationPropertyOutputReference {
	var returns TfService_CodeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"codeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) CodeConfigurationInput() *TfService_CodeConfigurationProperty {
	var returns *TfService_CodeConfigurationProperty
	_jsii_.Get(
		j,
		"codeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) InternalValue() *TfService_CodeRepositoryProperty {
	var returns *TfService_CodeRepositoryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) RepositoryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) SourceCodeVersion() TfService_SourceCodeVersionPropertyOutputReference {
	var returns TfService_SourceCodeVersionPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceCodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) SourceCodeVersionInput() *TfService_SourceCodeVersionProperty {
	var returns *TfService_SourceCodeVersionProperty
	_jsii_.Get(
		j,
		"sourceCodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) SourceDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) SourceDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfService_CodeRepositoryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfService_CodeRepositoryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfService_CodeRepositoryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfService_CodeRepositoryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.TfService.CodeRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfService_CodeRepositoryPropertyOutputReference_Override(t TfService_CodeRepositoryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.TfService.CodeRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference)SetInternalValue(val *TfService_CodeRepositoryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference)SetRepositoryUrl(val *string) {
	if err := j.validateSetRepositoryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryUrl",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference)SetSourceDirectory(val *string) {
	if err := j.validateSetSourceDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDirectory",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) PutCodeConfiguration(value *TfService_CodeConfigurationProperty) {
	if err := t.validatePutCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) PutSourceCodeVersion(value *TfService_SourceCodeVersionProperty) {
	if err := t.validatePutSourceCodeVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceCodeVersion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) ResetCodeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) ResetSourceDirectory() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceDirectory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfService_CodeRepositoryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

