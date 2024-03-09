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
    'ListDomainsResult',
    'AwaitableListDomainsResult',
    'list_domains',
    'list_domains_output',
]

@pulumi.output_type
class ListDomainsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListDomains':
        return pulumi.get(self, "items")


class AwaitableListDomainsResult(ListDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDomainsResult(
            items=self.items)


def list_domains(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDomainsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:domains/v2:listDomains', __args__, opts=opts, typ=ListDomainsResult).value

    return AwaitableListDomainsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_domains)
def list_domains_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDomainsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
