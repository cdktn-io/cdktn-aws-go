package awssesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRuleSet_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddHeader() TfRuleSet_AddHeaderPropertyList
	// Experimental.
	AddHeaderInput() interface{}
	// Experimental.
	Archive() TfRuleSet_ArchivePropertyList
	// Experimental.
	ArchiveInput() interface{}
	// Experimental.
	Bounce() TfRuleSet_BouncePropertyList
	// Experimental.
	BounceInput() interface{}
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
	DeliverToMailbox() TfRuleSet_DeliverToMailboxPropertyList
	// Experimental.
	DeliverToMailboxInput() interface{}
	// Experimental.
	DeliverToQBusiness() TfRuleSet_DeliverToQBusinessPropertyList
	// Experimental.
	DeliverToQBusinessInput() interface{}
	// Experimental.
	Drop() TfRuleSet_DropPropertyList
	// Experimental.
	DropInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	InvokeLambda() TfRuleSet_InvokeLambdaPropertyList
	// Experimental.
	InvokeLambdaInput() interface{}
	// Experimental.
	PublishToSns() TfRuleSet_PublishToSnsPropertyList
	// Experimental.
	PublishToSnsInput() interface{}
	// Experimental.
	Relay() TfRuleSet_RelayPropertyList
	// Experimental.
	RelayInput() interface{}
	// Experimental.
	ReplaceRecipient() TfRuleSet_ReplaceRecipientPropertyList
	// Experimental.
	ReplaceRecipientInput() interface{}
	// Experimental.
	Send() TfRuleSet_SendPropertyList
	// Experimental.
	SendInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WriteToS3() TfRuleSet_WriteToS3PropertyList
	// Experimental.
	WriteToS3Input() interface{}
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
	PutAddHeader(value interface{})
	// Experimental.
	PutArchive(value interface{})
	// Experimental.
	PutBounce(value interface{})
	// Experimental.
	PutDeliverToMailbox(value interface{})
	// Experimental.
	PutDeliverToQBusiness(value interface{})
	// Experimental.
	PutDrop(value interface{})
	// Experimental.
	PutInvokeLambda(value interface{})
	// Experimental.
	PutPublishToSns(value interface{})
	// Experimental.
	PutRelay(value interface{})
	// Experimental.
	PutReplaceRecipient(value interface{})
	// Experimental.
	PutSend(value interface{})
	// Experimental.
	PutWriteToS3(value interface{})
	// Experimental.
	ResetAddHeader()
	// Experimental.
	ResetArchive()
	// Experimental.
	ResetBounce()
	// Experimental.
	ResetDeliverToMailbox()
	// Experimental.
	ResetDeliverToQBusiness()
	// Experimental.
	ResetDrop()
	// Experimental.
	ResetInvokeLambda()
	// Experimental.
	ResetPublishToSns()
	// Experimental.
	ResetRelay()
	// Experimental.
	ResetReplaceRecipient()
	// Experimental.
	ResetSend()
	// Experimental.
	ResetWriteToS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRuleSet_ActionPropertyOutputReference
