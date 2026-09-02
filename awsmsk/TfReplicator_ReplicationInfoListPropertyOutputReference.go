package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicator_ReplicationInfoListPropertyOutputReference interface {
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
	ConsumerGroupReplication() TfReplicator_ConsumerGroupReplicationPropertyList
	// Experimental.
	ConsumerGroupReplicationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfReplicator_ReplicationInfoListProperty
	// Experimental.
	SetInternalValue(val *TfReplicator_ReplicationInfoListProperty)
	// Experimental.
	SourceKafkaClusterAlias() *string
	// Experimental.
	SourceKafkaClusterArn() *string
	// Experimental.
	SetSourceKafkaClusterArn(val *string)
	// Experimental.
	SourceKafkaClusterArnInput() *string
	// Experimental.
	TargetCompressionType() *string
	// Experimental.
	SetTargetCompressionType(val *string)
	// Experimental.
	TargetCompressionTypeInput() *string
	// Experimental.
	TargetKafkaClusterAlias() *string
	// Experimental.
	TargetKafkaClusterArn() *string
	// Experimental.
	SetTargetKafkaClusterArn(val *string)
	// Experimental.
	TargetKafkaClusterArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TopicReplication() TfReplicator_TopicReplicationPropertyList
	// Experimental.
	TopicReplicationInput() interface{}
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
	PutConsumerGroupReplication(value interface{})
	// Experimental.
	PutTopicReplication(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfReplicator_ReplicationInfoListPropertyOutputReference
type jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) ConsumerGroupReplication() TfReplicator_ConsumerGroupReplicationPropertyList {
	var returns TfReplicator_ConsumerGroupReplicationPropertyList
	_jsii_.Get(
		j,
		"consumerGroupReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) ConsumerGroupReplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumerGroupReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) InternalValue() *TfReplicator_ReplicationInfoListProperty {
	var returns *TfReplicator_ReplicationInfoListProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) SourceKafkaClusterAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKafkaClusterAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) SourceKafkaClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKafkaClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) SourceKafkaClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKafkaClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TargetCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TargetCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TargetKafkaClusterAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetKafkaClusterAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TargetKafkaClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetKafkaClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TargetKafkaClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetKafkaClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TopicReplication() TfReplicator_TopicReplicationPropertyList {
	var returns TfReplicator_TopicReplicationPropertyList
	_jsii_.Get(
		j,
		"topicReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) TopicReplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicReplicationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfReplicator_ReplicationInfoListPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfReplicator_ReplicationInfoListPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfReplicator_ReplicationInfoListPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.TfReplicator.ReplicationInfoListPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfReplicator_ReplicationInfoListPropertyOutputReference_Override(t TfReplicator_ReplicationInfoListPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.TfReplicator.ReplicationInfoListPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetInternalValue(val *TfReplicator_ReplicationInfoListProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetSourceKafkaClusterArn(val *string) {
	if err := j.validateSetSourceKafkaClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceKafkaClusterArn",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetTargetCompressionType(val *string) {
	if err := j.validateSetTargetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCompressionType",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetTargetKafkaClusterArn(val *string) {
	if err := j.validateSetTargetKafkaClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetKafkaClusterArn",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) PutConsumerGroupReplication(value interface{}) {
	if err := t.validatePutConsumerGroupReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConsumerGroupReplication",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) PutTopicReplication(value interface{}) {
	if err := t.validatePutTopicReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTopicReplication",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfReplicator_ReplicationInfoListPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

