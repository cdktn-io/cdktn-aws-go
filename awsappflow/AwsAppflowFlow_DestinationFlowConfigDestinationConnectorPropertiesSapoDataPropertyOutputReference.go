package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference interface {
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
	ErrorHandlingConfig() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigPropertyOutputReference
	// Experimental.
	ErrorHandlingConfigInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	IdFieldNames() *[]*string
	// Experimental.
	SetIdFieldNames(val *[]*string)
	// Experimental.
	IdFieldNamesInput() *[]*string
	// Experimental.
	InternalValue() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty)
	// Experimental.
	ObjectPath() *string
	// Experimental.
	SetObjectPath(val *string)
	// Experimental.
	ObjectPathInput() *string
	// Experimental.
	SuccessResponseHandlingConfig() AwsAppflowFlow_SuccessResponseHandlingConfigPropertyOutputReference
	// Experimental.
	SuccessResponseHandlingConfigInput() *AwsAppflowFlow_SuccessResponseHandlingConfigProperty
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
	PutErrorHandlingConfig(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty)
	// Experimental.
	PutSuccessResponseHandlingConfig(value *AwsAppflowFlow_SuccessResponseHandlingConfigProperty)
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

// The jsii proxy struct for AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
type jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ErrorHandlingConfig() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigPropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"errorHandlingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ErrorHandlingConfigInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty
	_jsii_.Get(
		j,
		"errorHandlingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) IdFieldNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) IdFieldNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) InternalValue() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ObjectPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ObjectPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) SuccessResponseHandlingConfig() AwsAppflowFlow_SuccessResponseHandlingConfigPropertyOutputReference {
	var returns AwsAppflowFlow_SuccessResponseHandlingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"successResponseHandlingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) SuccessResponseHandlingConfigInput() *AwsAppflowFlow_SuccessResponseHandlingConfigProperty {
	var returns *AwsAppflowFlow_SuccessResponseHandlingConfigProperty
	_jsii_.Get(
		j,
		"successResponseHandlingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) WriteOperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) WriteOperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference_Override(a AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetIdFieldNames(val *[]*string) {
	if err := j.validateSetIdFieldNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idFieldNames",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetInternalValue(val *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetObjectPath(val *string) {
	if err := j.validateSetObjectPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectPath",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference)SetWriteOperationType(val *string) {
	if err := j.validateSetWriteOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeOperationType",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) PutErrorHandlingConfig(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfigProperty) {
	if err := a.validatePutErrorHandlingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putErrorHandlingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) PutSuccessResponseHandlingConfig(value *AwsAppflowFlow_SuccessResponseHandlingConfigProperty) {
	if err := a.validatePutSuccessResponseHandlingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccessResponseHandlingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetErrorHandlingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorHandlingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetIdFieldNames() {
	_jsii_.InvokeVoid(
		a,
		"resetIdFieldNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetSuccessResponseHandlingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessResponseHandlingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ResetWriteOperationType() {
	_jsii_.InvokeVoid(
		a,
		"resetWriteOperationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

