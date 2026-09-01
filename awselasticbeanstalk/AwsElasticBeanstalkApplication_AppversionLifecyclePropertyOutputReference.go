package awselasticbeanstalk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselasticbeanstalk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselasticbeanstalk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference interface {
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
	DeleteSourceFromS3() interface{}
	// Experimental.
	SetDeleteSourceFromS3(val interface{})
	// Experimental.
	DeleteSourceFromS3Input() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsElasticBeanstalkApplication_AppversionLifecycleProperty
	// Experimental.
	SetInternalValue(val *AwsElasticBeanstalkApplication_AppversionLifecycleProperty)
	// Experimental.
	MaxAgeInDays() *float64
	// Experimental.
	SetMaxAgeInDays(val *float64)
	// Experimental.
	MaxAgeInDaysInput() *float64
	// Experimental.
	MaxCount() *float64
	// Experimental.
	SetMaxCount(val *float64)
	// Experimental.
	MaxCountInput() *float64
	// Experimental.
	ServiceRole() *string
	// Experimental.
	SetServiceRole(val *string)
	// Experimental.
	ServiceRoleInput() *string
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
	ResetDeleteSourceFromS3()
	// Experimental.
	ResetMaxAgeInDays()
	// Experimental.
	ResetMaxCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference
type jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) DeleteSourceFromS3() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteSourceFromS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) DeleteSourceFromS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteSourceFromS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) InternalValue() *AwsElasticBeanstalkApplication_AppversionLifecycleProperty {
	var returns *AwsElasticBeanstalkApplication_AppversionLifecycleProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) MaxAgeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) MaxAgeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) MaxCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) MaxCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elastic-beanstalk.AwsElasticBeanstalkApplication.AppversionLifecyclePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference_Override(a AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elastic-beanstalk.AwsElasticBeanstalkApplication.AppversionLifecyclePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetDeleteSourceFromS3(val interface{}) {
	if err := j.validateSetDeleteSourceFromS3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteSourceFromS3",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetInternalValue(val *AwsElasticBeanstalkApplication_AppversionLifecycleProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetMaxAgeInDays(val *float64) {
	if err := j.validateSetMaxAgeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAgeInDays",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetMaxCount(val *float64) {
	if err := j.validateSetMaxCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ResetDeleteSourceFromS3() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteSourceFromS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ResetMaxAgeInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAgeInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ResetMaxCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsElasticBeanstalkApplication_AppversionLifecyclePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

