// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAccount(ctx *pulumi.Context, args *GetAccountArgs, opts ...pulumi.InvokeOption) (*GetAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccountResult
	err := ctx.Invoke("digitalocean-native:account/v2:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetAccountArgs struct {
}

type GetAccountResult struct {
	Account *Account `pulumi:"account"`
}

// Defaults sets the appropriate defaults for GetAccountResult
func (val *GetAccountResult) Defaults() *GetAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Account = tmp.Account.Defaults()

	return &tmp
}

func GetAccountOutput(ctx *pulumi.Context, args GetAccountOutputArgs, opts ...pulumi.InvokeOption) GetAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountResultOutput, error) {
			args := v.(GetAccountArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetAccountResult
			secret, err := ctx.InvokePackageRaw("digitalocean-native:account/v2:getAccount", args, &rv, "", opts...)
			if err != nil {
				return GetAccountResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetAccountResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetAccountResultOutput), nil
			}
			return output, nil
		}).(GetAccountResultOutput)
}

type GetAccountOutputArgs struct {
}

func (GetAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountArgs)(nil)).Elem()
}

type GetAccountResultOutput struct{ *pulumi.OutputState }

func (GetAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountResult)(nil)).Elem()
}

func (o GetAccountResultOutput) ToGetAccountResultOutput() GetAccountResultOutput {
	return o
}

func (o GetAccountResultOutput) ToGetAccountResultOutputWithContext(ctx context.Context) GetAccountResultOutput {
	return o
}

func (o GetAccountResultOutput) Account() AccountPtrOutput {
	return o.ApplyT(func(v GetAccountResult) *Account { return v.Account }).(AccountPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountResultOutput{})
}
