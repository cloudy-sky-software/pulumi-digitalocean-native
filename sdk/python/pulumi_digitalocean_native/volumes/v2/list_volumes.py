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
    'ListVolumesItems',
    'AwaitableListVolumesItems',
    'list_volumes',
    'list_volumes_output',
]

@pulumi.output_type
class ListVolumesItems:
    def __init__(__self__, links=None, meta=None, volumes=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if volumes and not isinstance(volumes, list):
            raise TypeError("Expected argument 'volumes' to be a list")
        pulumi.set(__self__, "volumes", volumes)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def volumes(self) -> Sequence['outputs.VolumeFull']:
        """
        Array of volumes.
        """
        return pulumi.get(self, "volumes")


class AwaitableListVolumesItems(ListVolumesItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListVolumesItems(
            links=self.links,
            meta=self.meta,
            volumes=self.volumes)


def list_volumes(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListVolumesItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:volumes/v2:listVolumes', __args__, opts=opts, typ=ListVolumesItems).value

    return AwaitableListVolumesItems(
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'),
        volumes=pulumi.get(__ret__, 'volumes'))
def list_volumes_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListVolumesItems]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:volumes/v2:listVolumes', __args__, opts=opts, typ=ListVolumesItems)
    return __ret__.apply(lambda __response__: ListVolumesItems(
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta'),
        volumes=pulumi.get(__response__, 'volumes')))
