package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchTransformInput() AwsSagemakerMonitoringSchedule_BatchTransformInputPropertyOutputReference
	// Experimental.
	BatchTransformInputInput() *AwsSagemakerMonitoringSchedule_BatchTransformInputProperty
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
	EndpointInput() AwsSagemakerMonitoringSchedule_EndpointInputPropertyOutputReference
	// Experimental.
	EndpointInputInput() *AwsSagemakerMonitoringSchedule_EndpointInputProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty)
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
	PutBatchTransformInput(value *AwsSagemakerMonitoringSchedule_BatchTransformInputProperty)
	// Experimental.
	PutEndpointInput(value *AwsSagemakerMonitoringSchedule_EndpointInputProperty)
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

// The jsii proxy struct for AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference
type jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) BatchTransformInput() AwsSagemakerMonitoringSchedule_BatchTransformInputPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_BatchTransformInputPropertyOutputReference
	_jsii_.Get(
		j,
		"batchTransformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) BatchTransformInputInput() *AwsSagemakerMonitoringSchedule_BatchTransformInputProperty {
	var returns *AwsSagemakerMonitoringSchedule_BatchTransformInputProperty
	_jsii_.Get(
		j,
		"batchTransformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) EndpointInput() AwsSagemakerMonitoringSchedule_EndpointInputPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_EndpointInputPropertyOutputReference
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) EndpointInputInput() *AwsSagemakerMonitoringSchedule_EndpointInputProperty {
	var returns *AwsSagemakerMonitoringSchedule_EndpointInputProperty
	_jsii_.Get(
		j,
		"endpointInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) InternalValue() *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerMonitoringSchedule.MonitoringInputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference_Override(a AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerMonitoringSchedule.MonitoringInputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetInternalValue(val *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) PutBatchTransformInput(value *AwsSagemakerMonitoringSchedule_BatchTransformInputProperty) {
	if err := a.validatePutBatchTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBatchTransformInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) PutEndpointInput(value *AwsSagemakerMonitoringSchedule_EndpointInputProperty) {
	if err := a.validatePutEndpointInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) ResetBatchTransformInput() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchTransformInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) ResetEndpointInput() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