type jsiiProxy_TfRuleSet_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) AddHeader() TfRuleSet_AddHeaderPropertyList {
	var returns TfRuleSet_AddHeaderPropertyList
	_jsii_.Get(
		j,
		"addHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) AddHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) Archive() TfRuleSet_ArchivePropertyList {
	var returns TfRuleSet_ArchivePropertyList
	_jsii_.Get(
		j,
		"archive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ArchiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) Bounce() TfRuleSet_BouncePropertyList {
	var returns TfRuleSet_BouncePropertyList
	_jsii_.Get(
		j,
		"bounce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) BounceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bounceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) DeliverToMailbox() TfRuleSet_DeliverToMailboxPropertyList {
	var returns TfRuleSet_DeliverToMailboxPropertyList
	_jsii_.Get(
		j,
		"deliverToMailbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) DeliverToMailboxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToMailboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) DeliverToQBusiness() TfRuleSet_DeliverToQBusinessPropertyList {
	var returns TfRuleSet_DeliverToQBusinessPropertyList
	_jsii_.Get(
		j,
		"deliverToQBusiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) DeliverToQBusinessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToQBusinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) Drop() TfRuleSet_DropPropertyList {
	var returns TfRuleSet_DropPropertyList
	_jsii_.Get(
		j,
		"drop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) DropInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) InvokeLambda() TfRuleSet_InvokeLambdaPropertyList {
	var returns TfRuleSet_InvokeLambdaPropertyList
	_jsii_.Get(
		j,
		"invokeLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) InvokeLambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invokeLambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PublishToSns() TfRuleSet_PublishToSnsPropertyList {
	var returns TfRuleSet_PublishToSnsPropertyList
	_jsii_.Get(
		j,
		"publishToSns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PublishToSnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishToSnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) Relay() TfRuleSet_RelayPropertyList {
	var returns TfRuleSet_RelayPropertyList
	_jsii_.Get(
		j,
		"relay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) RelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ReplaceRecipient() TfRuleSet_ReplaceRecipientPropertyList {
	var returns TfRuleSet_ReplaceRecipientPropertyList
	_jsii_.Get(
		j,
		"replaceRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ReplaceRecipientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) Send() TfRuleSet_SendPropertyList {
	var returns TfRuleSet_SendPropertyList
	_jsii_.Get(
		j,
		"send",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) SendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) WriteToS3() TfRuleSet_WriteToS3PropertyList {
	var returns TfRuleSet_WriteToS3PropertyList
	_jsii_.Get(
		j,
		"writeToS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) WriteToS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeToS3Input",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRuleSet_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRuleSet_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRuleSet_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRuleSet_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfRuleSet.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRuleSet_ActionPropertyOutputReference_Override(t TfRuleSet_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfRuleSet.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutAddHeader(value interface{}) {
	if err := t.validatePutAddHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAddHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutArchive(value interface{}) {
	if err := t.validatePutArchiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArchive",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutBounce(value interface{}) {
	if err := t.validatePutBounceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBounce",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutDeliverToMailbox(value interface{}) {
	if err := t.validatePutDeliverToMailboxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeliverToMailbox",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutDeliverToQBusiness(value interface{}) {
	if err := t.validatePutDeliverToQBusinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeliverToQBusiness",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutDrop(value interface{}) {
	if err := t.validatePutDropParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDrop",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutInvokeLambda(value interface{}) {
	if err := t.validatePutInvokeLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInvokeLambda",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutPublishToSns(value interface{}) {
	if err := t.validatePutPublishToSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPublishToSns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutRelay(value interface{}) {
	if err := t.validatePutRelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelay",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutReplaceRecipient(value interface{}) {
	if err := t.validatePutReplaceRecipientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReplaceRecipient",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutSend(value interface{}) {
	if err := t.validatePutSendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSend",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) PutWriteToS3(value interface{}) {
	if err := t.validatePutWriteToS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWriteToS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetAddHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetAddHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetArchive() {
	_jsii_.InvokeVoid(
		t,
		"resetArchive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetBounce() {
	_jsii_.InvokeVoid(
		t,
		"resetBounce",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetDeliverToMailbox() {
	_jsii_.InvokeVoid(
		t,
		"resetDeliverToMailbox",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetDeliverToQBusiness() {
	_jsii_.InvokeVoid(
		t,
		"resetDeliverToQBusiness",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetDrop() {
	_jsii_.InvokeVoid(
		t,
		"resetDrop",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetInvokeLambda() {
	_jsii_.InvokeVoid(
		t,
		"resetInvokeLambda",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetPublishToSns() {
	_jsii_.InvokeVoid(
		t,
		"resetPublishToSns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetRelay() {
	_jsii_.InvokeVoid(
		t,
		"resetRelay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetReplaceRecipient() {
	_jsii_.InvokeVoid(
		t,
		"resetReplaceRecipient",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetSend() {
	_jsii_.InvokeVoid(
		t,
		"resetSend",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ResetWriteToS3() {
	_jsii_.InvokeVoid(
		t,
		"resetWriteToS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRuleSet_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

