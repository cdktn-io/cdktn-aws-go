package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference interface {
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
	InternalValue() *AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty)
	// Experimental.
	MaxConcurrency() *float64
	// Experimental.
	SetMaxConcurrency(val *float64)
	// Experimental.
	MaxConcurrencyInput() *float64
	// Experimental.
	MemorySizeInMb() *float64
	// Experimental.
	SetMemorySizeInMb(val *float64)
	// Experimental.
	MemorySizeInMbInput() *float64
	// Experimental.
	ProvisionedConcurrency() *float64
	// Experimental.
	SetProvisionedConcurrency(val *float64)
	// Experimental.
	ProvisionedConcurrencyInput() *float64
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
	ResetProvisionedConcurrency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) InternalValue() *AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty {
	var returns *AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) MaxConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) MaxConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) MemorySizeInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) MemorySizeInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) ProvisionedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) ProvisionedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpointConfiguration.ShadowProductionVariantsServerlessConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference_Override(a AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpointConfiguration.ShadowProductionVariantsServerlessConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetInternalValue(val *AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetMaxConcurrency(val *float64) {
	if err := j.validateSetMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetMemorySizeInMb(val *float64) {
	if err := j.validateSetMemorySizeInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeInMb",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetProvisionedConcurrency(val *float64) {
	if err := j.validateSetProvisionedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedConcurrency",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) ResetProvisionedConcurrency() {
	_jsii_.InvokeVoid(
		a,
		"resetProvisionedConcurrency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

