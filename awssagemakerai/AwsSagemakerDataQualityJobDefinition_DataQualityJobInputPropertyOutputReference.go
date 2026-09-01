package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchTransformInput() AwsSagemakerDataQualityJobDefinition_BatchTransformInputPropertyOutputReference
	// Experimental.
	BatchTransformInputInput() *AwsSagemakerDataQualityJobDefinition_BatchTransformInputProperty
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
	EndpointInput() AwsSagemakerDataQualityJobDefinition_EndpointInputPropertyOutputReference
	// Experimental.
	EndpointInputInput() *AwsSagemakerDataQualityJobDefinition_EndpointInputProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSagemakerDataQualityJobDefinition_DataQualityJobInputProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerDataQualityJobDefinition_DataQualityJobInputProperty)
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
	PutBatchTransformInput(value *AwsSagemakerDataQualityJobDefinition_BatchTransformInputProperty)
	// Experimental.
	PutEndpointInput(value *AwsSagemakerDataQualityJobDefinition_EndpointInputProperty)
	// Experimental.
	ResetBatchTransformInput()
	// Experimental.
	ResetEndpointInput()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference
type jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) BatchTransformInput() AwsSagemakerDataQualityJobDefinition_BatchTransformInputPropertyOutputReference {
	var returns AwsSagemakerDataQualityJobDefinition_BatchTransformInputPropertyOutputReference
	_jsii_.Get(
		j,
		"batchTransformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) BatchTransformInputInput() *AwsSagemakerDataQualityJobDefinition_BatchTransformInputProperty {
	var returns *AwsSagemakerDataQualityJobDefinition_BatchTransformInputProperty
	_jsii_.Get(
		j,
		"batchTransformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) EndpointInput() AwsSagemakerDataQualityJobDefinition_EndpointInputPropertyOutputReference {
	var returns AwsSagemakerDataQualityJobDefinition_EndpointInputPropertyOutputReference
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) EndpointInputInput() *AwsSagemakerDataQualityJobDefinition_EndpointInputProperty {
	var returns *AwsSagemakerDataQualityJobDefinition_EndpointInputProperty
	_jsii_.Get(
		j,
		"endpointInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) InternalValue() *AwsSagemakerDataQualityJobDefinition_DataQualityJobInputProperty {
	var returns *AwsSagemakerDataQualityJobDefinition_DataQualityJobInputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDataQualityJobDefinition.DataQualityJobInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference_Override(a AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDataQualityJobDefinition.DataQualityJobInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetInternalValue(val *AwsSagemakerDataQualityJobDefinition_DataQualityJobInputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) PutBatchTransformInput(value *AwsSagemakerDataQualityJobDefinition_BatchTransformInputProperty) {
	if err := a.validatePutBatchTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBatchTransformInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) PutEndpointInput(value *AwsSagemakerDataQualityJobDefinition_EndpointInputProperty) {
	if err := a.validatePutEndpointInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ResetBatchTransformInput() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchTransformInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ResetEndpointInput() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

