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
    'ListFunctionsNamespacesItems',
    'AwaitableListFunctionsNamespacesItems',
    'list_functions_namespaces',
    'list_functions_namespaces_output',
]

@pulumi.output_type
class ListFunctionsNamespacesItems:
    def __init__(__self__, namespaces=None):
        if namespaces and not isinstance(namespaces, list):
            raise TypeError("Expected argument 'namespaces' to be a list")
        pulumi.set(__self__, "namespaces", namespaces)

    @property
    @pulumi.getter
    def namespaces(self) -> Optional[Sequence['outputs.NamespaceInfo']]:
        return pulumi.get(self, "namespaces")


class AwaitableListFunctionsNamespacesItems(ListFunctionsNamespacesItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListFunctionsNamespacesItems(
            namespaces=self.namespaces)


def list_functions_namespaces(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListFunctionsNamespacesItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:functions/v2:listFunctionsNamespaces', __args__, opts=opts, typ=ListFunctionsNamespacesItems).value

    return AwaitableListFunctionsNamespacesItems(
        namespaces=pulumi.get(__ret__, 'namespaces'))
def list_functions_namespaces_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListFunctionsNamespacesItems]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:functions/v2:listFunctionsNamespaces', __args__, opts=opts, typ=ListFunctionsNamespacesItems)
    return __ret__.apply(lambda __response__: ListFunctionsNamespacesItems(
        namespaces=pulumi.get(__response__, 'namespaces')))
