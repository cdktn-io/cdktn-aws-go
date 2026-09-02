package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslambda/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslambda/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessConfig() TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList
	// Experimental.
	AccessConfigInput() interface{}
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
	EventRecordFormat() *string
	// Experimental.
	SetEventRecordFormat(val *string)
	// Experimental.
	EventRecordFormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
	// Experimental.
	SetInternalValue(val *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigProperty)
	// Experimental.
	SchemaRegistryUri() *string
	// Experimental.
	SetSchemaRegistryUri(val *string)
	// Experimental.
	SchemaRegistryUriInput() *string
	// Experimental.
	SchemaValidationConfig() TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList
	// Experimental.
	SchemaValidationConfigInput() interface{}
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
	PutAccessConfig(value interface{})
	// Experimental.
	PutSchemaValidationConfig(value interface{})
	// Experimental.
	ResetAccessConfig()
	// Experimental.
	ResetEventRecordFormat()
	// Experimental.
	ResetSchemaRegistryUri()
	// Experimental.
	ResetSchemaValidationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference
type jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) AccessConfig() TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList {
	var returns TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) EventRecordFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventRecordFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) EventRecordFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventRecordFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InternalValue() *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigProperty {
	var returns *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaRegistryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaRegistryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaRegistryUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaRegistryUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaValidationConfig() TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList {
	var returns TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList
	_jsii_.Get(
		j,
		"schemaValidationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaValidationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaValidationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lambda.TfEventSourceMapping.SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference_Override(t TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.TfEventSourceMapping.SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetEventRecordFormat(val *string) {
	if err := j.validateSetEventRecordFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventRecordFormat",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetInternalValue(val *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetSchemaRegistryUri(val *string) {
	if err := j.validateSetSchemaRegistryUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaRegistryUri",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) PutAccessConfig(value interface{}) {
	if err := t.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) PutSchemaValidationConfig(value interface{}) {
	if err := t.validatePutSchemaValidationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchemaValidationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetEventRecordFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetEventRecordFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetSchemaRegistryUri() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaRegistryUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetSchemaValidationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaValidationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

