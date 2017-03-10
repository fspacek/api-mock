# api-mock
Simple mock server for JSON API. Json files are served from defined directory, provided as CLI argument. 

Usage:
```bash
./api-mock -d /tmp/api-mock/
```
Sample output:
```
2017/03/10 19:27:13 Starting api-mock from dir /tmp/api-mock/
2017/03/10 19:27:13 Mapping file contact.json to path contact/1/
2017/03/10 19:27:13 Mapping file contact.json to path contact/2/
2017/03/10 19:27:13 Mapping file contacts.json to path contact/
2017/03/10 19:27:13 Creating handler for directory structure
```
