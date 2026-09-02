package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicator_TopicReplicationPropertyOutputReference interface {
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
	CopyAccessControlListsForTopics() interface{}
	// Experimental.
	SetCopyAccessControlListsForTopics(val interface{})
	// Experimental.
	CopyAccessControlListsForTopicsInput() interface{}
	// Experimental.
	CopyTopicConfigurations() interface{}
	// Experimental.
	SetCopyTopicConfigurations(val interface{})
	// Experimental.
	CopyTopicConfigurationsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DetectAndCopyNewTopics() interface{}
	// Experimental.
	SetDetectAndCopyNewTopics(val interface{})
	// Experimental.
	DetectAndCopyNewTopicsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	StartingPosition() TfReplicator_StartingPositionPropertyOutputReference
	// Experimental.
	StartingPositionInput() *TfReplicator_StartingPositionProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TopicNameConfiguration() TfReplicator_TopicNameConfigurationPropertyOutputReference
	// Experimental.
	TopicNameConfigurationInput() *TfReplicator_TopicNameConfigurationProperty
	// Experimental.
	TopicsToExclude() *[]*string
	// Experimental.
	SetTopicsToExclude(val *[]*string)
	// Experimental.
	TopicsToExcludeInput() *[]*string
	// Experimental.
	TopicsToReplicate() *[]*string
	// Experimental.
	SetTopicsToReplicate(val *[]*string)
	// Experimental.
	TopicsToReplicateInput() *[]*string
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
	PutStartingPosition(value *TfReplicator_StartingPositionProperty)
	// Experimental.
	PutTopicNameConfiguration(value *TfReplicator_TopicNameConfigurationProperty)
	// Experimental.
	ResetCopyAccessControlListsForTopics()
	// Experimental.
	ResetCopyTopicConfigurations()
	// Experimental.
	ResetDetectAndCopyNewTopics()
	// Experimental.
	ResetStartingPosition()
	// Experimental.
	ResetTopicNameConfiguration()
	// Experimental.
	ResetTopicsToExclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfReplicator_TopicReplicationPropertyOutputReference
type jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) CopyAccessControlListsForTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyAccessControlListsForTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) CopyAccessControlListsForTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyAccessControlListsForTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) CopyTopicConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTopicConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) CopyTopicConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTopicConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) DetectAndCopyNewTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) DetectAndCopyNewTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) StartingPosition() TfReplicator_StartingPositionPropertyOutputReference {
	var returns TfReplicator_StartingPositionPropertyOutputReference
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) StartingPositionInput() *TfReplicator_StartingPositionProperty {
	var returns *TfReplicator_StartingPositionProperty
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TopicNameConfiguration() TfReplicator_TopicNameConfigurationPropertyOutputReference {
	var returns TfReplicator_TopicNameConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"topicNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TopicNameConfigurationInput() *TfReplicator_TopicNameConfigurationProperty {
	var returns *TfReplicator_TopicNameConfigurationProperty
	_jsii_.Get(
		j,
		"topicNameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TopicsToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TopicsToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TopicsToReplicate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToReplicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) TopicsToReplicateInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToReplicateInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfReplicator_TopicReplicationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfReplicator_TopicReplicationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfReplicator_TopicReplicationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.TfReplicator.TopicReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfReplicator_TopicReplicationPropertyOutputReference_Override(t TfReplicator_TopicReplicationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.TfReplicator.TopicReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetCopyAccessControlListsForTopics(val interface{}) {
	if err := j.validateSetCopyAccessControlListsForTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyAccessControlListsForTopics",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetCopyTopicConfigurations(val interface{}) {
	if err := j.validateSetCopyTopicConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTopicConfigurations",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetDetectAndCopyNewTopics(val interface{}) {
	if err := j.validateSetDetectAndCopyNewTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectAndCopyNewTopics",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetTopicsToExclude(val *[]*string) {
	if err := j.validateSetTopicsToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicsToExclude",
		val,
	)
}

func (j *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference)SetTopicsToReplicate(val *[]*string) {
	if err := j.validateSetTopicsToReplicateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicsToReplicate",
		val,
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) PutStartingPosition(value *TfReplicator_StartingPositionProperty) {
	if err := t.validatePutStartingPositionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStartingPosition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) PutTopicNameConfiguration(value *TfReplicator_TopicNameConfigurationProperty) {
	if err := t.validatePutTopicNameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTopicNameConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ResetCopyAccessControlListsForTopics() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyAccessControlListsForTopics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ResetCopyTopicConfigurations() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTopicConfigurations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ResetDetectAndCopyNewTopics() {
	_jsii_.InvokeVoid(
		t,
		"resetDetectAndCopyNewTopics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		t,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ResetTopicNameConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTopicNameConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ResetTopicsToExclude() {
	_jsii_.InvokeVoid(
		t,
		"resetTopicsToExclude",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfReplicator_TopicReplicationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

