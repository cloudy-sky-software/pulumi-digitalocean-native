# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ._enums import *

__all__ = [
    'ListVolumeActionsItems',
    'AwaitableListVolumeActionsItems',
    'list_volume_actions',
    'list_volume_actions_output',
]

@pulumi.output_type
class ListVolumeActionsItems:
    def __init__(__self__, actions=None, links=None, meta=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence['outputs.VolumeAction']]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")


class AwaitableListVolumeActionsItems(ListVolumeActionsItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListVolumeActionsItems(
            actions=self.actions,
            links=self.links,
            meta=self.meta)


def list_volume_actions(volume_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListVolumeActionsItems:
    """
    Use this data source to access information about an existing resource.

    :param str volume_id: The ID of the block storage volume.
    """
    __args__ = dict()
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:volumes/v2:listVolumeActions', __args__, opts=opts, typ=ListVolumeActionsItems).value

    return AwaitableListVolumeActionsItems(
        actions=pulumi.get(__ret__, 'actions'),
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'))
def list_volume_actions_output(volume_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListVolumeActionsItems]:
    """
    Use this data source to access information about an existing resource.

    :param str volume_id: The ID of the block storage volume.
    """
    __args__ = dict()
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:volumes/v2:listVolumeActions', __args__, opts=opts, typ=ListVolumeActionsItems)
    return __ret__.apply(lambda __response__: ListVolumeActionsItems(
        actions=pulumi.get(__response__, 'actions'),
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta')))
