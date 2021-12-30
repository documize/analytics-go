# Zerabase Go Library

Zerabase provides a system-of-record for your products by unifying product deliver, customer conversations and app
analytics.

This module is for Go based applications that need to send app analytics and customer interactions into Zerabase. For
example, record user sign-up, customer purchase, contact form, downloaded, etc.

## Usage

## App Analytics

Send app analytics from your Go app.

1. Create a Metric within Zerabase.
2. Note the Tracker ID for the Metric.
3. ```go get github.com/documize/zerabase-go``` module and invoke as follows:

```
metric.TrackEvent("your-tracker-id")
```

## Customer Interactions

Coming soon.
