package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslambda/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslambda/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessConfig() TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList
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
	InternalValue() *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
	// Experimental.
	SetInternalValue(val *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty)
	// Experimental.
	SchemaRegistryUri() *string
	// Experimental.
	SetSchemaRegistryUri(val *string)
	// Experimental.
	SchemaRegistryUriInput() *string
	// Experimental.
	SchemaValidationConfig() TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList
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

// The jsii proxy struct for TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference
type jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) AccessConfig() TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList {
	var returns TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) EventRecordFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventRecordFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) EventRecordFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventRecordFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InternalValue() *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty {
	var returns *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaRegistryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaRegistryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaRegistryUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaRegistryUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaValidationConfig() TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList {
	var returns TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList
	_jsii_.Get(
		j,
		"schemaValidationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaValidationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaValidationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lambda.TfEventSourceMapping.AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference_Override(t TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.TfEventSourceMapping.AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetEventRecordFormat(val *string) {
	if err := j.validateSetEventRecordFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventRecordFormat",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetInternalValue(val *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetSchemaRegistryUri(val *string) {
	if err := j.validateSetSchemaRegistryUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaRegistryUri",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) PutAccessConfig(value interface{}) {
	if err := t.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) PutSchemaValidationConfig(value interface{}) {
	if err := t.validatePutSchemaValidationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchemaValidationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetEventRecordFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetEventRecordFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetSchemaRegistryUri() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaRegistryUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetSchemaValidationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaValidationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

