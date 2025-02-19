// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlUser struct {
	pulumi.CustomResourceState

	// The deletion policy for the user. Setting ABANDON allows the resource to be abandoned rather than deleted. This is
	// useful for Postgres, where users cannot be deleted from the API if they have been granted SQL roles. Possible values
	// are: "ABANDON".
	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL
	// instances. Can be an IP address. Changing this forces a new resource to be created.
	Host pulumi.StringOutput `pulumi:"host"`
	// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The name of the user. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to either
	// CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.
	Password       pulumi.StringPtrOutput         `pulumi:"password"`
	PasswordPolicy SqlUserPasswordPolicyPtrOutput `pulumi:"passwordPolicy"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project              pulumi.StringOutput                   `pulumi:"project"`
	SqlServerUserDetails SqlUserSqlServerUserDetailArrayOutput `pulumi:"sqlServerUserDetails"`
	SqlUserId            pulumi.StringOutput                   `pulumi:"sqlUserId"`
	Timeouts             SqlUserTimeoutsPtrOutput              `pulumi:"timeouts"`
	// The user type. It determines the method to authenticate the user during login. The default is the database's built-in
	// user type.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSqlUser registers a new resource with the given unique name, arguments, and options.
func NewSqlUser(ctx *pulumi.Context,
	name string, args *SqlUserArgs, opts ...pulumi.ResourceOption) (*SqlUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SqlUser
	err = ctx.RegisterPackageResource("google:index/sqlUser:SqlUser", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlUser gets an existing SqlUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlUserState, opts ...pulumi.ResourceOption) (*SqlUser, error) {
	var resource SqlUser
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/sqlUser:SqlUser", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlUser resources.
type sqlUserState struct {
	// The deletion policy for the user. Setting ABANDON allows the resource to be abandoned rather than deleted. This is
	// useful for Postgres, where users cannot be deleted from the API if they have been granted SQL roles. Possible values
	// are: "ABANDON".
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL
	// instances. Can be an IP address. Changing this forces a new resource to be created.
	Host *string `pulumi:"host"`
	// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
	Instance *string `pulumi:"instance"`
	// The name of the user. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to either
	// CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.
	Password       *string                `pulumi:"password"`
	PasswordPolicy *SqlUserPasswordPolicy `pulumi:"passwordPolicy"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project              *string                      `pulumi:"project"`
	SqlServerUserDetails []SqlUserSqlServerUserDetail `pulumi:"sqlServerUserDetails"`
	SqlUserId            *string                      `pulumi:"sqlUserId"`
	Timeouts             *SqlUserTimeouts             `pulumi:"timeouts"`
	// The user type. It determines the method to authenticate the user during login. The default is the database's built-in
	// user type.
	Type *string `pulumi:"type"`
}

type SqlUserState struct {
	// The deletion policy for the user. Setting ABANDON allows the resource to be abandoned rather than deleted. This is
	// useful for Postgres, where users cannot be deleted from the API if they have been granted SQL roles. Possible values
	// are: "ABANDON".
	DeletionPolicy pulumi.StringPtrInput
	// The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL
	// instances. Can be an IP address. Changing this forces a new resource to be created.
	Host pulumi.StringPtrInput
	// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
	Instance pulumi.StringPtrInput
	// The name of the user. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to either
	// CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.
	Password       pulumi.StringPtrInput
	PasswordPolicy SqlUserPasswordPolicyPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project              pulumi.StringPtrInput
	SqlServerUserDetails SqlUserSqlServerUserDetailArrayInput
	SqlUserId            pulumi.StringPtrInput
	Timeouts             SqlUserTimeoutsPtrInput
	// The user type. It determines the method to authenticate the user during login. The default is the database's built-in
	// user type.
	Type pulumi.StringPtrInput
}

func (SqlUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlUserState)(nil)).Elem()
}

