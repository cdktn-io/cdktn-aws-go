package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_BatchJobParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArrayProperties() TfPipe_ArrayPropertiesPropertyOutputReference
	// Experimental.
	ArrayPropertiesInput() *TfPipe_ArrayPropertiesProperty
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
	// Experimental.
	ContainerOverrides() TfPipe_ContainerOverridesPropertyOutputReference
	// Experimental.
	ContainerOverridesInput() *TfPipe_ContainerOverridesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DependsOn() TfPipe_DependsOnPropertyList
	// Experimental.
	DependsOnInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPipe_BatchJobParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_BatchJobParametersProperty)
	// Experimental.
	JobDefinition() *string
	// Experimental.
	SetJobDefinition(val *string)
	// Experimental.
	JobDefinitionInput() *string
	// Experimental.
	JobName() *string
	// Experimental.
	SetJobName(val *string)
	// Experimental.
	JobNameInput() *string
	// Experimental.
	Parameters() *map[string]*string
	// Experimental.
	SetParameters(val *map[string]*string)
	// Experimental.
	ParametersInput() *map[string]*string
	// Experimental.
	RetryStrategy() TfPipe_RetryStrategyPropertyOutputReference
	// Experimental.
	RetryStrategyInput() *TfPipe_RetryStrategyProperty
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
	PutArrayProperties(value *TfPipe_ArrayPropertiesProperty)
	// Experimental.
	PutContainerOverrides(value *TfPipe_ContainerOverridesProperty)
	// Experimental.
	PutDependsOn(value interface{})
	// Experimental.
	PutRetryStrategy(value *TfPipe_RetryStrategyProperty)
	// Experimental.
	ResetArrayProperties()
	// Experimental.
	ResetContainerOverrides()
	// Experimental.
	ResetDependsOn()
	// Experimental.
	ResetParameters()
	// Experimental.
	ResetRetryStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_BatchJobParametersPropertyOutputReference
type jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ArrayProperties() TfPipe_ArrayPropertiesPropertyOutputReference {
	var returns TfPipe_ArrayPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"arrayProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ArrayPropertiesInput() *TfPipe_ArrayPropertiesProperty {
	var returns *TfPipe_ArrayPropertiesProperty
	_jsii_.Get(
		j,
		"arrayPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ContainerOverrides() TfPipe_ContainerOverridesPropertyOutputReference {
	var returns TfPipe_ContainerOverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"containerOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ContainerOverridesInput() *TfPipe_ContainerOverridesProperty {
	var returns *TfPipe_ContainerOverridesProperty
	_jsii_.Get(
		j,
		"containerOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) DependsOn() TfPipe_DependsOnPropertyList {
	var returns TfPipe_DependsOnPropertyList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) InternalValue() *TfPipe_BatchJobParametersProperty {
	var returns *TfPipe_BatchJobParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) JobDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) JobDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) RetryStrategy() TfPipe_RetryStrategyPropertyOutputReference {
	var returns TfPipe_RetryStrategyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) RetryStrategyInput() *TfPipe_RetryStrategyProperty {
	var returns *TfPipe_RetryStrategyProperty
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_BatchJobParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_BatchJobParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_BatchJobParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.BatchJobParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_BatchJobParametersPropertyOutputReference_Override(t TfPipe_BatchJobParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.BatchJobParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetInternalValue(val *TfPipe_BatchJobParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetJobDefinition(val *string) {
	if err := j.validateSetJobDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobDefinition",
		val,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) PutArrayProperties(value *TfPipe_ArrayPropertiesProperty) {
	if err := t.validatePutArrayPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArrayProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) PutContainerOverrides(value *TfPipe_ContainerOverridesProperty) {
	if err := t.validatePutContainerOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContainerOverrides",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) PutDependsOn(value interface{}) {
	if err := t.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) PutRetryStrategy(value *TfPipe_RetryStrategyProperty) {
	if err := t.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ResetArrayProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetArrayProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ResetContainerOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		t,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_BatchJobParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

