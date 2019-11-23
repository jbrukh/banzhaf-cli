#!/usr/bin/env python

import sys

if __name__ == '__main__':

    weights = sorted([int(float(line)) for line in sys.stdin], reverse=True)
    weights = [weight for weight in weights if weight != 0]
    quota = int(sum(weights) / 2) + 1

    print(" ".join(str(item) for item in [quota]+weights))