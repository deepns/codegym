# Learnings

Noting down some learnings along the way

## python func args initialization

```python
def func_args_init():
    # the argument values is initialized with an empty list when interpreted.
    # subsequent calls to f() will use the same object assigned to values assigned during initialization
    def f(i, values = []):
        values.append(i)
        print(values)
        return values

    # will print [1], [1, 2], [1, 2, 3]
    f(1)
    f(2)
    f(3)

func_args_init()
```

## pretty python with pprint.pprint

Quick way to pretty print built in data types in python.

```python
>>> # data download from 'https://pypi.org/pypi/sampleproject/json'
>>> print(data)
{'author': 'The Python Packaging Authority', 'author_email': 'pypa-dev@googlegroups.com', 'bugtrack_url': None, 'classifiers': ['Development Status :: 3 - Alpha', 'Intended Audience :: Developers', 'License :: OSI Approved :: MIT License', 'Programming Language :: Python :: 2', 'Programming Language :: Python :: 2.6', 'Programming Language :: Python :: 2.7', 'Programming Language :: Python :: 3', 'Programming Language :: Python :: 3.2', 'Programming Language :: Python :: 3.3', 'Programming Language :: Python :: 3.4', 'Topic :: Software Development :: Build Tools'], 'description': 'A sample Python project\n=======================\n\nThis is the description file for the project.\n\nThe file should use UTF-8 encoding and be written using ReStructured Text. It\nwill be used to generate the project webpage on PyPI, and should be written for\nthat purpose.\n\nTypical contents for this file would include an overview of the project, basic\nusage examples, etc. Generally, including the project changelog in here is not\na good idea, although a simple "What\'s New" section for the most recent version\nmay be appropriate.', 'description_content_type': None, 'docs_url': None, 'download_url': 'UNKNOWN', 'downloads': {'last_day': -1, 'last_month': -1, 'last_week': -1}, 'home_page': 'https://github.com/pypa/sampleproject', 'keywords': 'sample setuptools development', 'license': 'MIT', 'maintainer': None, 'maintainer_email': None, 'name': 'sampleproject', 'package_url': 'https://pypi.org/project/sampleproject/', 'platform': 'UNKNOWN', 'project_url': 'https://pypi.org/project/sampleproject/', 'project_urls': {'Download': 'UNKNOWN', 'Homepage': 'https://github.com/pypa/sampleproject'}, 'release_url': 'https://pypi.org/project/sampleproject/1.2.0/', 'requires_dist': None, 'requires_python': None, 'summary': 'A sample Python project', 'version': '1.2.0'}
>>> pprint.pprint(data, indent=2, depth=2)
{ 'author': 'The Python Packaging Authority',
  'author_email': 'pypa-dev@googlegroups.com',
  'bugtrack_url': None,
  'classifiers': [ 'Development Status :: 3 - Alpha',
                   'Intended Audience :: Developers',
                   'License :: OSI Approved :: MIT License',
                   'Programming Language :: Python :: 2',
                   'Programming Language :: Python :: 2.6',
                   'Programming Language :: Python :: 2.7',
                   'Programming Language :: Python :: 3',
                   'Programming Language :: Python :: 3.2',
                   'Programming Language :: Python :: 3.3',
                   'Programming Language :: Python :: 3.4',
                   'Topic :: Software Development :: Build Tools'],
  'description': 'A sample Python project\n'
  ...
>>> pprint.pprint(data, indent=2, depth=1)
{ 'author': 'The Python Packaging Authority',
  'author_email': 'pypa-dev@googlegroups.com',
  'bugtrack_url': None,
  'classifiers': [...],
  'description': 'A sample Python project\n'
                 '=======================\n'
                 '\n'
                 'This is the description file for the project.\n'
                 '\n'
                 'The file should use UTF-8 encoding and be written using '
                 'ReStructured Text. It\n'
                 'will be used to generate the project webpage on PyPI, and '
                 'should be written for\n'
                 'that purpose.\n'
                 '\n'
                 'Typical contents for this file would include an overview of '
                 'the project, basic\n'
                 'usage examples, etc. Generally, including the project '
                 'changelog in here is not\n'
                 'a good idea, although a simple "What\'s New" section for the '
                 'most recent version\n'
                 'may be appropriate.',
  'description_content_type': None,
  'docs_url': None,
  'download_url': 'UNKNOWN',
  'downloads': {...},
```
