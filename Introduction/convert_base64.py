#!/usr/bin/python3
import base64

def convert_to_base64(text):
    encoded = base64.b64encode(text).encode()
    return encoded

x=input("Enter String>> ")
result=convert_to_base64(x)
print(x)

