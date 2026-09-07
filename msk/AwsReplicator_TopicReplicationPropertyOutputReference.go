package msk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/msk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/msk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsReplicator_TopicReplicationPropertyOutputReference interface {
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
	StartingPosition() AwsReplicator_StartingPositionPropertyOutputReference
	// Experimental.
	StartingPositionInput() *AwsReplicator_StartingPositionProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TopicNameConfiguration() AwsReplicator_TopicNameConfigurationPropertyOutputReference
	// Experimental.
	TopicNameConfigurationInput() *AwsReplicator_TopicNameConfigurationProperty
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
	PutStartingPosition(value *AwsReplicator_StartingPositionProperty)
	// Experimental.
	PutTopicNameConfiguration(value *AwsReplicator_TopicNameConfigurationProperty)
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

// The jsii proxy struct for AwsReplicator_TopicReplicationPropertyOutputReference
type jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) CopyAccessControlListsForTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyAccessControlListsForTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) CopyAccessControlListsForTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyAccessControlListsForTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) CopyTopicConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTopicConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) CopyTopicConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTopicConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) DetectAndCopyNewTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) DetectAndCopyNewTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) StartingPosition() AwsReplicator_StartingPositionPropertyOutputReference {
	var returns AwsReplicator_StartingPositionPropertyOutputReference
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) StartingPositionInput() *AwsReplicator_StartingPositionProperty {
	var returns *AwsReplicator_StartingPositionProperty
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TopicNameConfiguration() AwsReplicator_TopicNameConfigurationPropertyOutputReference {
	var returns AwsReplicator_TopicNameConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"topicNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TopicNameConfigurationInput() *AwsReplicator_TopicNameConfigurationProperty {
	var returns *AwsReplicator_TopicNameConfigurationProperty
	_jsii_.Get(
		j,
		"topicNameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TopicsToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TopicsToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TopicsToReplicate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToReplicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) TopicsToReplicateInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToReplicateInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsReplicator_TopicReplicationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsReplicator_TopicReplicationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsReplicator_TopicReplicationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.AwsReplicator.TopicReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsReplicator_TopicReplicationPropertyOutputReference_Override(a AwsReplicator_TopicReplicationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.AwsReplicator.TopicReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetCopyAccessControlListsForTopics(val interface{}) {
	if err := j.validateSetCopyAccessControlListsForTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyAccessControlListsForTopics",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetCopyTopicConfigurations(val interface{}) {
	if err := j.validateSetCopyTopicConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTopicConfigurations",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetDetectAndCopyNewTopics(val interface{}) {
	if err := j.validateSetDetectAndCopyNewTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectAndCopyNewTopics",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetTopicsToExclude(val *[]*string) {
	if err := j.validateSetTopicsToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicsToExclude",
		val,
	)
}

func (j *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference)SetTopicsToReplicate(val *[]*string) {
	if err := j.validateSetTopicsToReplicateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicsToReplicate",
		val,
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) PutStartingPosition(value *AwsReplicator_StartingPositionProperty) {
	if err := a.validatePutStartingPositionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStartingPosition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) PutTopicNameConfiguration(value *AwsReplicator_TopicNameConfigurationProperty) {
	if err := a.validatePutTopicNameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTopicNameConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ResetCopyAccessControlListsForTopics() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyAccessControlListsForTopics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ResetCopyTopicConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTopicConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ResetDetectAndCopyNewTopics() {
	_jsii_.InvokeVoid(
		a,
		"resetDetectAndCopyNewTopics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ResetTopicNameConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTopicNameConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ResetTopicsToExclude() {
	_jsii_.InvokeVoid(
		a,
		"resetTopicsToExclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsReplicator_TopicReplicationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

