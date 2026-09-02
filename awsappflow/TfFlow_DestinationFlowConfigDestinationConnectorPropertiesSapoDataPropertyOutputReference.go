package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference interface {
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
	ErrorHandlingConfig() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigPropertyOutputReference
	// Experimental.
	ErrorHandlingConfigInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	IdFieldNames() *[]*string
	// Experimental.
	SetIdFieldNames(val *[]*string)
	// Experimental.
	IdFieldNamesInput() *[]*string
	// Experimental.
	InternalValue() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	// Experimental.
	SetInternalValue(val *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty)
	// Experimental.
	ObjectPath() *string
	// Experimental.
	SetObjectPath(val *string)
	// Experimental.
	ObjectPathInput() *string
	// Experimental.
	SuccessResponseHandlingConfig() TfFlow_SuccessResponseHandlingConfigPropertyOutputReference
	// Experimental.
	SuccessResponseHandlingConfigInput() *TfFlow_SuccessResponseHandlingConfigProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WriteOperationType() *string
	// Experimental.
	SetWriteOperationType(val *string)
	// Experimental.
	WriteOperationTypeInput() *string
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
	PutErrorHandlingConfig(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty)
	// Experimental.
	PutSuccessResponseHandlingConfig(value *TfFlow_SuccessResponseHandlingConfigProperty)
	// Experimental.
	ResetErrorHandlingConfig()
	// Experimental.
	ResetIdFieldNames()
	// Experimental.
	ResetSuccessResponseHandlingConfig()
	// Experimental.
	ResetWriteOperationType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
type jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ErrorHandlingConfig() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigPropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"errorHandlingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ErrorHandlingConfigInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty
	_jsii_.Get(
		j,
		"errorHandlingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) IdFieldNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) IdFieldNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) InternalValue() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ObjectPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ObjectPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) SuccessResponseHandlingConfig() TfFlow_SuccessResponseHandlingConfigPropertyOutputReference {
	var returns TfFlow_SuccessResponseHandlingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"successResponseHandlingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) SuccessResponseHandlingConfigInput() *TfFlow_SuccessResponseHandlingConfigProperty {
	var returns *TfFlow_SuccessResponseHandlingConfigProperty
	_jsii_.Get(
		j,
		"successResponseHandlingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) WriteOperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) WriteOperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference_Override(t TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetIdFieldNames(val *[]*string) {
	if err := j.validateSetIdFieldNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idFieldNames",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetInternalValue(val *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetObjectPath(val *string) {
	if err := j.validateSetObjectPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectPath",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetWriteOperationType(val *string) {
	if err := j.validateSetWriteOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeOperationType",
		val,
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) PutErrorHandlingConfig(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty) {
	if err := t.validatePutErrorHandlingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putErrorHandlingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) PutSuccessResponseHandlingConfig(value *TfFlow_SuccessResponseHandlingConfigProperty) {
	if err := t.validatePutSuccessResponseHandlingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSuccessResponseHandlingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetErrorHandlingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorHandlingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetIdFieldNames() {
	_jsii_.InvokeVoid(
		t,
		"resetIdFieldNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetSuccessResponseHandlingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessResponseHandlingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetWriteOperationType() {
	_jsii_.InvokeVoid(
		t,
		"resetWriteOperationType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

