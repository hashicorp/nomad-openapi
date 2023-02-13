# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


from setuptools import setup, find_packages  # noqa: H301

NAME = "nomad-client"
VERSION = "1.1.4"
# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = [
  "urllib3 >= 1.25.3",
  "python-dateutil",
]

setup(
    name=NAME,
    version=VERSION,
    description="Nomad",
    author="OpenAPI Generator community",
    author_email="support@hashicorp.com",
    url="",
    keywords=["OpenAPI", "OpenAPI-Generator", "Nomad"],
    python_requires=">=3.6",
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    license="MPL 2",
    long_description="""\
    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501
    """
)
