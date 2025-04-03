# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs

__all__ = [
    'GetFunctionsNamespaceProperties',
    'AwaitableGetFunctionsNamespaceProperties',
    'get_functions_namespace',
    'get_functions_namespace_output',
]

@pulumi.output_type
class GetFunctionsNamespaceProperties:
    def __init__(__self__, namespace=None):
        if namespace and not isinstance(namespace, dict):
            raise TypeError("Expected argument 'namespace' to be a dict")
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def namespace(self) -> Optional['outputs.NamespaceInfo']:
        return pulumi.get(self, "namespace")


class AwaitableGetFunctionsNamespaceProperties(GetFunctionsNamespaceProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionsNamespaceProperties(
            namespace=self.namespace)


def get_functions_namespace(namespace_id: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionsNamespaceProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str namespace_id: The ID of the namespace to be managed.
    """
    __args__ = dict()
    __args__['namespaceId'] = namespace_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:functions/v2:getFunctionsNamespace', __args__, opts=opts, typ=GetFunctionsNamespaceProperties).value

    return AwaitableGetFunctionsNamespaceProperties(
        namespace=pulumi.get(__ret__, 'namespace'))
def get_functions_namespace_output(namespace_id: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFunctionsNamespaceProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str namespace_id: The ID of the namespace to be managed.
    """
    __args__ = dict()
    __args__['namespaceId'] = namespace_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:functions/v2:getFunctionsNamespace', __args__, opts=opts, typ=GetFunctionsNamespaceProperties)
    return __ret__.apply(lambda __response__: GetFunctionsNamespaceProperties(
        namespace=pulumi.get(__response__, 'namespace')))
