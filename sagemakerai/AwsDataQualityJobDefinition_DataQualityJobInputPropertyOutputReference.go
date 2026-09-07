package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchTransformInput() AwsDataQualityJobDefinition_BatchTransformInputPropertyOutputReference
	// Experimental.
	BatchTransformInputInput() *AwsDataQualityJobDefinition_BatchTransformInputProperty
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
	EndpointInput() AwsDataQualityJobDefinition_EndpointInputPropertyOutputReference
	// Experimental.
	EndpointInputInput() *AwsDataQualityJobDefinition_EndpointInputProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDataQualityJobDefinition_DataQualityJobInputProperty
	// Experimental.
	SetInternalValue(val *AwsDataQualityJobDefinition_DataQualityJobInputProperty)
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
	PutBatchTransformInput(value *AwsDataQualityJobDefinition_BatchTransformInputProperty)
	// Experimental.
	PutEndpointInput(value *AwsDataQualityJobDefinition_EndpointInputProperty)
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

// The jsii proxy struct for AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference
type jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) BatchTransformInput() AwsDataQualityJobDefinition_BatchTransformInputPropertyOutputReference {
	var returns AwsDataQualityJobDefinition_BatchTransformInputPropertyOutputReference
	_jsii_.Get(
		j,
		"batchTransformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) BatchTransformInputInput() *AwsDataQualityJobDefinition_BatchTransformInputProperty {
	var returns *AwsDataQualityJobDefinition_BatchTransformInputProperty
	_jsii_.Get(
		j,
		"batchTransformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) EndpointInput() AwsDataQualityJobDefinition_EndpointInputPropertyOutputReference {
	var returns AwsDataQualityJobDefinition_EndpointInputPropertyOutputReference
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) EndpointInputInput() *AwsDataQualityJobDefinition_EndpointInputProperty {
	var returns *AwsDataQualityJobDefinition_EndpointInputProperty
	_jsii_.Get(
		j,
		"endpointInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) InternalValue() *AwsDataQualityJobDefinition_DataQualityJobInputProperty {
	var returns *AwsDataQualityJobDefinition_DataQualityJobInputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDataQualityJobDefinition.DataQualityJobInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference_Override(a AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDataQualityJobDefinition.DataQualityJobInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetInternalValue(val *AwsDataQualityJobDefinition_DataQualityJobInputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) PutBatchTransformInput(value *AwsDataQualityJobDefinition_BatchTransformInputProperty) {
	if err := a.validatePutBatchTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBatchTransformInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) PutEndpointInput(value *AwsDataQualityJobDefinition_EndpointInputProperty) {
	if err := a.validatePutEndpointInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ResetBatchTransformInput() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchTransformInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ResetEndpointInput() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataQualityJobDefinition_DataQualityJobInputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

