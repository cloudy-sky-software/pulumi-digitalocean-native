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

__all__ = [
    'ListSizesItems',
    'AwaitableListSizesItems',
    'list_sizes',
    'list_sizes_output',
]

@pulumi.output_type
class ListSizesItems:
    def __init__(__self__, links=None, meta=None, sizes=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if sizes and not isinstance(sizes, list):
            raise TypeError("Expected argument 'sizes' to be a list")
        pulumi.set(__self__, "sizes", sizes)

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
    def sizes(self) -> Sequence['outputs.Size']:
        return pulumi.get(self, "sizes")


class AwaitableListSizesItems(ListSizesItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListSizesItems(
            links=self.links,
            meta=self.meta,
            sizes=self.sizes)


def list_sizes(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListSizesItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:sizes/v2:listSizes', __args__, opts=opts, typ=ListSizesItems).value

    return AwaitableListSizesItems(
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'),
        sizes=pulumi.get(__ret__, 'sizes'))


@_utilities.lift_output_func(list_sizes)
def list_sizes_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListSizesItems]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
