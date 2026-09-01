package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslambda/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslambda/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference interface {
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
	// Experimental.
	ConsumerGroupId() *string
	// Experimental.
	SetConsumerGroupId(val *string)
	// Experimental.
	ConsumerGroupIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty
	// Experimental.
	SetInternalValue(val *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty)
	// Experimental.
	SchemaRegistryConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference
	// Experimental.
	SchemaRegistryConfigInput() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
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
	PutSchemaRegistryConfig(value *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty)
	// Experimental.
	ResetConsumerGroupId()
	// Experimental.
	ResetSchemaRegistryConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference
type jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ConsumerGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ConsumerGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) InternalValue() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty {
	var returns *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) SchemaRegistryConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaRegistryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) SchemaRegistryConfigInput() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty {
	var returns *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
	_jsii_.Get(
		j,
		"schemaRegistryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping.AmazonManagedKafkaEventSourceConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference_Override(a AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping.AmazonManagedKafkaEventSourceConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference)SetConsumerGroupId(val *string) {
	if err := j.validateSetConsumerGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference)SetInternalValue(val *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) PutSchemaRegistryConfig(value *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty) {
	if err := a.validatePutSchemaRegistryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchemaRegistryConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ResetConsumerGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetConsumerGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ResetSchemaRegistryConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaRegistryConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

