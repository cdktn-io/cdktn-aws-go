package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BucketName() *string
	// Experimental.
	SetBucketName(val *string)
	// Experimental.
	BucketNameInput() *string
	// Experimental.
	BucketPrefix() *string
	// Experimental.
	SetBucketPrefix(val *string)
	// Experimental.
	BucketPrefixInput() *string
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
	FailOnFirstDestinationError() interface{}
	// Experimental.
	SetFailOnFirstDestinationError(val interface{})
	// Experimental.
	FailOnFirstDestinationErrorInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigProperty
	// Experimental.
	SetInternalValue(val *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigProperty)
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
	ResetBucketName()
	// Experimental.
	ResetBucketPrefix()
	// Experimental.
	ResetFailOnFirstDestinationError()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference
type jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) FailOnFirstDestinationError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstDestinationError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) FailOnFirstDestinationErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstDestinationErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) InternalValue() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference_Override(t TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetBucketPrefix(val *string) {
	if err := j.validateSetBucketPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetFailOnFirstDestinationError(val interface{}) {
	if err := j.validateSetFailOnFirstDestinationErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOnFirstDestinationError",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetInternalValue(val *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		t,
		"resetBucketName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) ResetFailOnFirstDestinationError() {
	_jsii_.InvokeVoid(
		t,
		"resetFailOnFirstDestinationError",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoErrorHandlingConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

