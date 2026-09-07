package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference interface {
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
	DataTransferApi() *string
	// Experimental.
	SetDataTransferApi(val *string)
	// Experimental.
	DataTransferApiInput() *string
	// Experimental.
	ErrorHandlingConfig() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigPropertyOutputReference
	// Experimental.
	ErrorHandlingConfigInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	IdFieldNames() *[]*string
	// Experimental.
	SetIdFieldNames(val *[]*string)
	// Experimental.
	IdFieldNamesInput() *[]*string
	// Experimental.
	InternalValue() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	// Experimental.
	SetInternalValue(val *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty)
	// Experimental.
	Object() *string
	// Experimental.
	SetObject(val *string)
	// Experimental.
	ObjectInput() *string
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
	PutErrorHandlingConfig(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigProperty)
	// Experimental.
	ResetDataTransferApi()
	// Experimental.
	ResetErrorHandlingConfig()
	// Experimental.
	ResetIdFieldNames()
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

// The jsii proxy struct for AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference
type jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) DataTransferApi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) DataTransferApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ErrorHandlingConfig() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigPropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"errorHandlingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ErrorHandlingConfigInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigProperty
	_jsii_.Get(
		j,
		"errorHandlingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) IdFieldNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) IdFieldNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) InternalValue() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) Object() *string {
	var returns *string
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) WriteOperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) WriteOperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference_Override(a AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetDataTransferApi(val *string) {
	if err := j.validateSetDataTransferApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTransferApi",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetIdFieldNames(val *[]*string) {
	if err := j.validateSetIdFieldNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idFieldNames",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetInternalValue(val *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetObject(val *string) {
	if err := j.validateSetObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"object",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference)SetWriteOperationType(val *string) {
	if err := j.validateSetWriteOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeOperationType",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) PutErrorHandlingConfig(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigProperty) {
	if err := a.validatePutErrorHandlingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putErrorHandlingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ResetDataTransferApi() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTransferApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ResetErrorHandlingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorHandlingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ResetIdFieldNames() {
	_jsii_.InvokeVoid(
		a,
		"resetIdFieldNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ResetWriteOperationType() {
	_jsii_.InvokeVoid(
		a,
		"resetWriteOperationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

