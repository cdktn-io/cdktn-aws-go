package awsroute53recoveryreadiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53recoveryreadiness/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53recoveryreadiness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference interface {
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
	InternalValue() *AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourceProperty
	// Experimental.
	SetInternalValue(val *AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourceProperty)
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
	TargetResource() AwsRoute53RecoveryreadinessResourceSet_TargetResourcePropertyOutputReference
	// Experimental.
	TargetResourceInput() *AwsRoute53RecoveryreadinessResourceSet_TargetResourceProperty
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
	PutTargetResource(value *AwsRoute53RecoveryreadinessResourceSet_TargetResourceProperty)
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

// The jsii proxy struct for AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference
type jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) HostedZoneArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) InternalValue() *AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourceProperty {
	var returns *AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) RecordSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) RecordSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) RecordType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) RecordTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) TargetResource() AwsRoute53RecoveryreadinessResourceSet_TargetResourcePropertyOutputReference {
	var returns AwsRoute53RecoveryreadinessResourceSet_TargetResourcePropertyOutputReference
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) TargetResourceInput() *AwsRoute53RecoveryreadinessResourceSet_TargetResourceProperty {
	var returns *AwsRoute53RecoveryreadinessResourceSet_TargetResourceProperty
	_jsii_.Get(
		j,
		"targetResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53-recovery-readiness.AwsRoute53RecoveryreadinessResourceSet.DnsTargetResourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference_Override(a AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53-recovery-readiness.AwsRoute53RecoveryreadinessResourceSet.DnsTargetResourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetHostedZoneArn(val *string) {
	if err := j.validateSetHostedZoneArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostedZoneArn",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetInternalValue(val *AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetRecordSetId(val *string) {
	if err := j.validateSetRecordSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSetId",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetRecordType(val *string) {
	if err := j.validateSetRecordTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordType",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) PutTargetResource(value *AwsRoute53RecoveryreadinessResourceSet_TargetResourceProperty) {
	if err := a.validatePutTargetResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ResetHostedZoneArn() {
	_jsii_.InvokeVoid(
		a,
		"resetHostedZoneArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ResetRecordSetId() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordSetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ResetRecordType() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ResetTargetResource() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRoute53RecoveryreadinessResourceSet_DnsTargetResourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

