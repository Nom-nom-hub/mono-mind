#!/usr/bin/env python3
import sys
import os
import datetime

print("Post-build plugin executed")
print(f"Current directory: {os.getcwd()}")
print(f"Timestamp: {datetime.datetime.now()}")
print(f"Arguments: {sys.argv}")