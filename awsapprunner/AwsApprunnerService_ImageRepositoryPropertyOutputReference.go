package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApprunnerService_ImageRepositoryPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	ImageConfiguration() AwsApprunnerService_ImageConfigurationPropertyOutputReference
	// Experimental.
	ImageConfigurationInput() *AwsApprunnerService_ImageConfigurationProperty
	// Experimental.
	ImageIdentifier() *string
	// Experimental.
	SetImageIdentifier(val *string)
	// Experimental.
	ImageIdentifierInput() *string
	// Experimental.
	ImageRepositoryType() *string
	// Experimental.
	SetImageRepositoryType(val *string)
	// Experimental.
	ImageRepositoryTypeInput() *string
	// Experimental.
	InternalValue() *AwsApprunnerService_ImageRepositoryProperty
	// Experimental.
	SetInternalValue(val *AwsApprunnerService_ImageRepositoryProperty)
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
	PutImageConfiguration(value *AwsApprunnerService_ImageConfigurationProperty)
	// Experimental.
	ResetImageConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsApprunnerService_ImageRepositoryPropertyOutputReference
type jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ImageConfiguration() AwsApprunnerService_ImageConfigurationPropertyOutputReference {
	var returns AwsApprunnerService_ImageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"imageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ImageConfigurationInput() *AwsApprunnerService_ImageConfigurationProperty {
	var returns *AwsApprunnerService_ImageConfigurationProperty
	_jsii_.Get(
		j,
		"imageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ImageIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ImageIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ImageRepositoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepositoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ImageRepositoryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepositoryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) InternalValue() *AwsApprunnerService_ImageRepositoryProperty {
	var returns *AwsApprunnerService_ImageRepositoryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApprunnerService_ImageRepositoryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApprunnerService_ImageRepositoryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApprunnerService_ImageRepositoryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsApprunnerService.ImageRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApprunnerService_ImageRepositoryPropertyOutputReference_Override(a AwsApprunnerService_ImageRepositoryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsApprunnerService.ImageRepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference)SetImageIdentifier(val *string) {
	if err := j.validateSetImageIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference)SetImageRepositoryType(val *string) {
	if err := j.validateSetImageRepositoryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageRepositoryType",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference)SetInternalValue(val *AwsApprunnerService_ImageRepositoryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) PutImageConfiguration(value *AwsApprunnerService_ImageConfigurationProperty) {
	if err := a.validatePutImageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ResetImageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetImageConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApprunnerService_ImageRepositoryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

