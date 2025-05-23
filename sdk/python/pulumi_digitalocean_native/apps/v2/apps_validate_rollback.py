# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = ['AppsValidateRollbackArgs', 'AppsValidateRollback']

@pulumi.input_type
class AppsValidateRollbackArgs:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[builtins.str]] = None,
                 deployment_id: Optional[pulumi.Input[builtins.str]] = None,
                 skip_pin: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a AppsValidateRollback resource.
        :param pulumi.Input[builtins.str] app_id: The app ID
        :param pulumi.Input[builtins.str] deployment_id: The ID of the deployment to rollback to.
        :param pulumi.Input[builtins.bool] skip_pin: Whether to skip pinning the rollback deployment. If false, the rollback deployment will be pinned and any new deployments including Auto Deploy on Push hooks will be disabled until the rollback is either manually committed or reverted via the CommitAppRollback or RevertAppRollback endpoints respectively. If true, the rollback will be immediately committed and the app will remain unpinned.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if deployment_id is not None:
            pulumi.set(__self__, "deployment_id", deployment_id)
        if skip_pin is not None:
            pulumi.set(__self__, "skip_pin", skip_pin)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The app ID
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the deployment to rollback to.
        """
        return pulumi.get(self, "deployment_id")

    @deployment_id.setter
    def deployment_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "deployment_id", value)

    @property
    @pulumi.getter(name="skipPin")
    def skip_pin(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether to skip pinning the rollback deployment. If false, the rollback deployment will be pinned and any new deployments including Auto Deploy on Push hooks will be disabled until the rollback is either manually committed or reverted via the CommitAppRollback or RevertAppRollback endpoints respectively. If true, the rollback will be immediately committed and the app will remain unpinned.
        """
        return pulumi.get(self, "skip_pin")

    @skip_pin.setter
    def skip_pin(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "skip_pin", value)


class AppsValidateRollback(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[builtins.str]] = None,
                 deployment_id: Optional[pulumi.Input[builtins.str]] = None,
                 skip_pin: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        Create a AppsValidateRollback resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] app_id: The app ID
        :param pulumi.Input[builtins.str] deployment_id: The ID of the deployment to rollback to.
        :param pulumi.Input[builtins.bool] skip_pin: Whether to skip pinning the rollback deployment. If false, the rollback deployment will be pinned and any new deployments including Auto Deploy on Push hooks will be disabled until the rollback is either manually committed or reverted via the CommitAppRollback or RevertAppRollback endpoints respectively. If true, the rollback will be immediately committed and the app will remain unpinned.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AppsValidateRollbackArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AppsValidateRollback resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AppsValidateRollbackArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppsValidateRollbackArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[builtins.str]] = None,
                 deployment_id: Optional[pulumi.Input[builtins.str]] = None,
                 skip_pin: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppsValidateRollbackArgs.__new__(AppsValidateRollbackArgs)

            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["deployment_id"] = deployment_id
            __props__.__dict__["skip_pin"] = skip_pin
            __props__.__dict__["error"] = None
            __props__.__dict__["valid"] = None
            __props__.__dict__["warnings"] = None
        super(AppsValidateRollback, __self__).__init__(
            'digitalocean-native:apps/v2:AppsValidateRollback',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AppsValidateRollback':
        """
        Get an existing AppsValidateRollback resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AppsValidateRollbackArgs.__new__(AppsValidateRollbackArgs)

        __props__.__dict__["deployment_id"] = None
        __props__.__dict__["error"] = None
        __props__.__dict__["skip_pin"] = None
        __props__.__dict__["valid"] = None
        __props__.__dict__["warnings"] = None
        return AppsValidateRollback(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the deployment to rollback to.
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def error(self) -> pulumi.Output[Optional['outputs.Error']]:
        return pulumi.get(self, "error")

    @property
    @pulumi.getter(name="skipPin")
    def skip_pin(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Whether to skip pinning the rollback deployment. If false, the rollback deployment will be pinned and any new deployments including Auto Deploy on Push hooks will be disabled until the rollback is either manually committed or reverted via the CommitAppRollback or RevertAppRollback endpoints respectively. If true, the rollback will be immediately committed and the app will remain unpinned.
        """
        return pulumi.get(self, "skip_pin")

    @property
    @pulumi.getter
    def valid(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Indicates whether the app can be rolled back to the specified deployment.
        """
        return pulumi.get(self, "valid")

    @property
    @pulumi.getter
    def warnings(self) -> pulumi.Output[Optional[Sequence['outputs.AppRollbackValidationCondition']]]:
        """
        Contains a list of warnings that may cause the rollback to run under unideal circumstances.
        """
        return pulumi.get(self, "warnings")

