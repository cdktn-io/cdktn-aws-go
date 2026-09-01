package awssesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMailmanagerRuleSet_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddHeader() AwsMailmanagerRuleSet_AddHeaderPropertyList
	// Experimental.
	AddHeaderInput() interface{}
	// Experimental.
	Archive() AwsMailmanagerRuleSet_ArchivePropertyList
	// Experimental.
	ArchiveInput() interface{}
	// Experimental.
	Bounce() AwsMailmanagerRuleSet_BouncePropertyList
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
	DeliverToMailbox() AwsMailmanagerRuleSet_DeliverToMailboxPropertyList
	// Experimental.
	DeliverToMailboxInput() interface{}
	// Experimental.
	DeliverToQBusiness() AwsMailmanagerRuleSet_DeliverToQBusinessPropertyList
	// Experimental.
	DeliverToQBusinessInput() interface{}
	// Experimental.
	Drop() AwsMailmanagerRuleSet_DropPropertyList
	// Experimental.
	DropInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	InvokeLambda() AwsMailmanagerRuleSet_InvokeLambdaPropertyList
	// Experimental.
	InvokeLambdaInput() interface{}
	// Experimental.
	PublishToSns() AwsMailmanagerRuleSet_PublishToSnsPropertyList
	// Experimental.
	PublishToSnsInput() interface{}
	// Experimental.
	Relay() AwsMailmanagerRuleSet_RelayPropertyList
	// Experimental.
	RelayInput() interface{}
	// Experimental.
	ReplaceRecipient() AwsMailmanagerRuleSet_ReplaceRecipientPropertyList
	// Experimental.
	ReplaceRecipientInput() interface{}
	// Experimental.
	Send() AwsMailmanagerRuleSet_SendPropertyList
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
	WriteToS3() AwsMailmanagerRuleSet_WriteToS3PropertyList
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

// The jsii proxy struct for AwsMailmanagerRuleSet_ActionPropertyOutputReference
type jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) AddHeader() AwsMailmanagerRuleSet_AddHeaderPropertyList {
	var returns AwsMailmanagerRuleSet_AddHeaderPropertyList
	_jsii_.Get(
		j,
		"addHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) AddHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) Archive() AwsMailmanagerRuleSet_ArchivePropertyList {
	var returns AwsMailmanagerRuleSet_ArchivePropertyList
	_jsii_.Get(
		j,
		"archive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ArchiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) Bounce() AwsMailmanagerRuleSet_BouncePropertyList {
	var returns AwsMailmanagerRuleSet_BouncePropertyList
	_jsii_.Get(
		j,
		"bounce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) BounceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bounceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) DeliverToMailbox() AwsMailmanagerRuleSet_DeliverToMailboxPropertyList {
	var returns AwsMailmanagerRuleSet_DeliverToMailboxPropertyList
	_jsii_.Get(
		j,
		"deliverToMailbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) DeliverToMailboxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToMailboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) DeliverToQBusiness() AwsMailmanagerRuleSet_DeliverToQBusinessPropertyList {
	var returns AwsMailmanagerRuleSet_DeliverToQBusinessPropertyList
	_jsii_.Get(
		j,
		"deliverToQBusiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) DeliverToQBusinessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToQBusinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) Drop() AwsMailmanagerRuleSet_DropPropertyList {
	var returns AwsMailmanagerRuleSet_DropPropertyList
	_jsii_.Get(
		j,
		"drop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) DropInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) InvokeLambda() AwsMailmanagerRuleSet_InvokeLambdaPropertyList {
	var returns AwsMailmanagerRuleSet_InvokeLambdaPropertyList
	_jsii_.Get(
		j,
		"invokeLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) InvokeLambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invokeLambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PublishToSns() AwsMailmanagerRuleSet_PublishToSnsPropertyList {
	var returns AwsMailmanagerRuleSet_PublishToSnsPropertyList
	_jsii_.Get(
		j,
		"publishToSns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PublishToSnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishToSnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) Relay() AwsMailmanagerRuleSet_RelayPropertyList {
	var returns AwsMailmanagerRuleSet_RelayPropertyList
	_jsii_.Get(
		j,
		"relay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) RelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ReplaceRecipient() AwsMailmanagerRuleSet_ReplaceRecipientPropertyList {
	var returns AwsMailmanagerRuleSet_ReplaceRecipientPropertyList
	_jsii_.Get(
		j,
		"replaceRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ReplaceRecipientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) Send() AwsMailmanagerRuleSet_SendPropertyList {
	var returns AwsMailmanagerRuleSet_SendPropertyList
	_jsii_.Get(
		j,
		"send",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) SendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) WriteToS3() AwsMailmanagerRuleSet_WriteToS3PropertyList {
	var returns AwsMailmanagerRuleSet_WriteToS3PropertyList
	_jsii_.Get(
		j,
		"writeToS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) WriteToS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeToS3Input",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMailmanagerRuleSet_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMailmanagerRuleSet_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMailmanagerRuleSet_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsMailmanagerRuleSet.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMailmanagerRuleSet_ActionPropertyOutputReference_Override(a AwsMailmanagerRuleSet_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsMailmanagerRuleSet.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutAddHeader(value interface{}) {
	if err := a.validatePutAddHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAddHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutArchive(value interface{}) {
	if err := a.validatePutArchiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArchive",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutBounce(value interface{}) {
	if err := a.validatePutBounceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBounce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutDeliverToMailbox(value interface{}) {
	if err := a.validatePutDeliverToMailboxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeliverToMailbox",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutDeliverToQBusiness(value interface{}) {
	if err := a.validatePutDeliverToQBusinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeliverToQBusiness",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutDrop(value interface{}) {
	if err := a.validatePutDropParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDrop",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutInvokeLambda(value interface{}) {
	if err := a.validatePutInvokeLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInvokeLambda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutPublishToSns(value interface{}) {
	if err := a.validatePutPublishToSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPublishToSns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutRelay(value interface{}) {
	if err := a.validatePutRelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelay",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutReplaceRecipient(value interface{}) {
	if err := a.validatePutReplaceRecipientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReplaceRecipient",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutSend(value interface{}) {
	if err := a.validatePutSendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSend",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) PutWriteToS3(value interface{}) {
	if err := a.validatePutWriteToS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWriteToS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetAddHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetAddHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetArchive() {
	_jsii_.InvokeVoid(
		a,
		"resetArchive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetBounce() {
	_jsii_.InvokeVoid(
		a,
		"resetBounce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetDeliverToMailbox() {
	_jsii_.InvokeVoid(
		a,
		"resetDeliverToMailbox",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetDeliverToQBusiness() {
	_jsii_.InvokeVoid(
		a,
		"resetDeliverToQBusiness",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetDrop() {
	_jsii_.InvokeVoid(
		a,
		"resetDrop",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetInvokeLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetInvokeLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetPublishToSns() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishToSns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetRelay() {
	_jsii_.InvokeVoid(
		a,
		"resetRelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetReplaceRecipient() {
	_jsii_.InvokeVoid(
		a,
		"resetReplaceRecipient",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetSend() {
	_jsii_.InvokeVoid(
		a,
		"resetSend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ResetWriteToS3() {
	_jsii_.InvokeVoid(
		a,
		"resetWriteToS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

