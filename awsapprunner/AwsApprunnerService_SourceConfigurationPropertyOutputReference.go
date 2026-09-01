package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApprunnerService_SourceConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationConfiguration() AwsApprunnerService_AuthenticationConfigurationPropertyOutputReference
	// Experimental.
	AuthenticationConfigurationInput() *AwsApprunnerService_AuthenticationConfigurationProperty
	// Experimental.
	AutoDeploymentsEnabled() interface{}
	// Experimental.
	SetAutoDeploymentsEnabled(val interface{})
	// Experimental.
	AutoDeploymentsEnabledInput() interface{}
	// Experimental.
	CodeRepository() AwsApprunnerService_CodeRepositoryPropertyOutputReference
	// Experimental.
	CodeRepositoryInput() *AwsApprunnerService_CodeRepositoryProperty
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
	ImageRepository() AwsApprunnerService_ImageRepositoryPropertyOutputReference
	// Experimental.
	ImageRepositoryInput() *AwsApprunnerService_ImageRepositoryProperty
	// Experimental.
	InternalValue() *AwsApprunnerService_SourceConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsApprunnerService_SourceConfigurationProperty)
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
	PutAuthenticationConfiguration(value *AwsApprunnerService_AuthenticationConfigurationProperty)
	// Experimental.
	PutCodeRepository(value *AwsApprunnerService_CodeRepositoryProperty)
	// Experimental.
	PutImageRepository(value *AwsApprunnerService_ImageRepositoryProperty)
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

// The jsii proxy struct for AwsApprunnerService_SourceConfigurationPropertyOutputReference
type jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) AuthenticationConfiguration() AwsApprunnerService_AuthenticationConfigurationPropertyOutputReference {
	var returns AwsApprunnerService_AuthenticationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) AuthenticationConfigurationInput() *AwsApprunnerService_AuthenticationConfigurationProperty {
	var returns *AwsApprunnerService_AuthenticationConfigurationProperty
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) AutoDeploymentsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) AutoDeploymentsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) CodeRepository() AwsApprunnerService_CodeRepositoryPropertyOutputReference {
	var returns AwsApprunnerService_CodeRepositoryPropertyOutputReference
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) CodeRepositoryInput() *AwsApprunnerService_CodeRepositoryProperty {
	var returns *AwsApprunnerService_CodeRepositoryProperty
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ImageRepository() AwsApprunnerService_ImageRepositoryPropertyOutputReference {
	var returns AwsApprunnerService_ImageRepositoryPropertyOutputReference
	_jsii_.Get(
		j,
		"imageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ImageRepositoryInput() *AwsApprunnerService_ImageRepositoryProperty {
	var returns *AwsApprunnerService_ImageRepositoryProperty
	_jsii_.Get(
		j,
		"imageRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) InternalValue() *AwsApprunnerService_SourceConfigurationProperty {
	var returns *AwsApprunnerService_SourceConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApprunnerService_SourceConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApprunnerService_SourceConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApprunnerService_SourceConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsApprunnerService.SourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApprunnerService_SourceConfigurationPropertyOutputReference_Override(a AwsApprunnerService_SourceConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsApprunnerService.SourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference)SetAutoDeploymentsEnabled(val interface{}) {
	if err := j.validateSetAutoDeploymentsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeploymentsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference)SetInternalValue(val *AwsApprunnerService_SourceConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) PutAuthenticationConfiguration(value *AwsApprunnerService_AuthenticationConfigurationProperty) {
	if err := a.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) PutCodeRepository(value *AwsApprunnerService_CodeRepositoryProperty) {
	if err := a.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) PutImageRepository(value *AwsApprunnerService_ImageRepositoryProperty) {
	if err := a.validatePutImageRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ResetAutoDeploymentsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoDeploymentsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ResetImageRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetImageRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApprunnerService_SourceConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

