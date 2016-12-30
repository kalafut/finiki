from setuptools import setup

setup(
    name='finiki',
    packages=['finiki'],
    include_package_data=True,
    install_requires=[
        'flask',
        'mistune',
    ],
)

