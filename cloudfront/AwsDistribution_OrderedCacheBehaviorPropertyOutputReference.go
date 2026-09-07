package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDistribution_OrderedCacheBehaviorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowedMethods() *[]*string
	// Experimental.
	SetAllowedMethods(val *[]*string)
	// Experimental.
	AllowedMethodsInput() *[]*string
	// Experimental.
	CachedMethods() *[]*string
	// Experimental.
	SetCachedMethods(val *[]*string)
	// Experimental.
	CachedMethodsInput() *[]*string
	// Experimental.
	CachePolicyId() *string
	// Experimental.
	SetCachePolicyId(val *string)
	// Experimental.
	CachePolicyIdInput() *string
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
	Compress() interface{}
	// Experimental.
	SetCompress(val interface{})
	// Experimental.
	CompressInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DefaultTtl() *float64
	// Experimental.
	SetDefaultTtl(val *float64)
	// Experimental.
	DefaultTtlInput() *float64
	// Experimental.
	FieldLevelEncryptionId() *string
	// Experimental.
	SetFieldLevelEncryptionId(val *string)
	// Experimental.
	FieldLevelEncryptionIdInput() *string
	// Experimental.
	ForwardedValues() AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference
	// Experimental.
	ForwardedValuesInput() *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	FunctionAssociation() AwsDistribution_OrderedCacheBehaviorFunctionAssociationPropertyList
	// Experimental.
	FunctionAssociationInput() interface{}
	// Experimental.
	GrpcConfig() AwsDistribution_OrderedCacheBehaviorGrpcConfigPropertyOutputReference
	// Experimental.
	GrpcConfigInput() *AwsDistribution_OrderedCacheBehaviorGrpcConfigProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaFunctionAssociation() AwsDistribution_OrderedCacheBehaviorLambdaFunctionAssociationPropertyList
	// Experimental.
	LambdaFunctionAssociationInput() interface{}
	// Experimental.
	MaxTtl() *float64
	// Experimental.
	SetMaxTtl(val *float64)
	// Experimental.
	MaxTtlInput() *float64
	// Experimental.
	MinTtl() *float64
	// Experimental.
	SetMinTtl(val *float64)
	// Experimental.
	MinTtlInput() *float64
	// Experimental.
	OriginRequestPolicyId() *string
	// Experimental.
	SetOriginRequestPolicyId(val *string)
	// Experimental.
	OriginRequestPolicyIdInput() *string
	// Experimental.
	PathPattern() *string
	// Experimental.
	SetPathPattern(val *string)
	// Experimental.
	PathPatternInput() *string
	// Experimental.
	RealtimeLogConfigArn() *string
	// Experimental.
	SetRealtimeLogConfigArn(val *string)
	// Experimental.
	RealtimeLogConfigArnInput() *string
	// Experimental.
	ResponseHeadersPolicyId() *string
	// Experimental.
	SetResponseHeadersPolicyId(val *string)
	// Experimental.
	ResponseHeadersPolicyIdInput() *string
	// Experimental.
	SmoothStreaming() interface{}
	// Experimental.
	SetSmoothStreaming(val interface{})
	// Experimental.
	SmoothStreamingInput() interface{}
	// Experimental.
	TargetOriginId() *string
	// Experimental.
	SetTargetOriginId(val *string)
	// Experimental.
	TargetOriginIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrustedKeyGroups() *[]*string
	// Experimental.
	SetTrustedKeyGroups(val *[]*string)
	// Experimental.
	TrustedKeyGroupsInput() *[]*string
	// Experimental.
	TrustedSigners() *[]*string
	// Experimental.
	SetTrustedSigners(val *[]*string)
	// Experimental.
	TrustedSignersInput() *[]*string
	// Experimental.
	ViewerProtocolPolicy() *string
	// Experimental.
	SetViewerProtocolPolicy(val *string)
	// Experimental.
	ViewerProtocolPolicyInput() *string
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
	PutForwardedValues(value *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty)
	// Experimental.
	PutFunctionAssociation(value interface{})
	// Experimental.
	PutGrpcConfig(value *AwsDistribution_OrderedCacheBehaviorGrpcConfigProperty)
	// Experimental.
	PutLambdaFunctionAssociation(value interface{})
	// Experimental.
	ResetCachePolicyId()
	// Experimental.
	ResetCompress()
	// Experimental.
	ResetDefaultTtl()
	// Experimental.
	ResetFieldLevelEncryptionId()
	// Experimental.
	ResetForwardedValues()
	// Experimental.
	ResetFunctionAssociation()
	// Experimental.
	ResetGrpcConfig()
	// Experimental.
	ResetLambdaFunctionAssociation()
	// Experimental.
	ResetMaxTtl()
	// Experimental.
	ResetMinTtl()
	// Experimental.
	ResetOriginRequestPolicyId()
	// Experimental.
	ResetRealtimeLogConfigArn()
	// Experimental.
	ResetResponseHeadersPolicyId()
	// Experimental.
	ResetSmoothStreaming()
	// Experimental.
	ResetTrustedKeyGroups()
	// Experimental.
	ResetTrustedSigners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDistribution_OrderedCacheBehaviorPropertyOutputReference
type jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) CachedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) CachedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) CachePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) Compress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) CompressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) FieldLevelEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ForwardedValues() AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference {
	var returns AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ForwardedValuesInput() *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty {
	var returns *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty
	_jsii_.Get(
		j,
		"forwardedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) FunctionAssociation() AwsDistribution_OrderedCacheBehaviorFunctionAssociationPropertyList {
	var returns AwsDistribution_OrderedCacheBehaviorFunctionAssociationPropertyList
	_jsii_.Get(
		j,
		"functionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) FunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GrpcConfig() AwsDistribution_OrderedCacheBehaviorGrpcConfigPropertyOutputReference {
	var returns AwsDistribution_OrderedCacheBehaviorGrpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"grpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GrpcConfigInput() *AwsDistribution_OrderedCacheBehaviorGrpcConfigProperty {
	var returns *AwsDistribution_OrderedCacheBehaviorGrpcConfigProperty
	_jsii_.Get(
		j,
		"grpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) LambdaFunctionAssociation() AwsDistribution_OrderedCacheBehaviorLambdaFunctionAssociationPropertyList {
	var returns AwsDistribution_OrderedCacheBehaviorLambdaFunctionAssociationPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) LambdaFunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) MinTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) OriginRequestPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) PathPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) PathPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) RealtimeLogConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResponseHeadersPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TargetOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TrustedKeyGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TrustedKeyGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TrustedSigners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) TrustedSignersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSignersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ViewerProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDistribution_OrderedCacheBehaviorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDistribution_OrderedCacheBehaviorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDistribution_OrderedCacheBehaviorPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.OrderedCacheBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDistribution_OrderedCacheBehaviorPropertyOutputReference_Override(a AwsDistribution_OrderedCacheBehaviorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.OrderedCacheBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetAllowedMethods(val *[]*string) {
	if err := j.validateSetAllowedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetCachedMethods(val *[]*string) {
	if err := j.validateSetCachedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachedMethods",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetCachePolicyId(val *string) {
	if err := j.validateSetCachePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachePolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetCompress(val interface{}) {
	if err := j.validateSetCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compress",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetFieldLevelEncryptionId(val *string) {
	if err := j.validateSetFieldLevelEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLevelEncryptionId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetMinTtl(val *float64) {
	if err := j.validateSetMinTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetOriginRequestPolicyId(val *string) {
	if err := j.validateSetOriginRequestPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originRequestPolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetPathPattern(val *string) {
	if err := j.validateSetPathPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathPattern",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetRealtimeLogConfigArn(val *string) {
	if err := j.validateSetRealtimeLogConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeLogConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetResponseHeadersPolicyId(val *string) {
	if err := j.validateSetResponseHeadersPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseHeadersPolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetSmoothStreaming(val interface{}) {
	if err := j.validateSetSmoothStreamingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetTargetOriginId(val *string) {
	if err := j.validateSetTargetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOriginId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetTrustedKeyGroups(val *[]*string) {
	if err := j.validateSetTrustedKeyGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedKeyGroups",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetTrustedSigners(val *[]*string) {
	if err := j.validateSetTrustedSignersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedSigners",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference)SetViewerProtocolPolicy(val *string) {
	if err := j.validateSetViewerProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewerProtocolPolicy",
		val,
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) PutForwardedValues(value *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty) {
	if err := a.validatePutForwardedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForwardedValues",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) PutFunctionAssociation(value interface{}) {
	if err := a.validatePutFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFunctionAssociation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) PutGrpcConfig(value *AwsDistribution_OrderedCacheBehaviorGrpcConfigProperty) {
	if err := a.validatePutGrpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) PutLambdaFunctionAssociation(value interface{}) {
	if err := a.validatePutLambdaFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionAssociation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetCachePolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetCachePolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetCompress() {
	_jsii_.InvokeVoid(
		a,
		"resetCompress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetFieldLevelEncryptionId() {
	_jsii_.InvokeVoid(
		a,
		"resetFieldLevelEncryptionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetForwardedValues() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardedValues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetFunctionAssociation() {
	_jsii_.InvokeVoid(
		a,
		"resetFunctionAssociation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetGrpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetLambdaFunctionAssociation() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionAssociation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetMinTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMinTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetOriginRequestPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginRequestPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetRealtimeLogConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRealtimeLogConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetResponseHeadersPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseHeadersPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		a,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedKeyGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ResetTrustedSigners() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedSigners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

