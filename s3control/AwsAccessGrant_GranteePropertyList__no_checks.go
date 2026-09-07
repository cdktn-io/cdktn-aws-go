//go:build no_runtime_type_checking

package s3control

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsAccessGrant_GranteePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsAccessGrant_GranteePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsAccessGrant_GranteePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsAccessGrant_GranteePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsAccessGrant_GranteePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsAccessGrant_GranteePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsAccessGrant_GranteePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsAccessGrant_GranteePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

