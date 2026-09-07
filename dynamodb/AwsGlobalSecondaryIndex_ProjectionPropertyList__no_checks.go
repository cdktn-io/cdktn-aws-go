//go:build no_runtime_type_checking

package dynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsGlobalSecondaryIndex_ProjectionPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsGlobalSecondaryIndex_ProjectionPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsGlobalSecondaryIndex_ProjectionPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_ProjectionPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_ProjectionPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_ProjectionPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsGlobalSecondaryIndex_ProjectionPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsGlobalSecondaryIndex_ProjectionPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

