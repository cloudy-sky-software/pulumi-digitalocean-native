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
    'ListDomainsRecordsResult',
    'AwaitableListDomainsRecordsResult',
    'list_domains_records',
    'list_domains_records_output',
]

@pulumi.output_type
class ListDomainsRecordsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListDomainsRecordsItems':
        return pulumi.get(self, "items")


class AwaitableListDomainsRecordsResult(ListDomainsRecordsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDomainsRecordsResult(
            items=self.items)


def list_domains_records(domain_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDomainsRecordsResult:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: The name of the domain itself.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:domains/v2:listDomainsRecords', __args__, opts=opts, typ=ListDomainsRecordsResult).value

    return AwaitableListDomainsRecordsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_domains_records)
def list_domains_records_output(domain_name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDomainsRecordsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: The name of the domain itself.
    """
    ...
