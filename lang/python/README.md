# Pyemo

[![pub package](https://img.shields.io/pypi/v/pyemo)](https://pypi.org/project/pyemo/)

## Install

```
pip install pyemo
```

## Usage

Declare an instance namespace

```python
from pyemo import Emo

emo = Emo("pipeline")
```

Then use it in the code

```python
emo.start("Starting data pipeline")
emo.arrow_in("Loading csv file")
# ...
emo.data("Processing data transformation")
# ...
emo.query("Processing indexes")
# ...
emo.ok("Finished")
```