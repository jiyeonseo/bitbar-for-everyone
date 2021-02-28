#!/usr/local/bin/python3

# <bitbar.title>Links</bitbar.title>
# <bitbar.version>v1.0</bitbar.version>
# <bitbar.author>Jiyeon Seo</bitbar.author>
# <bitbar.author.github>jiyeonseo</bitbar.author.github>
# <bitbar.desc>Bookmarks on the bar</bitbar.desc>
# <bitbar.dependencies>python</bitbar.dependencies>

import json
import os

import base64

ME_PATH = os.path.realpath(__file__)
ROOT = os.path.dirname(ME_PATH)
META = os.path.join(ROOT, './json/link.json')

def print_title(toprint):
    print(toprint)
    print('---')

if __name__ == '__main__':
    print_title("Bookmarks | color=green ")
    
    with open(META, 'r') as json_file:
        json_data = json.load(json_file)
        for data in json_data:
            print(f"{data}")
            for d in json_data[data]: 
                print(f"{d} | href={json_data[data][d]}")
            print("---")
    print(f'configuration | bash="/usr/bin/open" param1="-a" param2 ="TextEdit" param3={META}')
    