#!/usr/bin/env python

import random
import sys

def forfor(a):
    return [item for sublist in a for item in sublist]


if __name__ == '__main__':
 
  if len(sys.argv) == 1:
    max = 70
    mind = 1
    maxd = 2
    x = [random.randint(mind,maxd) for x in range(max)] # + [1 for x in range(random.randint(0, 2*max))]
    # times = int(max / maxd)
    # x = forfor(times*[item] for item in range(1, maxd+1))
  else:
    x = [int(x) for x in sys.argv[1:]]
  
  y = [int(sum(x)/2+1)] + sorted(x, reverse=True)
  print(" ".join(str(item) for item in y))

