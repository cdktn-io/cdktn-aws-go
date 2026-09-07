package apprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/apprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/apprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsService_CodeRepositoryPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CodeConfiguration() AwsService_CodeConfigurationPropertyOutputReference
	// Experimental.
	CodeConfigurationInput() *AwsService_CodeConfigurationProperty
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
	InternalValue() *AwsService_CodeRepositoryProperty
	// Experimental.
	SetInternalValue(val *AwsService_CodeRepositoryProperty)
	// Experimental.
	RepositoryUrl() *string
	// Experimental.
	SetRepositoryUrl(val *string)
	// Experimental.
	RepositoryUrlInput() *string
	// Experimental.
	SourceCodeVersion() AwsService_SourceCodeVersionPropertyOutputReference
	// Experimental.
	SourceCodeVersionInput() *AwsService_SourceCodeVersionProperty
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
	PutCodeConfiguration(value *AwsService_CodeConfigurationProperty)
	// Experimental.
	PutSourceCodeVersion(value *AwsService_SourceCodeVersionProperty)
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

// The jsii proxy struct for AwsService_CodeRepositoryPropertyOutputReference
type jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) CodeConfiguration() AwsService_CodeConfigurationPropertyOutputReference {
	var returns AwsService_CodeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"codeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) CodeConfigurationInput() *AwsService_CodeConfigurationProperty {
	var returns *AwsService_CodeConfigurationProperty
	_jsii_.Get(
		j,
		"codeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) InternalValue() *AwsService_CodeRepositoryProperty {
	var returns *AwsService_CodeRepositoryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) RepositoryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) SourceCodeVersion() AwsService_SourceCodeVersionPropertyOutputReference {
	var returns AwsService_SourceCodeVersionPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceCodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) SourceCodeVersionInput() *AwsService_SourceCodeVersionProperty {
	var returns *AwsService_SourceCodeVersionProperty
	_jsii_.Get(
		j,
		"sourceCodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) SourceDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) SourceDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsService_CodeRepositoryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsService_CodeRepositoryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsService_CodeRepositoryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsService.CodeRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsService_CodeRepositoryPropertyOutputReference_Override(a AwsService_CodeRepositoryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsService.CodeRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference)SetInternalValue(val *AwsService_CodeRepositoryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference)SetRepositoryUrl(val *string) {
	if err := j.validateSetRepositoryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryUrl",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference)SetSourceDirectory(val *string) {
	if err := j.validateSetSourceDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDirectory",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) PutCodeConfiguration(value *AwsService_CodeConfigurationProperty) {
	if err := a.validatePutCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) PutSourceCodeVersion(value *AwsService_SourceCodeVersionProperty) {
	if err := a.validatePutSourceCodeVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceCodeVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) ResetCodeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) ResetSourceDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsService_CodeRepositoryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

