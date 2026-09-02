package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	// Experimental.
	SetInternalValue(val *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty)
	// Experimental.
	ObjectPath() *string
	// Experimental.
	SetObjectPath(val *string)
	// Experimental.
	ObjectPathInput() *string
	// Experimental.
	PaginationConfig() TfFlow_PaginationConfigPropertyOutputReference
	// Experimental.
	PaginationConfigInput() *TfFlow_PaginationConfigProperty
	// Experimental.
	ParallelismConfig() TfFlow_ParallelismConfigPropertyOutputReference
	// Experimental.
	ParallelismConfigInput() *TfFlow_ParallelismConfigProperty
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
	PutPaginationConfig(value *TfFlow_PaginationConfigProperty)
	// Experimental.
	PutParallelismConfig(value *TfFlow_ParallelismConfigProperty)
	// Experimental.
	ResetPaginationConfig()
	// Experimental.
	ResetParallelismConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
type jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) InternalValue() *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty {
	var returns *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ObjectPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ObjectPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PaginationConfig() TfFlow_PaginationConfigPropertyOutputReference {
	var returns TfFlow_PaginationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"paginationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PaginationConfigInput() *TfFlow_PaginationConfigProperty {
	var returns *TfFlow_PaginationConfigProperty
	_jsii_.Get(
		j,
		"paginationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ParallelismConfig() TfFlow_ParallelismConfigPropertyOutputReference {
	var returns TfFlow_ParallelismConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelismConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ParallelismConfigInput() *TfFlow_ParallelismConfigProperty {
	var returns *TfFlow_ParallelismConfigProperty
	_jsii_.Get(
		j,
		"parallelismConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference_Override(t TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetInternalValue(val *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetObjectPath(val *string) {
	if err := j.validateSetObjectPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectPath",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PutPaginationConfig(value *TfFlow_PaginationConfigProperty) {
	if err := t.validatePutPaginationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPaginationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) PutParallelismConfig(value *TfFlow_ParallelismConfigProperty) {
	if err := t.validatePutParallelismConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParallelismConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ResetPaginationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetPaginationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ResetParallelismConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelismConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

