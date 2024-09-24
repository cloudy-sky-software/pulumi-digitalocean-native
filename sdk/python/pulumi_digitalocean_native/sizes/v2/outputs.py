# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs

__all__ = [
    'MetaMeta',
    'PageLinks',
    'PageLinksPagesProperties',
    'Size',
]

@pulumi.output_type
class MetaMeta(dict):
    def __init__(__self__, *,
                 total: Optional[int] = None):
        """
        :param int total: Number of objects returned by the request.
        """
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def total(self) -> Optional[int]:
        """
        Number of objects returned by the request.
        """
        return pulumi.get(self, "total")


@pulumi.output_type
class PageLinks(dict):
    def __init__(__self__, *,
                 pages: Optional['outputs.PageLinksPagesProperties'] = None):
        if pages is not None:
            pulumi.set(__self__, "pages", pages)

    @property
    @pulumi.getter
    def pages(self) -> Optional['outputs.PageLinksPagesProperties']:
        return pulumi.get(self, "pages")


@pulumi.output_type
class PageLinksPagesProperties(dict):
    def __init__(__self__, *,
                 first: Optional[str] = None,
                 last: Optional[str] = None,
                 next: Optional[str] = None,
                 prev: Optional[str] = None):
        if first is not None:
            pulumi.set(__self__, "first", first)
        if last is not None:
            pulumi.set(__self__, "last", last)
        if next is not None:
            pulumi.set(__self__, "next", next)
        if prev is not None:
            pulumi.set(__self__, "prev", prev)

    @property
    @pulumi.getter
    def first(self) -> Optional[str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def prev(self) -> Optional[str]:
        return pulumi.get(self, "prev")


@pulumi.output_type
class Size(dict):
    def __init__(__self__, *,
                 available: Optional[bool] = None,
                 description: str,
                 disk: int,
                 memory: int,
                 price_hourly: float,
                 price_monthly: float,
                 regions: Sequence[str],
                 slug: str,
                 transfer: float,
                 vcpus: int):
        """
        :param bool available: This is a boolean value that represents whether new Droplets can be created with this size.
        :param str description: A string describing the class of Droplets created from this size. For example: Basic, General Purpose, CPU-Optimized, Memory-Optimized, or Storage-Optimized.
        :param int disk: The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes.
        :param int memory: The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes.
        :param float price_hourly: This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.
        :param float price_monthly: This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars.
        :param Sequence[str] regions: An array containing the region slugs where this size is available for Droplet creates.
        :param str slug: A human-readable string that is used to uniquely identify each size.
        :param float transfer: The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.
        :param int vcpus: The integer of number CPUs allocated to Droplets of this size.
        """
        if available is None:
            available = True
        pulumi.set(__self__, "available", available)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "disk", disk)
        pulumi.set(__self__, "memory", memory)
        pulumi.set(__self__, "price_hourly", price_hourly)
        pulumi.set(__self__, "price_monthly", price_monthly)
        pulumi.set(__self__, "regions", regions)
        pulumi.set(__self__, "slug", slug)
        pulumi.set(__self__, "transfer", transfer)
        pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter
    def available(self) -> bool:
        """
        This is a boolean value that represents whether new Droplets can be created with this size.
        """
        return pulumi.get(self, "available")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A string describing the class of Droplets created from this size. For example: Basic, General Purpose, CPU-Optimized, Memory-Optimized, or Storage-Optimized.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disk(self) -> int:
        """
        The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes.
        """
        return pulumi.get(self, "disk")

    @property
    @pulumi.getter
    def memory(self) -> int:
        """
        The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="priceHourly")
    def price_hourly(self) -> float:
        """
        This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.
        """
        return pulumi.get(self, "price_hourly")

    @property
    @pulumi.getter(name="priceMonthly")
    def price_monthly(self) -> float:
        """
        This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars.
        """
        return pulumi.get(self, "price_monthly")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[str]:
        """
        An array containing the region slugs where this size is available for Droplet creates.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def slug(self) -> str:
        """
        A human-readable string that is used to uniquely identify each size.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def transfer(self) -> float:
        """
        The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.
        """
        return pulumi.get(self, "transfer")

    @property
    @pulumi.getter
    def vcpus(self) -> int:
        """
        The integer of number CPUs allocated to Droplets of this size.
        """
        return pulumi.get(self, "vcpus")


