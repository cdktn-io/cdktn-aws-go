package awsvpnsitetosite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpnsitetosite/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpnsitetosite/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BgpLogEnabled() interface{}
	// Experimental.
	SetBgpLogEnabled(val interface{})
	// Experimental.
	BgpLogEnabledInput() interface{}
	// Experimental.
	BgpLogGroupArn() *string
	// Experimental.
	SetBgpLogGroupArn(val *string)
	// Experimental.
	BgpLogGroupArnInput() *string
	// Experimental.
	BgpLogOutputFormat() *string
	// Experimental.
	SetBgpLogOutputFormat(val *string)
	// Experimental.
	BgpLogOutputFormatInput() *string
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsProperty
	// Experimental.
	SetInternalValue(val *TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsProperty)
	// Experimental.
	LogEnabled() interface{}
	// Experimental.
	SetLogEnabled(val interface{})
	// Experimental.
	LogEnabledInput() interface{}
	// Experimental.
	LogGroupArn() *string
	// Experimental.
	SetLogGroupArn(val *string)
	// Experimental.
	LogGroupArnInput() *string
	// Experimental.
	LogOutputFormat() *string
	// Experimental.
	SetLogOutputFormat(val *string)
	// Experimental.
	LogOutputFormatInput() *string
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
	ResetBgpLogEnabled()
	// Experimental.
	ResetBgpLogGroupArn()
	// Experimental.
	ResetBgpLogOutputFormat()
	// Experimental.
	ResetLogEnabled()
	// Experimental.
	ResetLogGroupArn()
	// Experimental.
	ResetLogOutputFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference
type jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) BgpLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) BgpLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) BgpLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) BgpLogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpLogGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) BgpLogOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpLogOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) BgpLogOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpLogOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) InternalValue() *TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsProperty {
	var returns *TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) LogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) LogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) LogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) LogOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) LogOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpn-site-to-site.TfConnection.Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference_Override(t TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpn-site-to-site.TfConnection.Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetBgpLogEnabled(val interface{}) {
	if err := j.validateSetBgpLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpLogEnabled",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetBgpLogGroupArn(val *string) {
	if err := j.validateSetBgpLogGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetBgpLogOutputFormat(val *string) {
	if err := j.validateSetBgpLogOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpLogOutputFormat",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetInternalValue(val *TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetLogEnabled(val interface{}) {
	if err := j.validateSetLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEnabled",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetLogGroupArn(val *string) {
	if err := j.validateSetLogGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetLogOutputFormat(val *string) {
	if err := j.validateSetLogOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logOutputFormat",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ResetBgpLogEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetBgpLogEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ResetBgpLogGroupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetBgpLogGroupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ResetBgpLogOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetBgpLogOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ResetLogEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetLogEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ResetLogGroupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetLogGroupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ResetLogOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetLogOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnection_Tunnel2LogOptionsCloudwatchLogOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

