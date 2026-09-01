package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApprunnerService_CodeRepositoryPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CodeConfiguration() AwsApprunnerService_CodeConfigurationPropertyOutputReference
	// Experimental.
	CodeConfigurationInput() *AwsApprunnerService_CodeConfigurationProperty
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
	InternalValue() *AwsApprunnerService_CodeRepositoryProperty
	// Experimental.
	SetInternalValue(val *AwsApprunnerService_CodeRepositoryProperty)
	// Experimental.
	RepositoryUrl() *string
	// Experimental.
	SetRepositoryUrl(val *string)
	// Experimental.
	RepositoryUrlInput() *string
	// Experimental.
	SourceCodeVersion() AwsApprunnerService_SourceCodeVersionPropertyOutputReference
	// Experimental.
	SourceCodeVersionInput() *AwsApprunnerService_SourceCodeVersionProperty
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
	PutCodeConfiguration(value *AwsApprunnerService_CodeConfigurationProperty)
	// Experimental.
	PutSourceCodeVersion(value *AwsApprunnerService_SourceCodeVersionProperty)
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

// The jsii proxy struct for AwsApprunnerService_CodeRepositoryPropertyOutputReference
type jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) CodeConfiguration() AwsApprunnerService_CodeConfigurationPropertyOutputReference {
	var returns AwsApprunnerService_CodeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"codeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) CodeConfigurationInput() *AwsApprunnerService_CodeConfigurationProperty {
	var returns *AwsApprunnerService_CodeConfigurationProperty
	_jsii_.Get(
		j,
		"codeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) InternalValue() *AwsApprunnerService_CodeRepositoryProperty {
	var returns *AwsApprunnerService_CodeRepositoryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) RepositoryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) SourceCodeVersion() AwsApprunnerService_SourceCodeVersionPropertyOutputReference {
	var returns AwsApprunnerService_SourceCodeVersionPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceCodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) SourceCodeVersionInput() *AwsApprunnerService_SourceCodeVersionProperty {
	var returns *AwsApprunnerService_SourceCodeVersionProperty
	_jsii_.Get(
		j,
		"sourceCodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) SourceDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) SourceDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApprunnerService_CodeRepositoryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApprunnerService_CodeRepositoryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApprunnerService_CodeRepositoryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsApprunnerService.CodeRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApprunnerService_CodeRepositoryPropertyOutputReference_Override(a AwsApprunnerService_CodeRepositoryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsApprunnerService.CodeRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference)SetInternalValue(val *AwsApprunnerService_CodeRepositoryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference)SetRepositoryUrl(val *string) {
	if err := j.validateSetRepositoryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryUrl",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference)SetSourceDirectory(val *string) {
	if err := j.validateSetSourceDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDirectory",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) PutCodeConfiguration(value *AwsApprunnerService_CodeConfigurationProperty) {
	if err := a.validatePutCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) PutSourceCodeVersion(value *AwsApprunnerService_SourceCodeVersionProperty) {
	if err := a.validatePutSourceCodeVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceCodeVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) ResetCodeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) ResetSourceDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApprunnerService_CodeRepositoryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

