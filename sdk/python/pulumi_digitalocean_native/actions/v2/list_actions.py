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
    'ListActionsItems',
    'AwaitableListActionsItems',
    'list_actions',
    'list_actions_output',
]

@pulumi.output_type
class ListActionsItems:
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
    def actions(self) -> Optional[Sequence['outputs.Action']]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")


class AwaitableListActionsItems(ListActionsItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListActionsItems(
            actions=self.actions,
            links=self.links,
            meta=self.meta)


def list_actions(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListActionsItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:actions/v2:listActions', __args__, opts=opts, typ=ListActionsItems).value

    return AwaitableListActionsItems(
        actions=pulumi.get(__ret__, 'actions'),
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'))


@_utilities.lift_output_func(list_actions)
def list_actions_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListActionsItems]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
