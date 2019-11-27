#!/usr/bin/env python

import sys
import argparse

def process(x, divisor=1):
    x = x.strip('"\n')
    x = float(x) / divisor
    x = round(x, 0)
    return int(x)

if __name__ == '__main__':

    parser = argparse.ArgumentParser()
    parser.add_argument('-d', dest='divisor', type=int, default=1)
    args = parser.parse_args()

    weights = sorted([process(line, args.divisor) for line in sys.stdin], reverse=True)
    weights = [weight for weight in weights if weight != 0]
    quota = int(sum(weights) / 2) + 1

    print(" ".join(str(item) for item in [quota]+weights))
