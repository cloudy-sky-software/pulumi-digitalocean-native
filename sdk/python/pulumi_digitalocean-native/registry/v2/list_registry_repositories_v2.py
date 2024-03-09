# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'ListRegistryRepositoriesV2Result',
    'AwaitableListRegistryRepositoriesV2Result',
    'list_registry_repositories_v2',
    'list_registry_repositories_v2_output',
]

@pulumi.output_type
class ListRegistryRepositoriesV2Result:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListRegistryRepositoriesV2':
        return pulumi.get(self, "items")


class AwaitableListRegistryRepositoriesV2Result(ListRegistryRepositoriesV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListRegistryRepositoriesV2Result(
            items=self.items)


def list_registry_repositories_v2(registry_name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListRegistryRepositoriesV2Result:
    """
    Use this data source to access information about an existing resource.

    :param str registry_name: The name of a container registry.
    """
    __args__ = dict()
    __args__['registryName'] = registry_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:registry/v2:listRegistryRepositoriesV2', __args__, opts=opts, typ=ListRegistryRepositoriesV2Result).value

    return AwaitableListRegistryRepositoriesV2Result(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_registry_repositories_v2)
def list_registry_repositories_v2_output(registry_name: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListRegistryRepositoriesV2Result]:
    """
    Use this data source to access information about an existing resource.

    :param str registry_name: The name of a container registry.
    """
    ...
