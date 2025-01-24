import json
import os

with open('data.json', 'r') as f:
    docs = json.load(f)

os.makedirs('docs', exist_ok=True)
for filename, content in docs.items():
    with open(f'docs/{filename}', 'w') as f:
        f.write(content)