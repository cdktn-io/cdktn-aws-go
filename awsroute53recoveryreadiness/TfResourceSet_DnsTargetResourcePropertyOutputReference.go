package awsroute53recoveryreadiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53recoveryreadiness/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53recoveryreadiness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResourceSet_DnsTargetResourcePropertyOutputReference interface {
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
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HostedZoneArn() *string
	// Experimental.
	SetHostedZoneArn(val *string)
	// Experimental.
	HostedZoneArnInput() *string
	// Experimental.
	InternalValue() *TfResourceSet_DnsTargetResourceProperty
	// Experimental.
	SetInternalValue(val *TfResourceSet_DnsTargetResourceProperty)
	// Experimental.
	RecordSetId() *string
	// Experimental.
	SetRecordSetId(val *string)
	// Experimental.
	RecordSetIdInput() *string
	// Experimental.
	RecordType() *string
	// Experimental.
	SetRecordType(val *string)
	// Experimental.
	RecordTypeInput() *string
	// Experimental.
	TargetResource() TfResourceSet_TargetResourcePropertyOutputReference
	// Experimental.
	TargetResourceInput() *TfResourceSet_TargetResourceProperty
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
	PutTargetResource(value *TfResourceSet_TargetResourceProperty)
	// Experimental.
	ResetHostedZoneArn()
	// Experimental.
	ResetRecordSetId()
	// Experimental.
	ResetRecordType()
	// Experimental.
	ResetTargetResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfResourceSet_DnsTargetResourcePropertyOutputReference
type jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) HostedZoneArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) InternalValue() *TfResourceSet_DnsTargetResourceProperty {
	var returns *TfResourceSet_DnsTargetResourceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) RecordSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) RecordSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) RecordType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) RecordTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) TargetResource() TfResourceSet_TargetResourcePropertyOutputReference {
	var returns TfResourceSet_TargetResourcePropertyOutputReference
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) TargetResourceInput() *TfResourceSet_TargetResourceProperty {
	var returns *TfResourceSet_TargetResourceProperty
	_jsii_.Get(
		j,
		"targetResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfResourceSet_DnsTargetResourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfResourceSet_DnsTargetResourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfResourceSet_DnsTargetResourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53-recovery-readiness.TfResourceSet.DnsTargetResourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfResourceSet_DnsTargetResourcePropertyOutputReference_Override(t TfResourceSet_DnsTargetResourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53-recovery-readiness.TfResourceSet.DnsTargetResourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetHostedZoneArn(val *string) {
	if err := j.validateSetHostedZoneArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostedZoneArn",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetInternalValue(val *TfResourceSet_DnsTargetResourceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetRecordSetId(val *string) {
	if err := j.validateSetRecordSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSetId",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetRecordType(val *string) {
	if err := j.validateSetRecordTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordType",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) PutTargetResource(value *TfResourceSet_TargetResourceProperty) {
	if err := t.validatePutTargetResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ResetHostedZoneArn() {
	_jsii_.InvokeVoid(
		t,
		"resetHostedZoneArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ResetRecordSetId() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordSetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ResetRecordType() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ResetTargetResource() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfResourceSet_DnsTargetResourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

