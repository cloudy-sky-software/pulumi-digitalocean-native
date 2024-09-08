# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'ListRegistryGarbageCollectionsProperties',
    'AwaitableListRegistryGarbageCollectionsProperties',
    'list_registry_garbage_collections',
    'list_registry_garbage_collections_output',
]

@pulumi.output_type
class ListRegistryGarbageCollectionsProperties:
    def __init__(__self__, garbage_collections=None):
        if garbage_collections and not isinstance(garbage_collections, list):
            raise TypeError("Expected argument 'garbage_collections' to be a list")
        pulumi.set(__self__, "garbage_collections", garbage_collections)

    @property
    @pulumi.getter(name="garbageCollections")
    def garbage_collections(self) -> Optional[Sequence['outputs.GarbageCollection']]:
        return pulumi.get(self, "garbage_collections")


class AwaitableListRegistryGarbageCollectionsProperties(ListRegistryGarbageCollectionsProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListRegistryGarbageCollectionsProperties(
            garbage_collections=self.garbage_collections)


def list_registry_garbage_collections(registry_name: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListRegistryGarbageCollectionsProperties:
    """
    Use this data source to access information about an existing resource.

    :param str registry_name: The name of a container registry.
    """
    __args__ = dict()
    __args__['registryName'] = registry_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:registry/v2:listRegistryGarbageCollections', __args__, opts=opts, typ=ListRegistryGarbageCollectionsProperties).value

    return AwaitableListRegistryGarbageCollectionsProperties(
        garbage_collections=pulumi.get(__ret__, 'garbage_collections'))


@_utilities.lift_output_func(list_registry_garbage_collections)
def list_registry_garbage_collections_output(registry_name: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListRegistryGarbageCollectionsProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str registry_name: The name of a container registry.
    """
    ...