type sqlUserArgs struct {
	// The deletion policy for the user. Setting ABANDON allows the resource to be abandoned rather than deleted. This is
	// useful for Postgres, where users cannot be deleted from the API if they have been granted SQL roles. Possible values
	// are: "ABANDON".
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL
	// instances. Can be an IP address. Changing this forces a new resource to be created.
	Host *string `pulumi:"host"`
	// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
	Instance string `pulumi:"instance"`
	// The name of the user. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to either
	// CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.
	Password       *string                `pulumi:"password"`
	PasswordPolicy *SqlUserPasswordPolicy `pulumi:"passwordPolicy"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project   *string          `pulumi:"project"`
	SqlUserId *string          `pulumi:"sqlUserId"`
	Timeouts  *SqlUserTimeouts `pulumi:"timeouts"`
	// The user type. It determines the method to authenticate the user during login. The default is the database's built-in
	// user type.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SqlUser resource.
type SqlUserArgs struct {
	// The deletion policy for the user. Setting ABANDON allows the resource to be abandoned rather than deleted. This is
	// useful for Postgres, where users cannot be deleted from the API if they have been granted SQL roles. Possible values
	// are: "ABANDON".
	DeletionPolicy pulumi.StringPtrInput
	// The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL
	// instances. Can be an IP address. Changing this forces a new resource to be created.
	Host pulumi.StringPtrInput
	// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
	Instance pulumi.StringInput
	// The name of the user. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to either
	// CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.
	Password       pulumi.StringPtrInput
	PasswordPolicy SqlUserPasswordPolicyPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project   pulumi.StringPtrInput
	SqlUserId pulumi.StringPtrInput
	Timeouts  SqlUserTimeoutsPtrInput
	// The user type. It determines the method to authenticate the user during login. The default is the database's built-in
	// user type.
	Type pulumi.StringPtrInput
}

func (SqlUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlUserArgs)(nil)).Elem()
}

type SqlUserInput interface {
	pulumi.Input

	ToSqlUserOutput() SqlUserOutput
	ToSqlUserOutputWithContext(ctx context.Context) SqlUserOutput
}

func (*SqlUser) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUser)(nil)).Elem()
}

func (i *SqlUser) ToSqlUserOutput() SqlUserOutput {
	return i.ToSqlUserOutputWithContext(context.Background())
}

func (i *SqlUser) ToSqlUserOutputWithContext(ctx context.Context) SqlUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserOutput)
}

type SqlUserOutput struct{ *pulumi.OutputState }

func (SqlUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUser)(nil)).Elem()
}

func (o SqlUserOutput) ToSqlUserOutput() SqlUserOutput {
	return o
}

func (o SqlUserOutput) ToSqlUserOutputWithContext(ctx context.Context) SqlUserOutput {
	return o
}

// The deletion policy for the user. Setting ABANDON allows the resource to be abandoned rather than deleted. This is
// useful for Postgres, where users cannot be deleted from the API if they have been granted SQL roles. Possible values
// are: "ABANDON".
func (o SqlUserOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

// The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL
// instances. Can be an IP address. Changing this forces a new resource to be created.
func (o SqlUserOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
func (o SqlUserOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The name of the user. Changing this forces a new resource to be created.
func (o SqlUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to either
// CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT.
func (o SqlUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SqlUserOutput) PasswordPolicy() SqlUserPasswordPolicyPtrOutput {
	return o.ApplyT(func(v *SqlUser) SqlUserPasswordPolicyPtrOutput { return v.PasswordPolicy }).(SqlUserPasswordPolicyPtrOutput)
}

// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
func (o SqlUserOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SqlUserOutput) SqlServerUserDetails() SqlUserSqlServerUserDetailArrayOutput {
	return o.ApplyT(func(v *SqlUser) SqlUserSqlServerUserDetailArrayOutput { return v.SqlServerUserDetails }).(SqlUserSqlServerUserDetailArrayOutput)
}

func (o SqlUserOutput) SqlUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.SqlUserId }).(pulumi.StringOutput)
}

func (o SqlUserOutput) Timeouts() SqlUserTimeoutsPtrOutput {
	return o.ApplyT(func(v *SqlUser) SqlUserTimeoutsPtrOutput { return v.Timeouts }).(SqlUserTimeoutsPtrOutput)
}

// The user type. It determines the method to authenticate the user during login. The default is the database's built-in
// user type.
func (o SqlUserOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlUserInput)(nil)).Elem(), &SqlUser{})
	pulumi.RegisterOutputType(SqlUserOutput{})
}
