package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchStrategy() *string
	// Experimental.
	SetBatchStrategy(val *string)
	// Experimental.
	BatchStrategyInput() *string
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
	Environment() *map[string]*string
	// Experimental.
	SetEnvironment(val *map[string]*string)
	// Experimental.
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxConcurrentTransforms() *float64
	// Experimental.
	SetMaxConcurrentTransforms(val *float64)
	// Experimental.
	MaxConcurrentTransformsInput() *float64
	// Experimental.
	MaxPayloadInMb() *float64
	// Experimental.
	SetMaxPayloadInMb(val *float64)
	// Experimental.
	MaxPayloadInMbInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransformInput() AwsSagemakerAlgorithm_TransformInputPropertyList
	// Experimental.
	TransformInputInput() interface{}
	// Experimental.
	TransformOutput() AwsSagemakerAlgorithm_TransformOutputPropertyList
	// Experimental.
	TransformOutputInput() interface{}
	// Experimental.
	TransformResources() AwsSagemakerAlgorithm_TransformResourcesPropertyList
	// Experimental.
	TransformResourcesInput() interface{}
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
	PutTransformInput(value interface{})
	// Experimental.
	PutTransformOutput(value interface{})
	// Experimental.
	PutTransformResources(value interface{})
	// Experimental.
	ResetBatchStrategy()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetMaxConcurrentTransforms()
	// Experimental.
	ResetMaxPayloadInMb()
	// Experimental.
	ResetTransformInput()
	// Experimental.
	ResetTransformOutput()
	// Experimental.
	ResetTransformResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference
type jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) BatchStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) BatchStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) MaxConcurrentTransforms() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTransforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) MaxConcurrentTransformsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTransformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) MaxPayloadInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) MaxPayloadInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TransformInput() AwsSagemakerAlgorithm_TransformInputPropertyList {
	var returns AwsSagemakerAlgorithm_TransformInputPropertyList
	_jsii_.Get(
		j,
		"transformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TransformInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TransformOutput() AwsSagemakerAlgorithm_TransformOutputPropertyList {
	var returns AwsSagemakerAlgorithm_TransformOutputPropertyList
	_jsii_.Get(
		j,
		"transformOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TransformOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TransformResources() AwsSagemakerAlgorithm_TransformResourcesPropertyList {
	var returns AwsSagemakerAlgorithm_TransformResourcesPropertyList
	_jsii_.Get(
		j,
		"transformResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) TransformResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformResourcesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.TransformJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference_Override(a AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.TransformJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetBatchStrategy(val *string) {
	if err := j.validateSetBatchStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetMaxConcurrentTransforms(val *float64) {
	if err := j.validateSetMaxConcurrentTransformsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentTransforms",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetMaxPayloadInMb(val *float64) {
	if err := j.validateSetMaxPayloadInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPayloadInMb",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) PutTransformInput(value interface{}) {
	if err := a.validatePutTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransformInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) PutTransformOutput(value interface{}) {
	if err := a.validatePutTransformOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransformOutput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) PutTransformResources(value interface{}) {
	if err := a.validatePutTransformResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransformResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ResetBatchStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ResetMaxConcurrentTransforms() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentTransforms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ResetMaxPayloadInMb() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxPayloadInMb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ResetTransformInput() {
	_jsii_.InvokeVoid(
		a,
		"resetTransformInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ResetTransformOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetTransformOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ResetTransformResources() {
	_jsii_.InvokeVoid(
		a,
		"resetTransformResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_TransformJobDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

