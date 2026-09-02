package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_SourceFlowConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiVersion() *string
	// Experimental.
	SetApiVersion(val *string)
	// Experimental.
	ApiVersionInput() *string
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
	ConnectorProfileName() *string
	// Experimental.
	SetConnectorProfileName(val *string)
	// Experimental.
	ConnectorProfileNameInput() *string
	// Experimental.
	ConnectorType() *string
	// Experimental.
	SetConnectorType(val *string)
	// Experimental.
	ConnectorTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncrementalPullConfig() TfFlow_IncrementalPullConfigPropertyOutputReference
	// Experimental.
	IncrementalPullConfigInput() *TfFlow_IncrementalPullConfigProperty
	// Experimental.
	InternalValue() *TfFlow_SourceFlowConfigProperty
	// Experimental.
	SetInternalValue(val *TfFlow_SourceFlowConfigProperty)
	// Experimental.
	SourceConnectorProperties() TfFlow_SourceConnectorPropertiesPropertyOutputReference
	// Experimental.
	SourceConnectorPropertiesInput() *TfFlow_SourceConnectorPropertiesProperty
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
	PutIncrementalPullConfig(value *TfFlow_IncrementalPullConfigProperty)
	// Experimental.
	PutSourceConnectorProperties(value *TfFlow_SourceConnectorPropertiesProperty)
	// Experimental.
	ResetApiVersion()
	// Experimental.
	ResetConnectorProfileName()
	// Experimental.
	ResetIncrementalPullConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_SourceFlowConfigPropertyOutputReference
type jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ConnectorProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ConnectorProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorProfileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ConnectorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) IncrementalPullConfig() TfFlow_IncrementalPullConfigPropertyOutputReference {
	var returns TfFlow_IncrementalPullConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"incrementalPullConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) IncrementalPullConfigInput() *TfFlow_IncrementalPullConfigProperty {
	var returns *TfFlow_IncrementalPullConfigProperty
	_jsii_.Get(
		j,
		"incrementalPullConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) InternalValue() *TfFlow_SourceFlowConfigProperty {
	var returns *TfFlow_SourceFlowConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) SourceConnectorProperties() TfFlow_SourceConnectorPropertiesPropertyOutputReference {
	var returns TfFlow_SourceConnectorPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceConnectorProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) SourceConnectorPropertiesInput() *TfFlow_SourceConnectorPropertiesProperty {
	var returns *TfFlow_SourceConnectorPropertiesProperty
	_jsii_.Get(
		j,
		"sourceConnectorPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_SourceFlowConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_SourceFlowConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_SourceFlowConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.SourceFlowConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_SourceFlowConfigPropertyOutputReference_Override(t TfFlow_SourceFlowConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.SourceFlowConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetApiVersion(val *string) {
	if err := j.validateSetApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiVersion",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetConnectorProfileName(val *string) {
	if err := j.validateSetConnectorProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorProfileName",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetConnectorType(val *string) {
	if err := j.validateSetConnectorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorType",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetInternalValue(val *TfFlow_SourceFlowConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) PutIncrementalPullConfig(value *TfFlow_IncrementalPullConfigProperty) {
	if err := t.validatePutIncrementalPullConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIncrementalPullConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) PutSourceConnectorProperties(value *TfFlow_SourceConnectorPropertiesProperty) {
	if err := t.validatePutSourceConnectorPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceConnectorProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ResetApiVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetApiVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ResetConnectorProfileName() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectorProfileName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ResetIncrementalPullConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetIncrementalPullConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

