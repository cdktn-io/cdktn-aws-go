package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMonitoringSchedule_MonitoringInputsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchTransformInput() TfMonitoringSchedule_BatchTransformInputPropertyOutputReference
	// Experimental.
	BatchTransformInputInput() *TfMonitoringSchedule_BatchTransformInputProperty
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
	EndpointInput() TfMonitoringSchedule_EndpointInputPropertyOutputReference
	// Experimental.
	EndpointInputInput() *TfMonitoringSchedule_EndpointInputProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfMonitoringSchedule_MonitoringInputsProperty
	// Experimental.
	SetInternalValue(val *TfMonitoringSchedule_MonitoringInputsProperty)
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
	PutBatchTransformInput(value *TfMonitoringSchedule_BatchTransformInputProperty)
	// Experimental.
	PutEndpointInput(value *TfMonitoringSchedule_EndpointInputProperty)
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

// The jsii proxy struct for TfMonitoringSchedule_MonitoringInputsPropertyOutputReference
type jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) BatchTransformInput() TfMonitoringSchedule_BatchTransformInputPropertyOutputReference {
	var returns TfMonitoringSchedule_BatchTransformInputPropertyOutputReference
	_jsii_.Get(
		j,
		"batchTransformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) BatchTransformInputInput() *TfMonitoringSchedule_BatchTransformInputProperty {
	var returns *TfMonitoringSchedule_BatchTransformInputProperty
	_jsii_.Get(
		j,
		"batchTransformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) EndpointInput() TfMonitoringSchedule_EndpointInputPropertyOutputReference {
	var returns TfMonitoringSchedule_EndpointInputPropertyOutputReference
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) EndpointInputInput() *TfMonitoringSchedule_EndpointInputProperty {
	var returns *TfMonitoringSchedule_EndpointInputProperty
	_jsii_.Get(
		j,
		"endpointInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) InternalValue() *TfMonitoringSchedule_MonitoringInputsProperty {
	var returns *TfMonitoringSchedule_MonitoringInputsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMonitoringSchedule_MonitoringInputsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMonitoringSchedule_MonitoringInputsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMonitoringSchedule_MonitoringInputsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringInputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMonitoringSchedule_MonitoringInputsPropertyOutputReference_Override(t TfMonitoringSchedule_MonitoringInputsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringInputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetInternalValue(val *TfMonitoringSchedule_MonitoringInputsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) PutBatchTransformInput(value *TfMonitoringSchedule_BatchTransformInputProperty) {
	if err := t.validatePutBatchTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBatchTransformInput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) PutEndpointInput(value *TfMonitoringSchedule_EndpointInputProperty) {
	if err := t.validatePutEndpointInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEndpointInput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) ResetBatchTransformInput() {
	_jsii_.InvokeVoid(
		t,
		"resetBatchTransformInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) ResetEndpointInput() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringInputsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

