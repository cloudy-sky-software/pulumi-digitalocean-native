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
from ._enums import *

__all__ = [
    'ListBillingHistoryItems',
    'AwaitableListBillingHistoryItems',
    'list_billing_history',
    'list_billing_history_output',
]

@pulumi.output_type
class ListBillingHistoryItems:
    def __init__(__self__, billing_history=None, links=None, meta=None):
        if billing_history and not isinstance(billing_history, list):
            raise TypeError("Expected argument 'billing_history' to be a list")
        pulumi.set(__self__, "billing_history", billing_history)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter(name="billingHistory")
    def billing_history(self) -> Optional[Sequence['outputs.BillingHistory']]:
        return pulumi.get(self, "billing_history")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaProperties':
        """
        Information about the response itself.
        """
        return pulumi.get(self, "meta")


class AwaitableListBillingHistoryItems(ListBillingHistoryItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListBillingHistoryItems(
            billing_history=self.billing_history,
            links=self.links,
            meta=self.meta)


def list_billing_history(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListBillingHistoryItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:customers/v2:listBillingHistory', __args__, opts=opts, typ=ListBillingHistoryItems).value

    return AwaitableListBillingHistoryItems(
        billing_history=pulumi.get(__ret__, 'billing_history'),
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'))
def list_billing_history_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListBillingHistoryItems]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:customers/v2:listBillingHistory', __args__, opts=opts, typ=ListBillingHistoryItems)
    return __ret__.apply(lambda __response__: ListBillingHistoryItems(
        billing_history=pulumi.get(__response__, 'billing_history'),
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta')))
