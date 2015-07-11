parsesearch
===========

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

This project shows example use of Cloud Code Webhooks to implement full text search for Parse objects.

Actual use will likey involve customizations.

Getting Started
---------------

The easiest way to get started is by using the 'Deploy to Heroku' button above to start running this project in a new Heroku dyno.

You will be prompted for your Parse Application's keys and the Parse Class you would like to index.

To auto-create the necessary triggers and webhook functions you must set the dyno url as an environment variable for your dyno.

After your app deploys successfully configure the endpoint URL like so:

* Manage App -> Settings -> Reveal Config Vars

Add a new environment variable named 'URL' and set its value to the URL of the Heroku dyno.

For example, you would set it to a string of the form 'https://DYNO_NAME_HERE.herokuapp.com/'

The dyno will restart and register the appropriate webhooks in your Parse app. You can confirm this by looking at your Webhooks page in the Parse web interface. If you set this to the incorrect value at first remove the webhooks in the Parse web interface and restart your dyno.

Querying
--------
parsesearch installs a webhook called 'search'

You can test a search by curling your Cloud Code Webhook like so:

```sh
$ curl -X POST https://${PARSE_APPLICATION_ID}:javascript-key:${PARSE_JAVASCRIPT_KEY}@api.parse.com/1/functions/search -d '{"q":"hello"}' 
{"result":["gKCRJ0tzgA","yWpBOZj4UB"]}
```

How it works
------------
On startup parsesearch iterates over all the objects in your specified class and indexes them (their full JSON representation). It registers an afterSave and afterDelete triggers to maintain the index.

