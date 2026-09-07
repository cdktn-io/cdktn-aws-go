package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference interface {
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
	DialogAction() AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepDialogActionPropertyList
	// Experimental.
	DialogActionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Intent() AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepIntentPropertyList
	// Experimental.
	IntentInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SessionAttributes() *map[string]*string
	// Experimental.
	SetSessionAttributes(val *map[string]*string)
	// Experimental.
	SessionAttributesInput() *map[string]*string
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
	PutDialogAction(value interface{})
	// Experimental.
	PutIntent(value interface{})
	// Experimental.
	ResetDialogAction()
	// Experimental.
	ResetIntent()
	// Experimental.
	ResetSessionAttributes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference
type jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) DialogAction() AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepDialogActionPropertyList {
	var returns AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepDialogActionPropertyList
	_jsii_.Get(
		j,
		"dialogAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) DialogActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dialogActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) Intent() AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepIntentPropertyList {
	var returns AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepIntentPropertyList
	_jsii_.Get(
		j,
		"intent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) IntentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) SessionAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sessionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) SessionAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sessionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference_Override(a AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference)SetSessionAttributes(val *map[string]*string) {
	if err := j.validateSetSessionAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) PutDialogAction(value interface{}) {
	if err := a.validatePutDialogActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDialogAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) PutIntent(value interface{}) {
	if err := a.validatePutIntentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) ResetDialogAction() {
	_jsii_.InvokeVoid(
		a,
		"resetDialogAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) ResetIntent() {
	_jsii_.InvokeVoid(
		a,
		"resetIntent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) ResetSessionAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchNextStepPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

