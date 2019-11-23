#!/usr/bin/env python

import sys

if __name__ == '__main__':
  count = {}
  for arg in sys.argv[1:]:
    if arg in count:
      count[arg] += 1
    else:
      count[arg] = 1
  print(count)
