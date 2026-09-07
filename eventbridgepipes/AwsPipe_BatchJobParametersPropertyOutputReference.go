package eventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipe_BatchJobParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArrayProperties() AwsPipe_ArrayPropertiesPropertyOutputReference
	// Experimental.
	ArrayPropertiesInput() *AwsPipe_ArrayPropertiesProperty
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
	ContainerOverrides() AwsPipe_ContainerOverridesPropertyOutputReference
	// Experimental.
	ContainerOverridesInput() *AwsPipe_ContainerOverridesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DependsOn() AwsPipe_DependsOnPropertyList
	// Experimental.
	DependsOnInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsPipe_BatchJobParametersProperty
	// Experimental.
	SetInternalValue(val *AwsPipe_BatchJobParametersProperty)
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
	RetryStrategy() AwsPipe_RetryStrategyPropertyOutputReference
	// Experimental.
	RetryStrategyInput() *AwsPipe_RetryStrategyProperty
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
	PutArrayProperties(value *AwsPipe_ArrayPropertiesProperty)
	// Experimental.
	PutContainerOverrides(value *AwsPipe_ContainerOverridesProperty)
	// Experimental.
	PutDependsOn(value interface{})
	// Experimental.
	PutRetryStrategy(value *AwsPipe_RetryStrategyProperty)
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

// The jsii proxy struct for AwsPipe_BatchJobParametersPropertyOutputReference
type jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ArrayProperties() AwsPipe_ArrayPropertiesPropertyOutputReference {
	var returns AwsPipe_ArrayPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"arrayProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ArrayPropertiesInput() *AwsPipe_ArrayPropertiesProperty {
	var returns *AwsPipe_ArrayPropertiesProperty
	_jsii_.Get(
		j,
		"arrayPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ContainerOverrides() AwsPipe_ContainerOverridesPropertyOutputReference {
	var returns AwsPipe_ContainerOverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"containerOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ContainerOverridesInput() *AwsPipe_ContainerOverridesProperty {
	var returns *AwsPipe_ContainerOverridesProperty
	_jsii_.Get(
		j,
		"containerOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) DependsOn() AwsPipe_DependsOnPropertyList {
	var returns AwsPipe_DependsOnPropertyList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) InternalValue() *AwsPipe_BatchJobParametersProperty {
	var returns *AwsPipe_BatchJobParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) JobDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) JobDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) RetryStrategy() AwsPipe_RetryStrategyPropertyOutputReference {
	var returns AwsPipe_RetryStrategyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) RetryStrategyInput() *AwsPipe_RetryStrategyProperty {
	var returns *AwsPipe_RetryStrategyProperty
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipe_BatchJobParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipe_BatchJobParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipe_BatchJobParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.BatchJobParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipe_BatchJobParametersPropertyOutputReference_Override(a AwsPipe_BatchJobParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.BatchJobParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetInternalValue(val *AwsPipe_BatchJobParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetJobDefinition(val *string) {
	if err := j.validateSetJobDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobDefinition",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) PutArrayProperties(value *AwsPipe_ArrayPropertiesProperty) {
	if err := a.validatePutArrayPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArrayProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) PutContainerOverrides(value *AwsPipe_ContainerOverridesProperty) {
	if err := a.validatePutContainerOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainerOverrides",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) PutDependsOn(value interface{}) {
	if err := a.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) PutRetryStrategy(value *AwsPipe_RetryStrategyProperty) {
	if err := a.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ResetArrayProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetArrayProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ResetContainerOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		a,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipe_BatchJobParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

